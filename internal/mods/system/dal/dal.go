/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/origadmin/contrib/database"
	"github.com/origadmin/entslog/v3"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/codec"
	"github.com/origadmin/toolkits/crypto/hash"

	"origadmin/application/admin/helpers/id"
	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"origadmin/application/admin/internal/mods/system/dto"
)

const (
	TreePathDelimiter = "."
)

// Data .
type Data struct {
	*ent.Database
}

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	wire.Struct(new(LoginData), "*"),
	NewData,
	NewAuthRepo,
	NewLoginRepo,
	NewPersonalRepo,
	NewMenuRepo,
	NewResourceRepo,
	NewRoleRepo,
	NewUserRepo,
	NewPermissionRepo,
	NewCasbinSourceRepo,
	RefreshTokenizer,
)

// NewTrans returns a transaction wit data
//func NewTrans(data *Data) database.Trans {
//	return data
//}

const FKSuffix = "_fk=1"

func FixSource(source string) string {
	// Check if the source already contains the FK parameter
	if strings.Contains(source, FKSuffix) {
		return source
	}

	// Check if the source already contains parameters
	if strings.Contains(source, "?") {
		// If parameters exist, append with &
		if !strings.HasSuffix(source, "&") {
			source += "&"
		}
		source += FKSuffix
	} else {
		// If no parameters exist, append with ?
		source += "?" + FKSuffix
	}
	return source
}

// NewData .
func NewData(bootstrap *configs.Bootstrap, logger log.KLogger) (*Data, func(), error) {
	hash.UseCrypto(hash.Type(bootstrap.GetCryptoType()))
	cfg := bootstrap.GetData().GetDatabase()
	if cfg == nil {
		return nil, nil, errors.New("data source not found")
	}
	if cfg.Dialect == "sqlite3" {
		cfg.Source = FixSource(cfg.GetSource())
	}
	sourceConfig, err := mysql.ParseDSN(cfg.Source)
	if err != nil {
		return nil, nil, err
	}
	log.Infow("msg", "connecting to database", "dialect", cfg.Dialect, "source", sourceConfig.Addr)
	drv, err := database.Open(cfg)
	log.Infow("msg", "connecting to database", "dialect", cfg.Dialect, "source", cfg.Source)
	if err != nil {
		log.Errorw("msg", "failed opening connection to database", "error", err)
		return nil, nil, err
	}

	//err = drv.Exec(context.Background(), "PRAGMA foreign_keys = ON;", []any{}, nil)
	//if err != nil {
	//	return nil, nil, err
	//}
	// Run the auto migration tool.
	debugDrv := entslog.New(sql.OpenDB(cfg.Dialect, drv))
	client := ent.NewClient(ent.Driver(debugDrv))
	if true || cfg.GetMigration().GetEnabled() {
		if err := client.Schema.Create(
			context.Background(),
			schema.WithDropIndex(true),
			schema.WithDropColumn(true),
			schema.WithForeignKeys(false)); err != nil {
			log.Errorw("msg", "failed creating schema resources", "error", err)
			return nil, nil, err
		}
	}

	d := &Data{
		Database: ent.NewDatabase(client),
	}
	// 初始化数据
	if err := d.InitDataFromPath(context.Background(), ""); err != nil {
		log.Errorw("failed to init data", "error", err)
		return nil, nil, err
	}

	return d, func() {
		log.Info("closing the data resources")
		if err := drv.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func NewDataWithClient(client *ent.Client) *Data {
	return &Data{
		Database: ent.NewDatabase(client),
	}
}

func (obj *Data) InitDataFromPath(ctx context.Context, path string, filters ...string) error {
	type data struct {
		name string
		fn   func(ctx context.Context, filename string) error
	}
	initializers := []data{
		{
			name: "resource",
			fn:   obj.InitResourceFromFile,
		},
		{
			name: "role",
			fn:   obj.InitRoleFromFile,
		},
		{
			name: "user",
			fn:   obj.InitUserFromFile,
		},
		{
			name: "department",
			fn:   obj.InitDepartmentFromFile,
		},
		{
			name: "position",
			fn:   obj.InitPositionFromFile,
		},
		{
			name: "permission",
			fn:   obj.InitPermissionFromFile,
		},
	}
	actions := make([]data, 0)
	for _, di := range initializers {
		for _, filter := range filters {
			if di.name == filter {
				actions = append(actions, di)
			}
		}

	}
	for _, action := range actions {
		action.name = filepath.Join(path, action.name+".json")
		err := action.fn(ctx, action.name)
		if err != nil {
			return err
		}
	}

	return nil
}
func (obj *Data) InitResourceFromFile(ctx context.Context, filename string) error {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	var resources []*dto.ResourcePB
	err = codec.DecodeFromFile(abs, &resources)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Warnw("Resource data file not found, skip init resource data from file", "file", abs)
			return nil
		}
		return err
	}
	for i, pb := range resources {
		log.Infow("msg", "Processing resource", "index", i, "resourceId", pb.Id, "resourceKeyword", pb.Keyword, "resourceName", pb.Name)
		if pb.Children != nil {
			for i2, child := range pb.Children {
				log.Infow("msg", "Processing child", "index", i2, "childId", child.Id, "childKeyword", child.Keyword, "childName", child.Name)
			}
		}
	}
	return obj.Tx(ctx, func(ctx context.Context) error {
		return obj.createResourceBatchWithParent(ctx, resources, nil)
	})
}

func (obj *Data) createResourceBatchWithParent(ctx context.Context, items []*dto.ResourcePB, parent *dto.ResourcePB) error {
	total := len(items)
	log.Infow("msg", "Starting createResourceBatchWithParent", "totalItems", total)

	for i, item := range items {
		log.Infow("msg", "Processing item", "index", i, "itemId", item.Id, "itemKeyword", item.Keyword, "itemName", item.Name)
		var pid int64
		if parent != nil {
			pid = parent.Id
			log.Infow("msg", "Parent ID set", "parentId", pid)
		}
		founded := false
		switch {
		case item.Id != 0:
			log.Infow("Checking item by ID", "itemId", item.Id)
			exists, err := obj.Resource(ctx).Query().Where(resource.ID(item.Id)).Exist(ctx)
			if err != nil {
				log.Errorw("msg", "Error checking item by ID", "itemId", item.Id, "error", err)
				return err
			}
			if exists {
				log.Infow("msg", "Item already exists by ID", "itemId", item.Id)
				continue
			}
		case item.Keyword != "":
			log.Infow("msg", "Checking item by Keyword", "itemKeyword", item.Keyword, "parentId", pid)
			var wheres = []predicate.Resource{
				resource.Keyword(item.Keyword),
			}
			if pid != 0 {
				wheres = append(wheres, resource.ParentID(pid))
			}
			exists, err := obj.Resource(ctx).Query().Where(wheres...).Exist(ctx)
			if err != nil {
				log.Errorw("msg", "Error checking item by Keyword", "itemKeyword", item.Keyword, "parentId", pid, "error", err)
				return err
			}
			if exists {
				resourceItem, err := obj.Resource(ctx).Query().Where(wheres...).First(ctx)
				if err != nil {
					log.Errorw("msg", "Error fetching item by Keyword", "itemKeyword", item.Keyword, "parentId", pid, "error", err)
					return err
				}
				founded = true
				item.Id = resourceItem.ID
				log.Infow("msg", "Item found by Keyword", "itemKeyword", item.Keyword, "itemId", item.Id)
			}
		case item.Name != "":
			log.Infow("msg", "Checking item by Name", "itemName", item.Name, "parentId", pid)
			var conditions = []predicate.Resource{
				resource.Name(item.Name),
			}
			if pid != 0 {
				conditions = append(conditions, resource.ParentID(pid))
			}
			exists, err := obj.Resource(ctx).Query().Where(conditions...).Exist(ctx)
			if err != nil {
				log.Errorw("msg", "Error checking item by Name", "itemName", item.Name, "parentId", pid, "error", err)
				return err
			}
			if exists {
				resourceItem, err := obj.Resource(ctx).Query().Where(conditions...).First(ctx)
				if err != nil {
					log.Errorw("msg", "Error fetching item by Name", "itemName", item.Name, "parentId", pid, "error", err)
					return err
				}
				founded = true
				item.Id = resourceItem.ID
				log.Infow("msg", "Item found by Name", "itemName", item.Name, "itemId", item.Id)
			}
		default:
			log.Infow("msg", "No ID, Keyword, or Name provided for item")
		}

		if !founded {
			if item.Id == 0 {
				item.Id = id.Gen()
				log.Infow("msg", "Generated new ID for item", "itemId", item.Id)
			}
			if item.Status == 0 {
				item.Status = int32(dto.UserStatusActive)
				log.Infow("msg", "Setting default status for item", "itemId", item.Id, "status", item.Status)
			}
			if item.Sequence == 0 {
				item.Sequence = int32(total - i)
				log.Infow("msg", "Setting default sequence for item", "itemId", item.Id, "sequence", item.Sequence)
			}

			item.ParentId = pid
			if parent != nil {
				item.TreePath = parent.TreePath + strconv.Itoa(int(pid)) + TreePathDelimiter
				log.Infow("msg", "Setting parent path for item", "itemId", item.Id, "treePath", item.TreePath)
			}
			itemObj := dto.ConvertResourcePB2Object(item)
			itemObj.UpdateTime = time.Now()
			itemObj.CreateTime = time.Now()
			if _, err := obj.Resource(ctx).Create().SetResource(itemObj).Save(ctx); err != nil {
				log.Errorw("msg", "Error creating resource item", "itemId", item.Id, "sequence", item.Sequence, "error", err)
				return err
			}
			log.Infow("msg", "Resource item created successfully", "itemId", item.Id)
		}

		if len(item.Children) != 0 {
			log.Infow("Processing children for item", "itemId", item.Id, "childCount", len(item.Children))
			if err := obj.createResourceBatchWithParent(ctx, item.Children, item); err != nil {
				log.Errorw("Error processing children", "itemId", item.Id, "error", err)
				return err
			}
			log.Infow("msg", "Children processed successfully", "itemId", item.Id)
		}
	}
	log.Infow("msg", "Finished createResourceBatchWithParent")
	return nil
}

func (obj *Data) InitUserFromFile(ctx context.Context, filename string) error {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	var users []*dto.UserNode
	err = codec.DecodeFromFile(abs, &users)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Warnw("User data file not found, skip init user data from file", "file", abs)
			return nil
		}
		return err
	}
	return obj.Tx(ctx, func(ctx context.Context) error {
		return obj.createUserBatch(ctx, users)
	})
}

func (obj *Data) createUserBatch(ctx context.Context, users []*dto.UserNode) error {
	total := len(users)
	log.Infow("msg", "Starting createUserBatch", "totalItems", total)
	for i, item := range users {
		log.Infow("msg", "Processing item", "index", i, "itemId", item.Id, "itemUsername", item.Username, "itemNickname", item.Nickname)
		user, ps, err := dto.MakeCreateUser(&item.UserPB, item.Username, item.Password, dto.UserMutationOption{})
		if err != nil {
			return err
		}
		fmt.Println("generate user: ", user.Username, "with password: ", ps)
		if _, err := obj.User(ctx).Create().SetIsSystem(item.IsSystem).SetUser(dto.ConvertUserPB2Object(user)).
			Save(ctx); err != nil {
			log.Errorw("msg", "Error creating user item", "itemId", item.Id, "error", err)
			return err
		}
		log.Infow("msg", "User item created successfully", "itemId", item.Id, "itemUuid", item.Uuid)
	}
	log.Infow("msg", "Finished createUserBatch")
	return nil
}

func (obj *Data) InitRoleFromFile(ctx context.Context, filename string) error {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	var roles []*dto.RolePB
	err = codec.DecodeFromFile(abs, &roles)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Warnw("Role data file not found, skip init role data from file", "file", abs)
			return nil
		}
		return err
	}
	return obj.Tx(ctx, func(ctx context.Context) error {
		return obj.createRoleBatch(ctx, roles)
	})
}

func (obj *Data) createRoleBatch(ctx context.Context, roles []*dto.RolePB) error {
	total := len(roles)
	log.Infow("msg", "Starting createRoleBatch", "totalItems", total)
	for i, item := range roles {
		log.Infow("msg", "Processing item", "index", i, "itemId", item.Id, "itemKeyword", item.Keyword, "itemName", item.Name)
		if item.Id == 0 {
			item.Id = id.Gen()
			log.Infow("msg", "Generated new ID for item", "itemId", item.Id)
		}
		if _, err := obj.Role(ctx).Create().SetRole(dto.ConvertRolePB2Object(item)).Save(ctx); err != nil {
			log.Errorw("msg", "Error creating role item", "itemId", item.Id, "error", err)
			return err
		}
		log.Infow("msg", "Role item created successfully", "itemId", item.Id)
	}
	log.Infow("msg", "Finished createRoleBatch")
	return nil
}

func (obj *Data) InitDepartmentFromFile(ctx context.Context, filename string) error {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	var departments []*dto.DepartmentPB
	err = codec.DecodeFromFile(abs, &departments)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Warnw("Department data file not found, skip init department data from file", "file", abs)
			return nil
		}
		return err
	}
	return obj.Tx(ctx, func(ctx context.Context) error {
		return obj.createDepartmentBatch(ctx, departments, nil)
	})
}

func (obj *Data) createDepartmentBatch(ctx context.Context, departments []*dto.DepartmentPB, parent *dto.DepartmentPB) error {
	total := len(departments)
	log.Infow("msg", "Starting createDepartmentBatch", "totalItems", total)
	for i, item := range departments {
		log.Infow("msg", "Processing item", "index", i, "itemId", item.Id, "itemKeyword", item.Keyword, "itemName", item.Name)
		if item.Id == 0 {
			item.Id = id.Gen()
			log.Infow("msg", "Generated new ID for item", "itemId", item.Id)
		}
		if parent != nil {
			item.ParentId = parent.Id
			item.TreePath = parent.TreePath + strconv.Itoa(int(parent.Id)) + TreePathDelimiter
		}

		if _, err := obj.Department(ctx).Create().SetDepartment(dto.ConvertDepartmentPB2Object(item)).Save(ctx); err != nil {
			log.Errorw("msg", "Error creating department item", "itemId", item.Id, "error", err)
			return err
		}

		log.Infow("msg", "Department item created successfully", "itemId", item.Id)
		if len(item.Children) != 0 {
			log.Infow("Processing children for item", "itemId", item.Id, "childCount", len(item.Children))
			if err := obj.createDepartmentBatch(ctx, item.Children, item); err != nil {
				log.Errorw("Error processing children", "itemId", item.Id, "error", err)
				return err
			}
			log.Infow("msg", "Children processed successfully", "itemId", item.Id)
		}
	}
	log.Infow("msg", "Finished createDepartmentBatch")
	return nil
}

func (obj *Data) InitPositionFromFile(ctx context.Context, filename string) error {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	var positions []*dto.PositionNode
	err = codec.DecodeFromFile(abs, &positions)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Warnw("Position data file not found, skip init position data from file", "file", abs)
			return nil
		}
		return err
	}
	return obj.Tx(ctx, func(ctx context.Context) error {
		return obj.createPositionBatch(ctx, positions)
	})
}

func (obj *Data) createPositionBatch(ctx context.Context, positions []*dto.PositionNode) error {
	total := len(positions)
	log.Infow("msg", "Starting createPositionBatch", "totalItems", total)
	for i, item := range positions {
		log.Infow("msg", "Processing item", "index", i, "itemId", item.Id, "itemKeyword", item.Keyword, "itemName", item.Name)
		if item.Id == 0 {
			item.Id = id.Gen()
			log.Infow("msg", "Generated new ID for item", "itemId", item.Id)
		}
		dept, err := obj.Department(ctx).Query().Where(department.Keyword(item.DepartmentKeyword)).Only(ctx)
		if err != nil {
			return err
		}

		if _, err := obj.Position(ctx).Create().SetPosition(&dto.Position{
			ID:           item.Id,
			CreateTime:   time.Now(),
			UpdateTime:   time.Now(),
			Name:         item.Name,
			Keyword:      item.Keyword,
			Description:  item.Description,
			DepartmentID: dept.ID,
		}).Save(ctx); err != nil {
			log.Errorw("msg", "Error creating position item", "itemId", item.Id, "error", err)
			return err
		}
		log.Infow("msg", "Position item created successfully", "itemId", item.Id)
	}
	log.Infow("msg", "Finished createPositionBatch")
	return nil
}

func (obj *Data) InitPermissionFromFile(ctx context.Context, filename string) error {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	var permissions []*dto.PermissionPB
	err = codec.DecodeFromFile(abs, &permissions)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Warnw("Permission data file not found, skip init permission data from file", "file", abs)
			return nil
		}
		return err
	}
	return obj.Tx(ctx, func(ctx context.Context) error {
		return obj.createPermissionBatch(ctx, permissions)
	})
}

func (obj *Data) createPermissionBatch(ctx context.Context, permissions []*dto.PermissionPB) error {
	total := len(permissions)
	log.Infow("msg", "Starting createPermissionBatch", "totalItems", total)
	for i, item := range permissions {
		log.Infow("msg", "Processing item", "index", i, "itemId", item.Id, "itemKeyword", item.Keyword, "itemName", item.Name)
		if item.Id == 0 {
			item.Id = id.Gen()
			log.Infow("msg", "Generated new ID for item", "itemId", item.Id)
		}
		if _, err := obj.Permission(ctx).Create().SetPermission(dto.ConvertPermissionPB2Object(item)).Save(ctx); err != nil {
			log.Errorw("msg", "Error creating permission item", "itemId", item.Id, "error", err)
			return err
		}
		log.Infow("msg", "Permission item created successfully", "itemId", item.Id)
	}
	log.Infow("msg", "Finished createPermissionBatch")
	return nil
}
