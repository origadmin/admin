/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"os"
	"path/filepath"

	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/goexts/generic/settings"
	"github.com/origadmin/contrib/config/envf"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/config"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/codec"
	"github.com/origadmin/toolkits/codec/json"
	"github.com/origadmin/toolkits/errors"

	"origadmin/application/admin/internal/configs"
)

type (
	Config    = configv1.SourceConfig
	Bootstrap = bootstrap.Bootstrap
)

func init() {
	runtime.RegisterConfigFunc("file", NewFileConfig)
}

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

func FromLocal(source *configv1.SourceConfig) (*configs.Bootstrap, error) {
	if source.File == nil {
		return nil, errors.String("file config is nil")
	}
	path := WorkPath("", source.File.Path)
	stat, err := os.Stat(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to state file %s", path)
	}
	if stat.IsDir() {
		return nil, errors.String("file config is a directory")
	}
	log.Infof("loading config from %s", path)
	return LoadFileBootstrap(path)
}

func FromLocalPath(path string, ss ...ConfigSetting) (*configs.Bootstrap, error) {
	source := NewFileSource(path)
	return FromLocal(settings.Apply(source, ss))
}

func NewFileConfig(cfg *Config, ss ...config.SourceSetting) (config.Config, error) {
	var configOption config.Option
	if cfg.EnvArgs != nil {
		configOption = config.WithSource(
			file.NewSource(cfg.File.Path),
			envf.WithEnv(cfg.EnvArgs, cfg.EnvPrefixes...),
		)
	} else {
		configOption = config.WithSource(
			file.NewSource(cfg.File.Path),
		)
	}
	opt := settings.Apply(&config.SourceOption{
		Options: []config.Option{configOption},
	}, ss)
	return config.New(opt.Options...), nil
}

func NewFileSource(path string) *Config {
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
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	return string(bytes)
}

func SourcePath(typo string, name, filename string) string {
	switch typo {
	case "file":
		return WorkPath("", filename)
	case "consul":
		return ConsulConfigPath(name, filename)
	//case "etcd":
	//	return source.ConfigPath(serviceName, "bootstrap.json")
	//default:
	//	return source.ConfigPath(serviceName, "bootstrap.json")
	default:

	}
	return filename
}

func ConsulConfigPath(name, filename string) string {
	return "/config/" + name + "/" + filename
}
