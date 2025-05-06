/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/bootstrap"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/interfaces/security"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/registry"
	"github.com/origadmin/runtime/service"

	"origadmin/application/admin/helpers/securityx"
	"origadmin/application/admin/internal/configs"
)

// AppOptions 包含微服务核心配置
type AppOptions struct {
	ID       string
	Name     string
	Version  string
	Metadata map[string]string
	Logger   log.KLogger
	Server   transport.Server // 改为通用传输层接口
}

var (
	ProviderSet = wire.NewSet(
		NewAuthConfig,
		NewRegistrar,
		NewTokenizer,
		NewAuthorizer,
		NewAuthenticator,
		wire.Struct(new(InjectorServer), "*"),
		wire.Struct(new(InjectorClient), "*"),
	)
)

var (
	_ *gins.Server
	_ *http.Server
	_ *grpc.Server
)

type Loader interface {
	SetupEnv() error
}

type InjectorClient struct {
	Logger    log.KLogger
	Bootstrap *configs.Bootstrap
	Server    *http.Server
}

type InjectorServer struct {
	Logger    log.KLogger
	Bootstrap *configs.Bootstrap
	Registrar registry.KRegistrar
	Servers   []transport.Server
}

func init() {
	runtime.RegisterConfigFunc("file", NewFileConfig)
	runtime.RegisterService("ORIGADMIN_SERVICE", service.DefaultServiceBuilder)
}

func NewAuthenticator(bootstrap *configs.Bootstrap) (security.Authenticator, error) {
	return securityx.NewAuthenticator(bootstrap)
}

func NewTokenizer(bootstrap *configs.Bootstrap) (security.Tokenizer, error) {
	authenticator, err := securityx.NewTokenizer(bootstrap)
	if err != nil {
		return nil, err
	}
	return authenticator, nil
}

func NewAuthorizer(bootstrap *configs.Bootstrap) (security.Authorizer, error) {
	return securityx.NewAuthorizer(bootstrap)
}

func NewAuthConfig(bootstrap *configs.Bootstrap) *configs.AuthConfig {
	// c := DefaultCaptcha()
	// todo Read from the configuration file
	return AuthConfig()
}

type loader struct {
	flags *bootstrap.Bootstrap
	cfg   *configv1.SourceConfig
}

func (l loader) SetupEnv() error {
	if len(l.cfg.EnvPrefixes) > 0 {
		SetupEnv(l.cfg.EnvArgs, l.cfg.EnvPrefixes[0])
	}
	return nil
}

func (l loader) Bootstrap() (*configs.Bootstrap, error) {
	var bs *configs.Bootstrap
	var err error
	switch l.cfg.GetType() {
	case "file":
		bs, err = LoadLocalBootstrap(l.cfg)
	default:
		bs, err = LoadRemoteBootstrap(l.cfg)
	}
	if err != nil {
		return nil, fmt.Errorf("load bootstrap error: %s", err.Error())
	}

	log.Infof("load config: %+v\n", bs)
	return bs, nil
}

func New(flags *Bootstrap) (Loader, error) {
	load := &loader{
		flags: flags,
	}
	sourceConfig, err := bootstrap.LoadSourceConfig(flags)
	if err != nil {
		return nil, err
	}
	load.cfg = sourceConfig
	return load, nil
}

// 删除复杂的多级错误收集机制，简化为优先级失败模式
func LoadBootstrap(cfg BootstrapConfig) (*configs.Bootstrap, error) {
	var bs *configs.Bootstrap
	var err error

	// 优先尝试加载远程配置 ✅ 首选远程配置中心
	bs, err = LoadRemoteBootstrap(&cfg.Source)
	if err == nil && bs != nil {
		if envErr := ReplaceObject(bs, cfg.Source.EnvArgs); envErr == nil {
			return bs, nil
		}
		return nil, fmt.Errorf("remote config replace failed: %w", envErr)
	}

	// 远程加载失败时回退到本地配置 ⛔️ 仅作为降级方案
	bs, err = LoadLocalBootstrap(&cfg.Source)
	if err == nil && bs != nil {
		if envErr := ReplaceObject(bs, cfg.Source.EnvArgs); envErr == nil {
			return bs, nil
		}
		return nil, fmt.Errorf("local config replace failed: %w", envErr)
	}

	return nil, fmt.Errorf("failed to load config: remote[%v], local[%v]",
		err, os.ErrNotExist)
}
