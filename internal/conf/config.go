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

const (
	// APIPrefix is the prefix for all API routes.
	APIPrefix = "/api/v1"
)

type Config struct {
	Bootstrap confpb.Bootstrap
}

// Data returns the data configuration.
func (c *Config) Data() *datav1.Data {
	return c.Bootstrap.GetData()
}

// Caches returns the caches configuration.
func (c *Config) Caches() *datav1.Caches {
	return c.Bootstrap.GetData().GetCaches()
}

// Databases returns the databases configuration.
func (c *Config) Databases() *datav1.Databases {
	return c.Bootstrap.GetData().GetDatabases()
}

// ObjectStores returns the object stores configuration.
func (c *Config) ObjectStores() *datav1.ObjectStores {
	return c.Bootstrap.GetData().GetObjectStores()
}

// DefaultDiscovery returns the default discovery name.
func (c *Config) DefaultDiscovery() string {
	return c.Bootstrap.GetDefaultDiscovery()
}

// Discoveries returns the discoveries configuration.
func (c *Config) Discoveries() *discoveryv1.Discoveries {
	return c.Bootstrap.GetDiscoveries()
}

// Logger returns the logger configuration.
func (c *Config) Logger() *loggerv1.Logger {
	return c.Bootstrap.GetLogger()
}

// Middlewares returns the middlewares configuration.
func (c *Config) Middlewares() *middlewarev1.Middlewares {
	return c.Bootstrap.GetMiddlewares()
}

// Servers returns the servers configuration.
func (c *Config) Servers() *transportv1.Servers {
	return c.Bootstrap.GetServers()
}

// Clients returns the clients configuration.
func (c *Config) Clients() *transportv1.Clients {
	return c.Bootstrap.GetClients()
}

// Captcha returns the captcha configuration.
func (c *Config) Captcha() *confpb.Captcha {
	return c.Bootstrap.GetCaptcha()
}

// RootUser returns the root user configuration.
func (c *Config) RootUser() *confpb.RootUser {
	return c.Bootstrap.GetRootUser()
}

// --- Runtime Interface Adapters (Deprecated: Use direct getters above) ---

func (c *Config) DecodeData() (*datav1.Data, error) {
	return c.Data(), nil
}

func (c *Config) DecodeCaches() (*datav1.Caches, error) {
	return c.Caches(), nil
}

func (c *Config) DecodeDatabases() (*datav1.Databases, error) {
	return c.Databases(), nil
}

func (c *Config) DecodeObjectStores() (*datav1.ObjectStores, error) {
	return c.ObjectStores(), nil
}

func (c *Config) DecodeDefaultDiscovery() (string, error) {
	return c.DefaultDiscovery(), nil
}

func (c *Config) DecodeDiscoveries() (*discoveryv1.Discoveries, error) {
	return c.Discoveries(), nil
}

func (c *Config) DecodeLogger() (*loggerv1.Logger, error) {
	return c.Logger(), nil
}

func (c *Config) DecodeMiddlewares() (*middlewarev1.Middlewares, error) {
	return c.Middlewares(), nil
}

func (c *Config) DecodeServers() (*transportv1.Servers, error) {
	return c.Servers(), nil
}

func (c *Config) DecodeClients() (*transportv1.Clients, error) {
	return c.Clients(), nil
}

func (c *Config) GetCaptcha() (*confpb.Captcha, error) {
	return c.Captcha(), nil
}

func (c *Config) GetRootUser() (*confpb.RootUser, error) {
	return c.RootUser(), nil
}

func (c *Config) GetBootstrap() *confpb.Bootstrap {
	return &c.Bootstrap
}

func (c *Config) DecodedConfig() any {
	return &c.Bootstrap
}

func (c *Config) Transform(config interfaces.ConfigLoader, sc interfaces.StructuredConfig) (interfaces.StructuredConfig, error) {
	err := config.Decode("", &c.Bootstrap)
	if err != nil {
		return nil, err
	}
	//var debugMap map[string]any
	//if err = config.Decode("", &debugMap); err != nil {
	//	return nil, err
	//}
	return c, nil
}

func New() bootstrap.ConfigTransformer {
	return &Config{}
}
