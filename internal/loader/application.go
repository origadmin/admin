/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

//
//func NewApp(ctx context.Context, injector *InjectorClient) *kratos.App {
//	opts := []kratos.Option{
//		kratos.ID(flags.ServiceID()),
//		kratos.Name(flags.ServiceName()),
//		kratos.Version(flags.Version()),
//		kratos.Metadata(map[string]string{}),
//		kratos.Context(ctx),
//		kratos.Signal(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT),
//		kratos.Logger(injector.Logger),
//		kratos.Server(injector.Server),
//	}
//
//	if flags.Env() == "release" {
//		gin.SetMode(gin.ReleaseMode)
//	}
//
//	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
//		log.Infow("msg", "GIN route", "method", httpMethod, "path", absolutePath, "operation", handlerName, "handlers", nuHandlers)
//	}
//
//	return kratos.New(opts...)
//}
