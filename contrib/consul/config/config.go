/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package config

import (
	"encoding/json"

	"github.com/hashicorp/consul/api"
	"github.com/origadmin/toolkits/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/origadmin/runtime"
	configv1 "github.com/origadmin/runtime/api/gen/go/config/v1"
	"github.com/origadmin/runtime/config"
)

func init() {
	runtime.RegisterConfigFunc(Type, NewConsulConfig)
	runtime.RegisterConfigSync(Type, config.SyncFunc(SyncConfig))
}

// NewConsulConfig create a new consul config.
func NewConsulConfig(ccfg *configv1.SourceConfig, options *config.Options) (config.KSource, error) {
	consul := ccfg.GetConsul()
	if consul == nil {
		return nil, errors.New("consul config error")
	}

	cfg := api.DefaultConfig()
	cfg.Address = consul.Address
	cfg.Scheme = consul.Scheme

	apiClient, err := api.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "consul client error")
	}

	if consul.Path == "" {
		consul.Path = FileConfigPath(ccfg.Name, DefaultPathName)
	}

	source, err := New(apiClient, WithPath(consul.Path))
	if err != nil {
		return nil, errors.Wrap(err, "consul source error")
	}

	//var configSources = []config.KSource{source}
	//if ccfg.EnvPrefixes != nil {
	//	configSources = append(configSources, env.NewSource(ccfg.EnvPrefixes...))
	//}
	//
	//options.Sources = append(options.Sources, configSources...)
	//if options.Decoder != nil {
	//	options.ConfigOptions = append(options.ConfigOptions, config.WithDecoder(options.Decoder))
	//}
	return source, nil
}

func SyncConfig(ccfg *configv1.SourceConfig, k string, v any, options *config.Options) error {
	consul := ccfg.GetConsul()
	if consul == nil {
		return errors.New("consul config error")
	}

	cfg := api.DefaultConfig()
	cfg.Address = consul.Address
	cfg.Scheme = consul.Scheme
	apiClient, err := api.NewClient(cfg)
	if err != nil {
		return errors.Wrap(err, "consul client error")
	}

	if consul.Path == "" {
		consul.Path = FileConfigPath(ccfg.Name, DefaultPathName)
	}

	encode := marshalJSON
	if options.Encoder != nil {
		encode = options.Encoder
	}
	marshal, err := encode(v)
	if err != nil {
		return errors.Wrap(err, "marshal config error")
	}

	if _, err := apiClient.KV().Put(&api.KVPair{
		Key:   consul.Path,
		Value: marshal,
	}, nil); err != nil {
		return errors.Wrap(err, "consul put error")
	}
	return nil
}

func FileConfigPath(serviceName, filename string) string {
	return "/config/" + serviceName + "/" + filename
}
func marshalJSON(v any) ([]byte, error) {
	if data, ok := v.(proto.Message); ok {
		opt := protojson.MarshalOptions{
			EmitUnpopulated: true,
			Indent:          " ",
		}
		return opt.Marshal(data)
	}
	return json.Marshal(v)
}
