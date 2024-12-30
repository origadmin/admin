/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"errors"
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
	"google.golang.org/protobuf/types/known/timestamppb"

	"origadmin/application/admin/helpers/id"
	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
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
var ProviderSet = wire.NewSet(NewData, NewAuthRepo, NewCurrentRepo, NewMenuRepo, NewRoleRepo, NewUserRepo)

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
	return d, func() {
		log.Info("closing the data resources")
		if err := drv.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func (obj *Data) InitFromFile(ctx context.Context, filename string) error {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	var menus []*dto.MenuPB
	err = codec.DecodeFromFile(filename, &menus)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Warnw("Menu data file not found, skip init menu data from file", "file", abs)
			return nil
		}
		return err
	}
	for i, pb := range menus {
		log.Infow("msg", "Processing menu", "index", i, "menuId", pb.Id, "menuKeyword", pb.Keyword, "menuName", pb.Name)
		if pb.Children != nil {
			for i2, child := range pb.Children {
				log.Infow("msg", "Processing child", "index", i2, "childId", child.Id, "childKeyword", child.Keyword, "childName", child.Name)
			}
		}
	}
	return obj.Tx(ctx, func(ctx context.Context) error {
		return obj.createInBatchByParent(ctx, menus, nil)
	})
}

func (obj *Data) createInBatchByParent(ctx context.Context, items []*dto.MenuPB, parent *dto.MenuPB) error {
	total := len(items)
	log.Infow("msg", "Starting createInBatchByParent", "totalItems", total)

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
			exists, err := obj.Menu(ctx).Query().Where(menu.ID(item.Id)).Exist(ctx)
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
			var wheres = []predicate.Menu{
				menu.Keyword(item.Keyword),
			}
			if pid != 0 {
				wheres = append(wheres, menu.ParentID(pid))
			}
			exists, err := obj.Menu(ctx).Query().Where(wheres...).Exist(ctx)
			if err != nil {
				log.Errorw("msg", "Error checking item by Keyword", "itemKeyword", item.Keyword, "parentId", pid, "error", err)
				return err
			}
			if exists {
				menuItem, err := obj.Menu(ctx).Query().Where(wheres...).First(ctx)
				if err != nil {
					log.Errorw("msg", "Error fetching item by Keyword", "itemKeyword", item.Keyword, "parentId", pid, "error", err)
					return err
				}
				founded = true
				item.Id = menuItem.ID
				log.Infow("msg", "Item found by Keyword", "itemKeyword", item.Keyword, "itemId", item.Id)
			}
		case item.Name != "":
			log.Infow("msg", "Checking item by Name", "itemName", item.Name, "parentId", pid)
			var wheres = []predicate.Menu{
				menu.Name(item.Name),
			}
			if pid != 0 {
				wheres = append(wheres, menu.ParentID(pid))
			}
			exists, err := obj.Menu(ctx).Query().Where(wheres...).Exist(ctx)
			if err != nil {
				log.Errorw("msg", "Error checking item by Name", "itemName", item.Name, "parentId", pid, "error", err)
				return err
			}
			if exists {
				menuItem, err := obj.Menu(ctx).Query().Where(wheres...).First(ctx)
				if err != nil {
					log.Errorw("msg", "Error fetching item by Name", "itemName", item.Name, "parentId", pid, "error", err)
					return err
				}
				founded = true
				item.Id = menuItem.ID
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
				item.Status = dto.MenuStatusActivated
				log.Infow("msg", "Setting default status for item", "itemId", item.Id, "status", item.Status)
			}
			if item.Sequence == 0 {
				item.Sequence = int32(total - i)
				log.Infow("msg", "Setting default sequence for item", "itemId", item.Id, "sequence", item.Sequence)
			}

			item.ParentId = pid
			if parent != nil {
				item.ParentPath = parent.ParentPath + strconv.Itoa(int(pid)) + TreePathDelimiter
				log.Infow("msg", "Setting parent path for item", "itemId", item.Id, "parentPath", item.ParentPath)
			}
			item.CreateTime = timestamppb.New(time.Now())
			item.UpdateTime = timestamppb.New(time.Now())
			//columns := menu.OmitColumns()
			//if item.ParentId == "" {
			//	columns = menu.OmitColumns(menu.FieldParentID)
			//}
			if _, err := obj.Menu(ctx).Create().SetMenu(dto.MenuObject(item)).Save(ctx); err != nil {
				log.Errorw("msg", "Error creating menu item", "itemId", item.Id, "sequence", item.Sequence, "error", err)
				return err
			}
			log.Infow("msg", "Menu item created successfully", "itemId", item.Id)
		}

		for _, res := range item.Resources {
			log.Infow("msg", "Processing resource", "resourceId", res.Id, "resourcePath", res.Path, "resourceMethod", res.Method)

			if res.Id != 0 {
				log.Infow("msg", "Checking resource by ID", "resourceId", res.Id)
				exists, err := obj.Resource(ctx).Query().Where(resource.ID(res.Id)).Exist(ctx)
				if err != nil {
					log.Errorw("msg", "Error checking resource by ID", "resourceId", res.Id, "error", err)
					return err
				} else if exists {
					log.Infow("msg", "Resource already exists by ID", "resourceId", res.Id)
					continue
				}
			}
			if res.Path != "" {
				log.Infow("msg", "Checking resource by Path and Method", "resourcePath", res.Path, "resourceMethod", res.Method, "menuId", item.Id)

				exists, err := obj.Resource(ctx).Query().Where(resource.Path(res.Path), resource.Method(res.Method), resource.ID(res.Id)).Exist(ctx)
				if err != nil {
					log.Errorw("msg", "Error checking resource by Path and Method", "resourcePath", res.Path, "resourceMethod", res.Method, "resourceId", res.Id, "error", err)
					return err
				}
				if exists {
					log.Infow("msg", "Resource already exists by Path and Method", "resourcePath", res.Path, "resourceMethod", res.Method)
					continue
				}
			}
			if res.Id == 0 {
				res.Id = id.Gen()
				log.Infow("msg", "Generated new ID for resource", "resourceId", res.Id)
			}
			res.MenuId = item.Id
			res.CreateTime = timestamppb.New(time.Now())
			res.UpdateTime = timestamppb.New(time.Now())
			//columns := resource.OmitColumns()
			if _, err := obj.Resource(ctx).Create().SetResource(dto.ResourceObject(res)).SetID(res.Id).Save(ctx); err != nil {
				log.Errorw("msg", "Error creating resource", "resourceId", res.Id, "error", err)
				return err
			}
			log.Infow("msg", "Resource created successfully", "resourceId", res.Id)
		}

		if len(item.Children) != 0 {
			log.Infow("Processing children for item", "itemId", item.Id, "childCount", len(item.Children))
			if err := obj.createInBatchByParent(ctx, item.Children, item); err != nil {
				log.Errorw("Error processing children", "itemId", item.Id, "error", err)
				return err
			}
			log.Infow("msg", "Children processed successfully", "itemId", item.Id)
		}
	}

	log.Infow("msg", "Finished createInBatchByParent")
	return nil
}
