/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goexts/generic/settings"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/bootstrap"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/codec"
	"github.com/origadmin/toolkits/errors"

	"origadmin/application/admin/internal/configs"
)

// decodeFile loads the config file from the given path
func decodeFile(path string, cfg any) error {
	// skip stupid temp file
	if strings.HasSuffix(path, "~") || strings.HasSuffix(path, ".bak") || strings.HasSuffix(path, ".tmp") || strings.HasSuffix(path, ".lock") {
		return nil
	}
	// Decode the file into the config struct
	if err := codec.DecodeFromFile(path, cfg); err != nil {
		return errors.Wrapf(err, "failed to parse config file %s", path)
	}
	return nil
}

// decodeDir loads the config file from the given directory
func decodeDir(path string, cfg any) error {
	found := false
	// Walk through the directory and load each file
	err := filepath.WalkDir(path, func(walkpath string, d os.DirEntry, err error) error {
		if err != nil {
			return errors.Wrapf(err, "failed to get config file %s", walkpath)
		}
		// Check if the path is a directory
		if d.IsDir() {
			return nil
		}

		// Decode the file into the config struct
		if err := decodeFile(walkpath, cfg); err != nil {
			return err
		}
		found = true
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "load config error")
	}
	if !found {
		return errors.New("no config file found in " + path)
	}
	return nil
}

// LoadFileBootstrap load config from file
func LoadFileBootstrap(path string) (*configs.Bootstrap, error) {
	typo := codec.TypeFromPath(path)
	if typo == codec.UNKNOWN {
		return nil, fmt.Errorf("unknown file type: %s", path)
	}

	cfg := DefaultBootstrap()
	err := decodeFile(path, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
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
	var bs configs.Bootstrap
	if stat.IsDir() {
		err := decodeDir(path, &bs)
		if err != nil {
			return nil, err
		}
		return &bs, nil
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
	cfg := DefaultBootstrap()
	if err := config.Scan(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
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

	log.Infof("load source config: %+v", sourceConfig)
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
	log.Infof("load config: %+v\n", bs)
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
