package broker

import "github.com/google/wire"

// ProviderSet is the provider set for the broker initializer package.
var ProviderSet = wire.NewSet(NewInitializer)
