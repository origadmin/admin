/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package main

import (
	"github.com/google/wire"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func provideHasher() (hash.Crypto, error) {
	// Using a default cost for bcrypt. In a real application, this might come from config.
	return hash.NewCrypto(types.BCRYPT, bcrypt.WithCost(bcrypt.DefaultCost))
}

func provideLogger(app *runtime.App) log.Logger {
	return app.Logger()
}

var infraProviderSet = wire.NewSet(provideLogger, provideHasher)
