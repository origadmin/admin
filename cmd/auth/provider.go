/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package main

import (
	"github.com/google/wire"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
)

func provideLogger(app *runtime.App) log.Logger {
	return app.Logger()
}

// infraProviderSet provides basic infrastructure dependencies.
var infraProviderSet = wire.NewSet(provideLogger)
