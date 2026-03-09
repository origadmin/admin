/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/casbin/casbin/v3/persist"
	"github.com/google/wire"

	"github.com/origadmin/contrib/security/authn/jwt"
	"github.com/origadmin/contrib/security/authz"
	securitycasbin "github.com/origadmin/contrib/security/authz/casbin"
	"github.com/origadmin/contrib/security/credential"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/helpers/comp"
)

// ProviderBackendSet provides backend-specific dependencies for Business Logic.
// It relies on engine to provide infrastructure instances.
var ProviderBackendSet = wire.NewSet(
	ProviderCommonSet,
	NewDebounceExecutor,
	ProvideWatcher,
	ProvidePublisher,
	ProvideAuthenticator,
	wire.Bind(new(credential.Creator), new(*jwt.Authenticator)),
	ProvideAuthorizer,
	wire.Bind(new(authz.Authorizer), new(*securitycasbin.Authorizer)),
	// Other business-logic providers go here...
)

// ProvideWatcher bridges the engine-managed Watcher to wire-based services.
func ProvideWatcher(app *runtime.App) (persist.Watcher, error) {
	return comp.GetDefault[persist.Watcher](app.Context(), app.Container().In(CategoryWatcher))
}

// ProvidePublisher bridges the engine-managed Publisher to wire-based services.
func ProvidePublisher(app *runtime.App) (message.Publisher, error) {
	return comp.GetDefault[message.Publisher](app.Context(), app.Container().In(CategoryPublisher))
}
