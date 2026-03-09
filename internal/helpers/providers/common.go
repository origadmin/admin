/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/contracts/component"
	"github.com/origadmin/runtime/data/storage"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"origadmin/application/admin/internal/conf"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/helpers/captcha"
)

// ProviderCommonSet provides common dependencies for all modules.
var ProviderCommonSet = wire.NewSet(
	ProvideConfig,
	ProvideLogger,
	ProvideHasher,
	ProvideCaptcha,
	ProvideServers,
)

// ProvideConfig bridges the PB-based Bootstrap config to the internal business Config.
func ProvideConfig(b *confpb.Bootstrap) *conf.Config {
	return &conf.Config{Bootstrap: *b}
}

// ProvideLogger provides the project's runtime logger.
func ProvideLogger(app *runtime.App) log.Logger {
	return app.Logger()
}

// ProvideServers extracts the server configuration from bootstrap.
func ProvideServers(cfg *confpb.Bootstrap) *transportv1.Servers {
	return cfg.GetServers()
}

// ProvideCache provides a cache provider from the runtime App.
func ProvideCache(app *runtime.App) storage.Provider {
	return storage.NewProvider(app.Container().In(""))
}

// ProvideHasher provides a password hasher.
func ProvideHasher() (hash.Crypto, error) {
	return hash.NewCrypto(types.BCRYPT, bcrypt.WithCost(bcrypt.DefaultCost))
}

// ProvideCaptcha provides a captcha instance from the runtime container.
func ProvideCaptcha(app *runtime.App) (*captcha.Captcha, error) {
	return comp.GetDefault[*captcha.Captcha](app.Context(), app.Container().In(component.CategorySecurity))
}
