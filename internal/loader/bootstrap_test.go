/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"origadmin/application/admin/internal/configs"
)

const (
	testPath = "test"
)

func init() {
	_, err := os.Stat(testPath)
	if err != nil {
		os.MkdirAll(testPath, 0755)
	}
}

func TestSaveConfig(t *testing.T) {
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
				conf: DefaultBootstrap(),
			},
		},
		{
			name: "test",
			args: args{
				path: "test.yml",
				conf: DefaultBootstrap(),
			},
		},
		{
			name: "test",
			args: args{
				path: "test.json",
				conf: DefaultBootstrap(),
			},
		},
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
