/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package seeder contains the implementation for database seeding tasks.
package seeder

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"unicode"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/origadmin/contrib/security"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/features/system/dto"
)

// ProviderSet is for wire injection.
var ProviderSet = wire.NewSet(NewSeeder)

const (
	passwordCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	passwordLength  = 16
)

// Seeder is the container for initialization tasks.
type Seeder struct {
	userUseCase     *biz.UserUseCase
	resourceUseCase *biz.ResourceUseCase
	viewUseCase     *biz.ViewUseCase
	rootUserCfg     *confpb.RootUser
	log             *log.Helper
}

// NewSeeder creates a new Seeder.
func NewSeeder(userUseCase *biz.UserUseCase, resourceUseCase *biz.ResourceUseCase, viewUseCase *biz.ViewUseCase, cfg *conf.Config, logger log.Logger) (*Seeder, error) {
	return &Seeder{
		userUseCase:     userUseCase,
		resourceUseCase: resourceUseCase,
		viewUseCase:     viewUseCase,
		rootUserCfg:     cfg.RootUser(),
		log:             log.NewHelper(log.With(logger, "module", "seeder")),
	}, nil
}

// Run executes all seeding tasks.
func (s *Seeder) Run() error {
	if err := s.createRootUser(); err != nil {
		return err
	}
	if err := s.createInitialResources(); err != nil {
		return err
	}
	if err := s.createInitialViews(); err != nil {
		return err
	}
	return nil
}

// createRootUser creates the initial administrator user if it does not exist.
// If the password in the config is empty, a random one will be generated and printed.
func (s *Seeder) createRootUser() error {
	if s.rootUserCfg == nil || !s.rootUserCfg.Enabled {
		s.log.Info("Root user seeding is disabled in config.")
		return nil
	}

	ctx := context.Background()
	username := s.rootUserCfg.GetUsername()

	// 1. Check if the user already exists by trying to list them.
	listReq := &system.ListUsersRequest{
		Keyword:  username,
		PageSize: 1,
	}
	existingUsers, total, err := s.userUseCase.ListUsers(ctx, listReq)
	if err != nil {
		s.log.Errorf("Failed to check for root user: %v", err)
		return err
	}
	if total > 0 || len(existingUsers) > 0 {
		s.log.Infof("Root user '%s' already exists, skipping creation.", username)
		return nil
	}

	// 2. If not found, create the new user.
	s.log.Infof("Root user '%s' not found, creating...", username)

	password := s.rootUserCfg.GetPassword()
	if password == "" {
		var err error
		password, err = generateRandomPassword(passwordLength)
		if err != nil {
			s.log.Errorf("Failed to generate random password: %v", err)
			return err
		}
		// Use fmt.Printf with ANSI colors for high visibility
		fmt.Printf("\n")
		fmt.Printf("\033[1;32m==================== [Root User Created] ===================\033[0m\n")
		fmt.Printf("\033[1;33mUsername: %s\033[0m\n", username)
		fmt.Printf("\033[1;33mPassword: %s\033[0m\n", password)
		fmt.Printf("\033[1;32m============================================================\033[0m\n")
		fmt.Printf("\n")
	}

	newUser := &types.User{
		Username: username,
		Nickname: s.rootUserCfg.GetNickname(),
		Email:    s.rootUserCfg.GetEmail(),
	}

	createdUser, err := s.userUseCase.CreateUser(ctx, newUser, password)
	if err != nil {
		s.log.Errorf("Failed to create root user '%s': %v", username, err)
		return err
	}

	s.log.Infof("Successfully created root user '%s' with ID: %d", createdUser.Username, createdUser.Id)
	return nil
}

func (s *Seeder) createInitialResources() error {
	ctx := context.Background()

	// Collect resources by module:service
	moduleResources := make(map[string]map[string][]*security.Policy)

	for _, policy := range security.RegisteredPolicies() {
		// Parse gRPC method: /package.Service/Method
		// e.g. /api.v1.services.auth.AuthService/Login
		parts := strings.Split(policy.ServiceMethod, "/")
		if len(parts) < 3 {
			s.log.Warnf("Skipping malformed service method: %s", policy.ServiceMethod)
			continue
		}

		fullService := parts[1]

		// Parse Service: api.v1.services.auth.AuthService
		serviceParts := strings.Split(fullService, ".")
		if len(serviceParts) < 2 {
			s.log.Warnf("Skipping malformed service name: %s", fullService)
			continue
		}

		// Extract Module Name (e.g. "auth" from "api.v1.services.auth.AuthService")
		var moduleName string
		if len(serviceParts) >= 2 {
			moduleName = serviceParts[len(serviceParts)-2]
		} else {
			moduleName = "system"
		}

		// Extract Resource Name (e.g. "Auth" from "AuthService")
		serviceName := serviceParts[len(serviceParts)-1]

		// Ensure module and service collections exist
		if moduleResources[moduleName] == nil {
			moduleResources[moduleName] = make(map[string][]*security.Policy)
		}
		if moduleResources[moduleName][serviceName] == nil {
			moduleResources[moduleName][serviceName] = []*security.Policy{}
		}

		moduleResources[moduleName][serviceName] = append(moduleResources[moduleName][serviceName], &policy)
	}

	// Create resources with 3-level structure: Module -> Service -> API
	seq := 1
	for moduleName, services := range moduleResources {
		// Create module level resource
		moduleKeyword := toSnakeCase(moduleName)
		moduleNameDisplay := toTitleCase(moduleName) + " Module"
		moduleI18n := "resource." + moduleKeyword + ".module"
		moduleTreePath := "/" + moduleKeyword

		moduleResource := &dto.ResourceFromPolicyInput{
			Resource: types.Resource{
				Name:        moduleNameDisplay,
				I18N:        moduleI18n,
				Sequence:    int32(seq),
				Keyword:     moduleKeyword,
				ServiceName: moduleName,
				Type:        "module",
				TreePath:    moduleTreePath,
			},
			Policy: nil, // No policy for module
		}

		moduleID, err := s.createOrUpdateResource(ctx, moduleResource, moduleKeyword, true)
		if err != nil {
			s.log.Errorf("failed to create module resource '%s': %v", moduleName, err)
			continue
		}
		seq++

		// Create service level resources
		for serviceName, policies := range services {
			resourceName := strings.TrimSuffix(serviceName, "Service")
			serviceKeyword := strings.Join([]string{moduleKeyword, toSnakeCase(resourceName)}, ":")
			serviceNameDisplay := toTitleCase(resourceName) + " Service"
			serviceI18n := "resource." + serviceKeyword + ".service"
			serviceTreePath := moduleTreePath + "/" + toSnakeCase(resourceName)

			serviceResource := &dto.ResourceFromPolicyInput{
				Resource: types.Resource{
					Name:        serviceNameDisplay,
					I18N:        serviceI18n,
					Sequence:    int32(seq),
					Keyword:     serviceKeyword,
					ServiceName: resourceName,
					Type:        "service",
					TreePath:    serviceTreePath,
					ParentId:    moduleID,
				},
				Policy: nil, // No policy for service
			}

			serviceID, err := s.createOrUpdateResource(ctx, serviceResource, serviceKeyword, false)
			if err != nil {
				s.log.Errorf("failed to create service resource '%s': %v", serviceName, err)
				continue
			}
			seq++

			// Create API level resources
			for _, policy := range policies {
				parts := strings.Split(policy.ServiceMethod, "/")
				method := parts[2]

				apiKeyword := strings.Join([]string{moduleKeyword, toSnakeCase(resourceName), toSnakeCase(method)}, ":")
				apiNameDisplay := toTitleCase(resourceName) + " " + toTitleCase(method)
				apiI18n := "resource." + apiKeyword + ".api"
				apiTreePath := serviceTreePath + "/" + toSnakeCase(method)

				// Extract method and path from GatewayPath
				var httpMethod, path string
				if policy.GatewayPath != "" {
					if gatewayParts := strings.SplitN(policy.GatewayPath, ":", 2); len(gatewayParts) == 2 {
						httpMethod = gatewayParts[0]
						path = gatewayParts[1]
					}
				}

				apiResource := &dto.ResourceFromPolicyInput{
					Resource: types.Resource{
						Name:        apiNameDisplay,
						I18N:        apiI18n,
						Sequence:    int32(seq),
						Keyword:     apiKeyword,
						ServiceName: resourceName,
						Type:        "api",
						TreePath:    apiTreePath,
						ParentId:    serviceID,
						Method:      httpMethod,
						Path:        path,
					},
					Policy: policy,
				}

				_, err := s.createOrUpdateResource(ctx, apiResource, apiKeyword, false)
				if err != nil {
					s.log.Errorf("failed to create API resource from policy '%s': %v", policy.ServiceMethod, err)
				} else {
					s.log.Infof("Successfully created resource from policy: %s", policy.ServiceMethod)
				}
				seq++
			}
		}
	}

	return nil
}

// createOrUpdateResource creates or updates a resource based on keyword
func (s *Seeder) createOrUpdateResource(ctx context.Context, input *dto.ResourceFromPolicyInput, keyword string, isModule bool) (int64, error) {
	// Check if resource already exists
	resources, count, err := s.resourceUseCase.ListResources(ctx,
		&system.ListResourcesRequest{
			Keyword:   keyword,
			OnlyCount: false,
			PageSize:  1,
		})
	if err != nil && count == 0 {
		s.log.Errorf("failed to check existing resource '%s': %v", keyword, err)
		return 0, err
	}

	if count > 0 && len(resources) > 0 {
		s.log.Infof("Resource '%s' already exists, skipping.", keyword)
		return resources[0].Id, nil
	}

	created, err := s.resourceUseCase.CreateResourceFromPolicy(ctx, input)
	if err != nil {
		return 0, err
	}

	return created.Id, nil
}

func (s *Seeder) createInitialViews() error {
	ctx := context.Background()
	views := []*types.View{
		{Name: "Dashboard", Keyword: "dashboard", Path: "/dashboard"}, // TODO: Add Component: "default" after proto regen
		{Name: "System", Keyword: "system", Path: "/system"},          // TODO: Add Component: "default" after proto regen
	}

	for _, view := range views {
		_, _, err := s.viewUseCase.ListViews(ctx, &system.ListViewsRequest{Keyword: view.Keyword})
		if err == nil {
			s.log.Infof("View '%s' already exists, skipping.", view.Name)
			continue
		}
		if _, err := s.viewUseCase.CreateView(ctx, view); err != nil {
			s.log.Errorf("failed to create view %s: %v", view.Name, err)
		} else {
			s.log.Infof("Successfully created view: %s", view.Name)
		}
	}
	return nil
}

// toTitleCase converts CamelCase to Title Case (e.g. "GetProfile" -> "Get Profile")
func toTitleCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune(' ')
		}
		result.WriteRune(r)
	}
	return result.String()
}

// toSnakeCase converts CamelCase to snake_case (e.g. "GetProfile" -> "get_profile")
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}

// generateRandomPassword creates a random string of a given length.
func generateRandomPassword(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(passwordCharset))))
		if err != nil {
			return "", err
		}
		b[i] = passwordCharset[num.Int64()]
	}
	return string(b), nil
}
