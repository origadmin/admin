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

	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/codec"
)

func TestNewFileSourceConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *configv1.SourceConfig
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				path: "resources/configs",
			},
			want: &configv1.SourceConfig{
				Type: "consul",
				File: &configv1.SourceConfig_File{
					Path: "resources/configs",
				},
				Consul: &configv1.SourceConfig_Consul{
					Address: "127.0.0.1:8500",
					Scheme:  "http",
				},
				Etcd: &configv1.SourceConfig_ETCD{
					Endpoints: []string{"127.0.0.1:2379"},
				},
				EnvArgs: map[string]string{
					"env": "dev",
				},
				EnvPrefixes: []string{"env"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &configv1.SourceConfig{
				Type: "consul",
				File: &configv1.SourceConfig_File{
					Path: "resources/configs",
				},
				Consul: &configv1.SourceConfig_Consul{
					Address: "127.0.0.1:8500",
					Scheme:  "http",
				},
				EnvArgs: map[string]string{
					"env": "dev",
				},
				EnvPrefixes: []string{"env"},
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileSourceConfig() = %v, want %v", got, tt.want)
			}
			wd, _ := os.Getwd()
			path := filepath.Join(wd, "../../resources/local2.toml")
			err := codec.EncodeToFile(path, got)
			if err != nil {
				t.Errorf("NewFileSourceConfig() = %v", err)
			}
		})
	}
}
