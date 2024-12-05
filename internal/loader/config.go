/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"os"
	"path/filepath"

	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/goexts/generic/settings"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/config"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/codec"
	"github.com/origadmin/toolkits/codec/json"
	"github.com/origadmin/toolkits/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"origadmin/application/admin/internal/configs"
)

type (
	Config    = configv1.SourceConfig
	Bootstrap = bootstrap.Bootstrap
)

type ConfigSetting = func(config *Config)

func WithPath(name, filename string) ConfigSetting {
	return func(config *Config) {
		config.Name = name
		switch config.Type {
		case "file":
			config.File.Path = WorkPath("", filename)
		case "consul":
			config.Consul.Path = ConsulConfigPath(name, filename)
		}
	}
}

func WithServiceName(name string) ConfigSetting {
	return func(config *Config) {
		config.Name = name
	}
}

// LoadEnvFiles Loads configuration files in various formats from a directory,
func LoadEnvFiles(paths ...string) (map[string]string, error) {
	envs := make(map[string]string)
	for i := range paths {
		if err := filepath.WalkDir(paths[i], func(walkpath string, d os.DirEntry, err error) error {
			if err != nil {
				return errors.Wrapf(err, "failed to get config file %s", walkpath)
			} else if d.IsDir() {
				return nil
			}
			typo := codec.TypeFromExt(filepath.Ext(walkpath))
			if typo == codec.UNKNOWN {
				return nil
			}
			if err := codec.DecodeFromFile(walkpath, &envs); err != nil {
				return errors.Wrapf(err, "failed to parse config file %s", walkpath)
			}
			return nil
		}); err != nil {
			return nil, err
		}
	}

	return envs, nil
}

func FromLocalPath(path string, ss ...ConfigSetting) (*configs.Bootstrap, error) {
	source := FileSourceConfig(path)
	return LoadLocalBootstrap(settings.Apply(source, ss))
}

func NewFileConfig(cfg *Config, rc *config.RuntimeConfig) (config.Config, error) {
	var sources = []config.Source{file.NewSource(cfg.File.Path)}
	if cfg.EnvPrefixes != nil {
		sources = append(sources, env.NewSource(cfg.EnvPrefixes...))
		SetupEnv(cfg.EnvArgs, cfg.EnvPrefixes[0])
	}

	source := rc.Source()
	source.Options = append(source.Options, config.WithSource(sources...))
	return config.New(source.Options...), nil
}

func FileSourceConfig(path string) *Config {
	return &Config{
		Type: "file",
		File: &configv1.SourceConfig_File{
			Path: path,
		},
	}
}

func WorkPath(wd, path string) string {
	if wd != "" && !filepath.IsAbs(path) {
		path = filepath.Join(wd, path)
	}
	path, _ = filepath.Abs(path)
	return path
}

func PrintString(v any) string {
	if message, ok := v.(proto.Message); ok {
		option := protojson.MarshalOptions{
			Indent:            " ",
			EmitDefaultValues: false,
			EmitUnpopulated:   false,
		}
		bytes, _ := option.Marshal(message)
		return string(bytes)
	}

	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	return string(bytes)
}

func ConsulConfigPath(name, filename string) string {
	return "/config/" + name + "/" + filename
}
