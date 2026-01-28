/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package api

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	hashtypes "github.com/origadmin/toolkits/crypto/hash/types"

	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/enttest"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/features/system/dal"
	"origadmin/application/admin/internal/features/system/service"
)

// testServerComponents holds all components needed for an integration test.
type testServerComponents struct {
	DBClient   *ent.Client
	HTTPServer *httptest.Server
	Ctx        context.Context
}

// noopPublisher is a no-op publisher for testing.
type noopPublisher struct{}

func (n *noopPublisher) Publish(topic string, messages ...*message.Message) error {
	return nil
}

func (n *noopPublisher) Close() error {
	return nil
}

// setupTestServer initializes a test server with an in-memory SQLite database
// and all necessary dependencies for integration testing.
// 修正了数据库初始化: 正确使用 enttest.Open 的参数
func setupTestServer(t *testing.T) *testServerComponents {
	t.Helper()

	// 1. Initialize in-memory SQLite database
	// 修正: enttest.Open 的第二个参数是 dialect，第三个参数才是 DSN
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	t.Cleanup(func() { client.Close() })
	database := ent.NewDatabaseWithClient(client)

	// 2. Manually construct the dependency graph (DAL -> Biz -> Service -> Server)
	logger := log.NewStdLogger(io.Discard)
	hasher, err := hash.NewCrypto(hashtypes.BCRYPT)
	require.NoError(t, err)

	// DAL Layer
	userRepo := dal.NewUserRepo(database)
	roleRepo := dal.NewRoleRepo(database)
	permissionRepo := dal.NewPermissionRepo(database)
	resourceRepo := dal.NewResourceRepo(database)
	viewRepo := dal.NewViewRepo(database)
	authzRepo := dal.NewAuthorizationRepo(database, logger)

	// Biz Layer
	userUseCase := biz.NewUserUseCase(userRepo, hasher, logger)
	roleUseCase := biz.NewRoleUseCase(roleRepo)
	permissionUseCase := biz.NewPermissionUseCase(permissionRepo)
	resourceUseCase := biz.NewResourceUseCase(resourceRepo)
	viewUseCase := biz.NewViewUseCase(viewRepo)
	authzUseCase := biz.NewAuthorizationUseCase(authzRepo, logger)

	// Service Layer
	// For testing, we can use a no-op publisher
	noopPublisher := &noopPublisher{}
	userService := service.NewUserService(userUseCase, noopPublisher, logger)
	roleService := service.NewRoleService(roleUseCase)
	permissionService := service.NewPermissionService(permissionUseCase)
	resourceService := service.NewResourceService(resourceUseCase)
	viewService := service.NewViewService(viewUseCase)
	authzService := service.NewAuthorizationService(authzUseCase, logger)

	systemService := service.NewSystemService(
		resourceService,
		roleService,
		userService,
		permissionService,
		viewService,
		authzService,
	)

	// 3. Setup HTTP server using Gin
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Manually register routes similar to how Kratos would.
	RegisterRoleServiceRoutes(router, systemService.Role)
	RegisterPermissionServiceRoutes(router, systemService.Permission)

	// 4. Create httptest server
	testServer := httptest.NewServer(router)
	t.Cleanup(func() { testServer.Close() })

	return &testServerComponents{
		DBClient:   client,
		HTTPServer: testServer,
		Ctx:        context.Background(),
	}
}

// handleServiceError translates service-layer errors into appropriate HTTP responses.
func handleServiceError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	// 导入 errors 包检查错误类型
	switch {
	case err.Error() == "not found":
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case err.Error() == "bad request":
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// RegisterRoleServiceRoutes manually registers HTTP routes for RoleService.
func RegisterRoleServiceRoutes(r *gin.Engine, svc *service.RoleService) {
	g := r.Group("/v1/system/roles")
	{
		g.POST("", func(c *gin.Context) {
			req := &systemv1.CreateRoleRequest{}
			if err := c.ShouldBindJSON(req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resp, err := svc.CreateRole(c.Request.Context(), req)
			if err != nil {
				handleServiceError(c, err)
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		g.GET("/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role ID"})
				return
			}
			req := &systemv1.GetRoleRequest{Id: id}
			resp, err := svc.GetRole(c.Request.Context(), req)
			if err != nil {
				handleServiceError(c, err)
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		g.PUT("/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role ID"})
				return
			}
			req := &systemv1.UpdateRoleRequest{}
			if err := c.ShouldBindJSON(req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if req.Role != nil && req.Role.Id != 0 && req.Role.Id != id {
				c.JSON(http.StatusBadRequest, gin.H{"error": "role ID in path and body mismatch"})
				return
			}
			if req.Role == nil {
				req.Role = &types.Role{}
			}
			req.Role.Id = id // Set ID from path
			resp, err := svc.UpdateRole(c.Request.Context(), req)
			if err != nil {
				handleServiceError(c, err)
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		g.DELETE("/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role ID"})
				return
			}
			req := &systemv1.DeleteRoleRequest{Id: id}
			resp, err := svc.DeleteRole(c.Request.Context(), req)
			if err != nil {
				handleServiceError(c, err)
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		g.GET("", func(c *gin.Context) {
			req := &systemv1.ListRolesRequest{}
			if err := c.ShouldBindQuery(req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resp, err := svc.ListRoles(c.Request.Context(), req)
			if err != nil {
				handleServiceError(c, err)
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		g.PUT("/:id/permissions", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role ID"})
				return
			}
			var req struct {
				PermissionIDs []int64 `json:"permission_ids"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			err = svc.UpdateRolePermissions(c.Request.Context(), id, req.PermissionIDs)
			if err != nil {
				handleServiceError(c, err)
				return
			}
			c.Status(http.StatusNoContent)
		})
	}
}

// RegisterPermissionServiceRoutes manually registers HTTP routes for PermissionService.
func RegisterPermissionServiceRoutes(r *gin.Engine, svc *service.PermissionService) {
	g := r.Group("/v1/system/permissions")
	{
		g.POST("", func(c *gin.Context) {
			req := &systemv1.CreatePermissionRequest{}
			if err := c.ShouldBindJSON(req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resp, err := svc.CreatePermission(c.Request.Context(), req)
			if err != nil {
				handleServiceError(c, err)
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		g.GET("/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid permission ID"})
				return
			}
			req := &systemv1.GetPermissionRequest{Id: id}
			resp, err := svc.GetPermission(c.Request.Context(), req)
			if err != nil {
				handleServiceError(c, err)
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		g.PUT("/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid permission ID"})
				return
			}
			req := &systemv1.UpdatePermissionRequest{}
			if err := c.ShouldBindJSON(req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if req.Permission != nil && req.Permission.Id != 0 && req.Permission.Id != id {
				c.JSON(http.StatusBadRequest, gin.H{"error": "permission ID in path and body mismatch"})
				return
			}
			if req.Permission == nil {
				req.Permission = &types.Permission{}
			}
			req.Permission.Id = id // Set ID from path
			resp, err := svc.UpdatePermission(c.Request.Context(), req)
			if err != nil {
				handleServiceError(c, err)
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		g.DELETE("/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid permission ID"})
				return
			}
			req := &systemv1.DeletePermissionRequest{Id: id}
			resp, err := svc.DeletePermission(c.Request.Context(), req)
			if err != nil {
				handleServiceError(c, err)
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		g.GET("", func(c *gin.Context) {
			req := &systemv1.ListPermissionsRequest{}
			if err := c.ShouldBindQuery(req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resp, err := svc.ListPermissions(c.Request.Context(), req)
			if err != nil {
				handleServiceError(c, err)
				return
			}
			c.JSON(http.StatusOK, resp)
		})
	}
}
