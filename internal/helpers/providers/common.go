/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/casbin/casbin/v3/persist"
	"github.com/google/wire"

	"github.com/origadmin/contrib/security/authn"
	"github.com/origadmin/contrib/security/authn/jwt"
	"github.com/origadmin/contrib/security/authz"
	securitycasbin "github.com/origadmin/contrib/security/authz/casbin"
	"github.com/origadmin/contrib/security/credential"
	"github.com/origadmin/runtime"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	hashtypes "github.com/origadmin/toolkits/crypto/hash/types"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/helpers/captcha"
	"origadmin/application/admin/internal/helpers/debounce"
	"origadmin/application/admin/internal/helpers/providers/internal/security"
)

// ProviderCommonSet provides common dependencies for all modules.
var ProviderCommonSet = wire.NewSet(
	ProvideLogger,
	ProvideHasher,
	ProvideServers,
)

// ProviderBackendSet provides backend dependencies.
var ProviderBackendSet = wire.NewSet(
	ProviderCommonSet,
	ProvideAuthenticator,
	wire.Bind(new(credential.Creator), new(*jwt.Authenticator)),
	wire.Bind(new(authn.Authenticator), new(*jwt.Authenticator)),
	ProvideAuthorizer,
	wire.Bind(new(authz.Authorizer), new(*securitycasbin.Authorizer)),
	ProvideDebouncer,
	ProvideWatcher,
	ProvidePublisher,
	ProvideCaptcha,
)

// ProviderSet combines all provider sets for backend.
var ProviderSet = wire.NewSet(
	ProviderBackendSet,
)

// ProvideServers extracts server configurations from the bootstrap config.
func ProvideServers(app *runtime.App) *transportv1.Servers {
	v, ok := app.Config().(*confpb.Bootstrap)
	if !ok {
		return nil
	}
	return v.GetServers()
}

// ProvideLogger provides the project's runtime logger.
func ProvideLogger(app *runtime.App) log.Logger {
	return app.Logger()
}

// ProvideHasher provides a hasher instance.
func ProvideHasher() (hash.Crypto, error) {
	return hash.NewCrypto(hashtypes.BCRYPT)
}

// ProvideDebouncer provides a debouncer instance.
func ProvideDebouncer(app *runtime.App) (debounce.Executor, error) {
	return comp.GetDefault[debounce.Executor](app.Context(), app.Container().In(debounce.CategoryDebounce))
}

// ProvideAuthenticator bridges the engine-managed Authenticator to wire.
func ProvideAuthenticator(app *runtime.App) (*jwt.Authenticator, error) {
	return comp.GetDefault[*jwt.Authenticator](app.Context(), app.Container().In(CategoryAuthn).WithInTags(GatewayTag))
}

// ProvideAuthorizer bridges the engine-managed Authorizer to wire.
func ProvideAuthorizer(app *runtime.App) (*securitycasbin.Authorizer, error) {
	return comp.GetDefault[*securitycasbin.Authorizer](app.Context(), app.Container().In(CategoryAuthz).WithInTags(
		FeatureTag))
}

// ProvideWatcher provides the watcher instance from the engine.
func ProvideWatcher(app *runtime.App) (persist.Watcher, error) {
	return comp.GetDefault[persist.Watcher](app.Context(), app.Container().In(CategoryWatcher))
}

// ProvidePublisher provides the publisher instance from the engine.
func ProvidePublisher(app *runtime.App) (message.Publisher, error) {
	return comp.GetDefault[message.Publisher](app.Context(), app.Container().In(CategoryPublisher))
}

func ProvideCaptcha(app *runtime.App) (*captcha.Captcha, error) {
	return comp.Get[*captcha.Captcha](app.Context(), app.Container().In(runtime.CategorySecurity), security.NameCaptcha)
}
