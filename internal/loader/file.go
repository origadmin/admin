/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/goexts/generic/settings"
	"github.com/origadmin/contrib/replacer"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/config"
	"github.com/origadmin/runtime/config/file"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/codec"
	"github.com/origadmin/toolkits/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"origadmin/application/admin/internal/configs"
)

// SaveOption represents an option for saving configuration data.
type SaveOption = func(*protojson.MarshalOptions)

var (
	r = replacer.New(replacer.WithStart("${"), replacer.WithEnd("}"), replacer.WithSeparator(":"))
)

func init() {
	runtime.RegisterConfig("file", FileConfig(NewFileConfig))
}

// SaveConfig saves the configuration data to the specified file path.
func SaveConfig(path string, data any, opts ...SaveOption) error {
	if v, ok := data.(proto.Message); ok && strings.HasSuffix(path, ".json") {
		opt := settings.Apply(&protojson.MarshalOptions{
			Indent: " ",
		}, opts)
		bytes, err := opt.Marshal(v)
		if err != nil {
			return err
		}
		if err := os.WriteFile(path, bytes, 0644); err != nil {
			return err
		}
		return nil
	}
	if err := codec.EncodeToFile(path, data); err != nil {
		return err
	}
	return nil
}

func Replace(s []byte, envs map[string]string) []byte {
	return r.Replace(s, envs)
}

func ReplaceObject(s any, envs map[string]string) error {
	marshal, err := json.Marshal(s)
	if err != nil {
		return err
	}
	marshal = Replace(marshal, envs)
	return json.Unmarshal(marshal, s)
}

type FileConfig func(*configv1.SourceConfig, *config.Options) (config.KSource, error)

func (f FileConfig) NewSource(sourceConfig *configv1.SourceConfig, _ *config.Options) (config.KSource, error) {
	return f(sourceConfig, nil)
}

func NewFileConfig(sourceConfig *configv1.SourceConfig, _ *config.Options) (config.KSource, error) {
	cfg := sourceConfig.GetFile()
	if cfg == nil {
		return nil, config.ErrInvalidConfigType
	}
	var options []file.Option
	if len(cfg.Ignores) > 0 {
		options = append(options, file.WithIgnores(cfg.Ignores...))
	}
	v := new(configs.Bootstrap)
	options = append(options, file.WithFormatter(fileFormatter(v)))
	path, _ := filepath.Abs(cfg.Path)
	log.NewHelper(log.DefaultLogger).Infof("loading config from %s", path)
	return file.NewSource(cfg.Path, options...), nil
}

func fileFormatter(typo any) file.Formatter {
	return func(key string, value []byte) (*config.KKeyValue, error) {
		err := encoding.GetCodec(format(key)).Unmarshal(value, typo)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal config")
		}
		//if v, ok := typo.(proto.Message); ok {
		//	c := encoding.GetCodec("proto")
		//	marshal, err := c.Marshal(v)
		//	if err != nil {
		//		return nil, err
		//	}
		//	return &config.KKeyValue{
		//		Key:    key,
		//		Format: "proto",
		//		Value:  marshal,
		//	}, nil
		//}
		j := encoding.GetCodec("json")
		marshal, err := j.Marshal(typo)
		if err != nil {
			return nil, err
		}
		key = strings.TrimSuffix(key, filepath.Ext(key))
		return &config.KKeyValue{
			Key:    key + ".json",
			Format: "json",
			Value:  marshal,
		}, nil
	}
}

func format(name string) string {
	if p := strings.Split(name, "."); len(p) > 1 {
		return p[len(p)-1]
	}
	return ""
}
