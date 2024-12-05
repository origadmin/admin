/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"fmt"
	"os"

	"github.com/goexts/generic/settings"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/bootstrap"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/codec"
	"github.com/origadmin/toolkits/errors"

	"origadmin/application/admin/internal/configs"
)

// LoadFileBootstrap load config from file
func LoadFileBootstrap(path string) (*configs.Bootstrap, error) {
	typo := codec.TypeFromPath(path)
	if typo == codec.UNKNOWN {
		return nil, fmt.Errorf("unknown file type: %s", path)
	}

	var cfg configs.Bootstrap
	err := codec.DecodeFromFile(path, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LoadLocalBootstrap(source *configv1.SourceConfig) (*configs.Bootstrap, error) {
	if source.File == nil {
		return nil, errors.String("file config is nil")
	}
	path := WorkPath("", source.File.Path)
	log.Infof("loading config from %s", path)
	stat, err := os.Stat(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to state file %s", path)
	}
	if stat.IsDir() {
		return nil, errors.String("file config is a directory")
	}
	return LoadFileBootstrap(path)
}

// LoadRemoteBootstrap load config from source
func LoadRemoteBootstrap(source *configv1.SourceConfig) (*configs.Bootstrap, error) {
	config, err := runtime.NewConfig(source)
	if err != nil {
		return nil, err
	}
	if err := config.Load(); err != nil {
		return nil, err
	}
	var cfg configs.Bootstrap
	if err := config.Scan(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// LoadBootstrap load config from file
func LoadBootstrap(flags *Bootstrap) (*configs.Bootstrap, error) {
	fmt.Println("load config from: ", flags.WorkPath())
	sourceConfig, err := bootstrap.LoadSourceConfig(flags)
	if err != nil {
		return nil, err
	}
	if len(sourceConfig.EnvPrefixes) > 0 {
		SetupEnv(sourceConfig.EnvArgs, sourceConfig.EnvPrefixes[0])
	}

	log.Infof("load soure config: %+v", sourceConfig)
	var bs *configs.Bootstrap
	switch sourceConfig.GetType() {
	case "file":
		bs, err = LoadLocalBootstrap(sourceConfig)
	default:
		bs, err = LoadRemoteBootstrap(sourceConfig)
	}
	if err != nil {
		return nil, err
	}
	if err := ReplaceObject(bs, sourceConfig.EnvArgs); err != nil {
		return nil, err
	}
	return bs, nil
}

// LoadRemoteServiceBootstrap get the configuration from the remote Configuration Center
func LoadRemoteServiceBootstrap(flags *Bootstrap, name string, ss ...ConfigSetting) (*configs.Bootstrap, error) {
	var cfg configs.Bootstrap
	sourceConfig, err := bootstrap.LoadSourceConfig(flags)
	if err != nil {
		return nil, err
	}

	sourceConfig.Name = name
	sourceConfig = settings.Apply(sourceConfig, ss)
	config, err := runtime.NewConfig(sourceConfig)
	if err != nil {
		return nil, err
	}
	if err := config.Load(); err != nil {
		return nil, err
	}
	if err := config.Scan(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
