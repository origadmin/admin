// Package conf implements the functions, types, and interfaces for the module.
package conf

import (
	datav1 "github.com/origadmin/runtime/api/gen/go/config/data/v1"
	discoveryv1 "github.com/origadmin/runtime/api/gen/go/config/discovery/v1"
	loggerv1 "github.com/origadmin/runtime/api/gen/go/config/logger/v1"
	middlewarev1 "github.com/origadmin/runtime/api/gen/go/config/middleware/v1"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/interfaces"
	confpb "origadmin/application/admin/internal/conf/pb"
)

type Config struct {
	bootstrap confpb.Bootstrap
}

func (c *Config) DecodeData() (*datav1.Data, error) {
	return c.bootstrap.GetData(), nil
}

func (c *Config) DecodeCaches() (*datav1.Caches, error) {
	return c.bootstrap.GetData().GetCaches(), nil
}

func (c *Config) DecodeDatabases() (*datav1.Databases, error) {
	return c.bootstrap.GetData().GetDatabases(), nil
}

func (c *Config) DecodeObjectStores() (*datav1.ObjectStores, error) {
	return c.bootstrap.GetData().GetObjectStores(), nil
}

func (c *Config) DecodeDefaultDiscovery() (string, error) {
	return c.bootstrap.GetDefaultDiscovery(), nil
}

func (c *Config) DecodeDiscoveries() (*discoveryv1.Discoveries, error) {
	return c.bootstrap.GetDiscoveries(), nil
}

func (c *Config) DecodeLogger() (*loggerv1.Logger, error) {
	return c.bootstrap.GetLogger(), nil
}

func (c *Config) DecodeMiddlewares() (*middlewarev1.Middlewares, error) {
	return c.bootstrap.GetMiddlewares(), nil
}

func (c *Config) DecodeServers() (*transportv1.Servers, error) {
	return c.bootstrap.GetServers(), nil
}

func (c *Config) DecodeClients() (*transportv1.Clients, error) {
	return c.bootstrap.GetClients(), nil
}

func (c *Config) GetCaptcha() (*confpb.Captcha, error) {
	return c.bootstrap.GetCaptcha(), nil
}

func (c *Config) GetRootUser() (*confpb.RootUser, error) {
	return c.bootstrap.GetRootUser(), nil
}

func (c *Config) GetBootstrap() *confpb.Bootstrap {
	return &c.bootstrap
}

func (c *Config) DecodedConfig() any {
	return &c.bootstrap
}

func (c *Config) Transform(config interfaces.Config, config2 interfaces.StructuredConfig) (interfaces.
StructuredConfig, error) {
	return c, nil
}

func New() bootstrap.ConfigTransformer {
	return &Config{}
}
