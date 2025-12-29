package dal

import "github.com/google/wire"

// ProviderSet is dal providers.
var ProviderSet = wire.NewSet(NewAuthRepo, NewMeRepo, NewCaptchaRepo)
