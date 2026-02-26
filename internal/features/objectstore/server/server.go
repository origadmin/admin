/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	stdhttp "net/http"
	"os"
	"path/filepath"

	kratosgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	kratoshttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/origadmin/runtime"
	grpcv1 "github.com/origadmin/runtime/api/gen/go/config/transport/grpc/v1"
	httpv1 "github.com/origadmin/runtime/api/gen/go/config/transport/http/v1"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/container"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service/transport"
	runtimegrpc "github.com/origadmin/runtime/service/transport/grpc"
	runtimehttp "github.com/origadmin/runtime/service/transport/http"

	objPb "origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/internal/features/objectstore/dal"
	objSvc "origadmin/application/admin/internal/features/objectstore/service"
)

var ProviderSet = wire.NewSet(NewServers)

func NewServers(app *runtime.App, cfg *transportv1.Servers, objectStoreSvc *objSvc.ObjectStoreService, middlewareProvider container.ServerMiddlewareProvider, storageCfg *dal.LocalStorageConfig) ([]transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("servers config is nil")
	}
	var transportServers []transport.Server
	for _, serverCfg := range cfg.GetConfigs() {
		if serverCfg.GetName() != "objectstore" && serverCfg.GetName() != "origadmin.service.objectstore" {
			continue
		}
		switch serverCfg.GetProtocol() {
		case "http":
			srv, err := NewHTTPServer(app, serverCfg.GetHttp(), objectStoreSvc, middlewareProvider, storageCfg)
			if err == nil {
				transportServers = append(transportServers, srv)
			}
		case "grpc":
			srv, err := NewGRPCServer(app, serverCfg.GetGrpc(), objectStoreSvc, middlewareProvider)
			if err == nil {
				transportServers = append(transportServers, srv)
			}
		}
	}
	return transportServers, nil
}

func NewHTTPServer(_ *runtime.App, cfg *httpv1.Server, objectStoreSvc *objSvc.ObjectStoreService, provider container.ServerMiddlewareProvider, storageCfg *dal.LocalStorageConfig) (*transport.HTTPServer, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}
	mws, _ := provider.ServerMiddlewares()
	srv, _ := runtimehttp.NewServer(cfg, &runtimehttp.ServerOptions{ServerMiddlewares: mws})
	objPb.RegisterObjectStoreServiceHTTPServer(srv, objectStoreSvc)

	srv.Route("/").PUT("/multipart/{upload_id}/{part_number}", func(ctx kratoshttp.Context) error {
		uploadID, partNumber := ctx.Vars().Get("upload_id"), ctx.Vars().Get("part_number")
		r := ctx.Request()

		uploadPath := filepath.Join(storageCfg.MultipartPath, uploadID)
		_ = os.MkdirAll(uploadPath, 0755)

		file, err := os.Create(filepath.Join(uploadPath, partNumber))
		if err != nil {
			return err
		}
		defer file.Close()

		// 1. INITIALIZE MD5 ENGINE FOR PART VALIDATION
		partHasher := md5.New()
		writer := bufio.NewWriter(file)
		multiWriter := io.MultiWriter(writer, partHasher)

		var bodyReader io.Reader = r.Body
		if r.ContentLength > 0 {
			bodyReader = io.LimitReader(r.Body, r.ContentLength)
		}

		n, err := io.Copy(multiWriter, bodyReader)
		if err != nil {
			return err
		}
		writer.Flush()

		// 2. GENERATE REAL ETAG
		calculatedMD5 := hex.EncodeToString(partHasher.Sum(nil))

		// Optional: Verify against client-provided MD5 if present (X-Content-MD5)
		clientMD5 := r.Header.Get("X-Content-MD5")
		if clientMD5 != "" && clientMD5 != calculatedMD5 {
			file.Close()
			_ = os.Remove(filepath.Join(uploadPath, partNumber))
			return fmt.Errorf("data corruption: MD5 mismatch for part %s", partNumber)
		}

		log.Context(ctx).Infof("Saved part %s: %d bytes, ETag: %s", partNumber, n, calculatedMD5)
		ctx.Response().Header().Set("ETag", calculatedMD5)
		return ctx.Result(200, nil)
	})

	staticPath := "/objects/"
	srv.HandlePrefix(staticPath, stdhttp.StripPrefix(staticPath, stdhttp.FileServer(stdhttp.Dir(storageCfg.BasePath))))
	return srv, nil
}

func NewGRPCServer(_ *runtime.App, cfg *grpcv1.Server, objectStoreSvc *objSvc.ObjectStoreService, provider container.ServerMiddlewareProvider) (*transport.GRPCServer, error) {
	if cfg == nil {
		return nil, errors.New("grpc config is nil")
	}
	mws, _ := provider.ServerMiddlewares()
	srv, _ := runtimegrpc.NewServer(cfg, &runtimegrpc.ServerOptions{
		ServerOptions:     []kratosgrpc.ServerOption{kratosgrpc.Options(grpc.MaxRecvMsgSize(512 << 20))},
		ServerMiddlewares: mws,
	})
	objPb.RegisterObjectStoreServiceServer(srv, objectStoreSvc)
	return srv, nil
}
