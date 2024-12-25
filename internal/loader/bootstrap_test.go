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
	"reflect"
	"testing"
	"time"

	_ "github.com/origadmin/contrib/database"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/rand"
	"github.com/origadmin/toolkits/idgen/uuid"
	"google.golang.org/protobuf/encoding/protojson"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/system/dal"
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConf() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_InitFromFile(t *testing.T) {
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
				Bootstrap: DefaultBootstrap(),
			},
			args: args{
				filename: "../../resources/data/menu.json",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, cleanup, err := dal.NewData(tt.fields.Bootstrap, log.DefaultLogger)
			if err != nil {
				t.Errorf("NewData() error = %v", err)
				return
			}
			defer cleanup()
			if err := d.InitFromFile(context.Background(), tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("InitFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
