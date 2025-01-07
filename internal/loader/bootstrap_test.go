/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/origadmin/contrib/database"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/slog-kratos"
	"github.com/origadmin/toolkits/crypto/rand"
	"github.com/origadmin/toolkits/idgen/uuid"
	"google.golang.org/protobuf/encoding/protojson"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/system/dal"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	_ "origadmin/application/admin/internal/mods/system/dal/entity/ent/runtime"
)

const (
	testPath = "test"
)

var (
	key = rand.GenerateRandom(32)
)

func init() {
	_, err := os.Stat(testPath)
	if err != nil {
		os.MkdirAll(testPath, 0755)
	}
}

func TestSaveConfig(t *testing.T) {
	fmt.Println("unix:", time.Now().Unix())
	fmt.Println("unixmillis:", time.Now().UnixMilli())
	bootstrap := DefaultBootstrap()
	//bootstrap.Security.Authn.Jwt.SigningMethod = "HS256"
	bootstrap.Security.Authn.Jwt.SigningKey = key
	bootstrap.Middleware.Jwt.Config.Key = key
	bootstrap.Middleware.Jwt.Config.SigningMethod = "HS512"
	bootstrap.Service.Middleware.Jwt.Config.Key = key
	bootstrap.Service.Middleware.Jwt.Config.SigningMethod = "HS512"
	type args struct {
		path string
		conf *configs.Bootstrap
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				path: "test.toml",
				conf: bootstrap,
			},
		},
		{
			name: "test",
			args: args{
				path: "test.yml",
				conf: bootstrap,
			},
		},
		{
			name: "test",
			args: args{
				path: "test.json",
				conf: bootstrap,
			},
		},
		//{
		//	name: "test",
		//	args: args{
		//		path: "test.ini",
		//		conf: DefaultBootstrap(),
		//	},
		//},
		//{
		//	name: "test",
		//	args: args{
		//		path: "test.xml",
		//		conf: DefaultBootstrap(),
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveConfig(filepath.Join(testPath, tt.args.path), tt.args.conf); (err != nil) != tt.wantErr {
				t.Errorf("SaveConf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		opt := protojson.MarshalOptions{
			EmitUnpopulated: true,
			Indent:          " ",
		}
		bs, _ := opt.Marshal(DefaultBootstrap())
		_ = os.WriteFile("test.json", bs, os.ModePerm)
	}
}

func TestLoadConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *configs.Bootstrap
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				path: "test.toml",
			},
			want:    DefaultBootstrap(),
			wantErr: false,
		},
		{
			name: "test",
			args: args{
				path: "test.json",
			},
			want:    DefaultBootstrap(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadFileBootstrap(filepath.Join(testPath, tt.args.path))
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Mode != tt.want.Mode {
				t.Errorf("LoadConf() got = %v, want %v", got.Mode, tt.want.Mode)
			}
			if got.Name != tt.want.Name {
				t.Errorf("LoadConf() got = %v, want %v", got.Name, tt.want.Name)
			}
		})
	}
}

func TestData_InitResourceFromFile(t *testing.T) {
	log.SetLogger(slog.NewLogger())
	_ = uuid.UUID{}
	type fields struct {
		Bootstrap *configs.Bootstrap
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				//Bootstrap: DefaultBootstrap(),
			},
			args: args{
				filename: "../../resources/data",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.Bootstrap == nil {
				abs, err := filepath.Abs("../../resources/configs/system/bootstrap.toml")
				if err != nil {
					return
				}
				log.Infof("abs: %s", abs)
				bs, err := LoadFileBootstrap("../../resources/configs/system/bootstrap.toml")
				if err != nil {
					t.Fatal(err)
					return
				}
				tt.fields.Bootstrap = bs
			}

			d, cleanup, err := dal.NewData(tt.fields.Bootstrap, log.DefaultLogger)
			if err != nil {
				t.Errorf("NewData() error = %v", err)
				return
			}
			defer cleanup()
			if err := d.InitDataFromPath(context.Background(), tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("InitFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func setupTestDB(t *testing.T) *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Fatalf("failed opening connection to sqlite: %v", err)
	}
	t.Cleanup(func() {
		client.Close()
	})

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func createTestRole(t *testing.T, client *ent.Client) *ent.Role {
	r, err := client.Role.Create().
		SetName("test-role").
		SetDescription("Test role for testing").
		SetKeyword("test-role").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed creating role: %v", err)
	}
	return r
}

func createTestResource(t *testing.T, client *ent.Client) *ent.Resource {
	m, err := client.Resource.Create().
		SetName("test-resource").
		SetURI("/test").
		SetDescription("Test menu for testing").
		SetKeyword("test-resource").
		//SetI18nKey("test-resource").
		//SetOperation("test-resource").
		//SetMethod("GET").
		//SetComponent("test-resource").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed creating resource: %v", err)
	}
	return m
}

func createTestPermission(t *testing.T, client *ent.Client) *ent.Permission {
	p, err := client.Permission.Create().
		SetName("test-permission").
		SetDescription("Test permission for testing").
		SetKeyword("test-permission").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed creating permission: %v", err)
	}
	return p
}
func TestAddRole(t *testing.T) {
	client := setupTestDB(t)
	role := createTestRole(t, client)

	if role.Name != "test-role" {
		t.Errorf("Expected role name to be 'test-role', got '%s'", role.Name)
	}
}

func TestAddResource(t *testing.T) {
	client := setupTestDB(t)
	menu := createTestResource(t, client)

	if menu.Name != "test-menu" {
		t.Errorf("Expected menu name to be 'test-menu', got '%s'", menu.Name)
	}
}

func TestAddPermission(t *testing.T) {
	client := setupTestDB(t)
	permission := createTestPermission(t, client)

	if permission.Name != "test-permission" {
		t.Errorf("Expected permission name to be 'test-permission', got '%s'", permission.Name)
	}
}

func TestAddRolePermission(t *testing.T) {
	client := setupTestDB(t)
	role := createTestRole(t, client)
	permission := createTestPermission(t, client)

	role, err := role.Update().
		//	AddRolePermissions(ent.RolePermission{
		//	ID:           0,
		//	RoleID:       0,
		//	PermissionID: 0,
		//	Edges:        ent.RolePermissionEdges{},
		//})
		AddPermissions(permission).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed adding permission to role: %v", err)
	}
	permissions := role.QueryPermissions().AllX(context.Background())
	if len(permissions) != 1 {
		t.Errorf("Expected role to have 1 permission, got %d", len(role.Edges.Permissions))
	}
}

func TestAddRoleResource(t *testing.T) {
	client := setupTestDB(t)
	role := createTestRole(t, client)
	perm := createTestPermission(t, client)
	res := createTestResource(t, client)
	pr, err := client.PermissionResource.Create().
		SetPermission(perm).
		SetResource(res).
		SetActions("abc123").Save(context.Background())
	//perm, err := perm.Update().AddPermissionResources().Save(context.Background())
	if err != nil {
		t.Fatalf("failed adding resource to permission: %v", err)
	}
	role, err = role.Update().
		AddPermissionIDs(pr.PermissionID).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed adding permission to role: %v", err)
	}
	ress := role.QueryPermissions().QueryResources().AllX(context.Background())
	if len(ress) != 1 {
		t.Errorf("Expected role to have 1 menu, got %d", len(ress))
	}
	resources := role.QueryPermissions().QueryPermissionResources().AllX(context.Background())
	t.Logf("resources: %v", resources)
}

func TestDeleteRole(t *testing.T) {
	client := setupTestDB(t)
	role := createTestRole(t, client)

	err := client.Role.DeleteOne(role).
		Exec(context.Background())
	if err != nil {
		t.Fatalf("failed deleting role: %v", err)
	}

	_, err = client.Role.Get(context.Background(), role.ID)
	if !ent.IsNotFound(err) {
		t.Errorf("Expected role to be deleted, got %v", err)
	}
}
