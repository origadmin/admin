/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"encoding/json"
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// ResourceServiceHookedBridge is a menu service.
type ResourceServiceHookedBridge struct {
	pb.UnimplementedResourceServiceHooked
	log *log.KHelper
}

func (h ResourceServiceHookedBridge) CompleteCreateResource(ctx transhttp.Context, request *pb.CreateResourceRequest, response *pb.CreateResourceResponse) error {
	marshal, err := json.Marshal(response.Resource)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func (h ResourceServiceHookedBridge) CompleteDeleteResource(ctx transhttp.Context, request *pb.DeleteResourceRequest, response *pb.DeleteResourceResponse) error {
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    nil,
	})
}

func (h ResourceServiceHookedBridge) CompleteGetResource(ctx transhttp.Context, request *pb.GetResourceRequest, response *pb.GetResourceResponse) error {
	marshal, err := json.Marshal(response.Resource)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func (h ResourceServiceHookedBridge) CompleteListResources(ctx transhttp.Context, request *pb.ListResourcesRequest, response *pb.ListResourcesResponse) error {
	marshal, err := json.Marshal(response.Resources)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourcePage{
		Success: true,
		Total:   response.GetTotalSize(),
		Data:    marshal,
		//Current:  request.GetCurrent(),
		//PageSize: nil,
		//Extra: "",
	})
}

func (h ResourceServiceHookedBridge) CompleteUpdateResource(ctx transhttp.Context, request *pb.UpdateResourceRequest, response *pb.UpdateResourceResponse) error {
	marshal, err := json.Marshal(response.Resource)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func NewResourceServiceHookedBridge(r runtime.Runtime, client pb.ResourceServiceHTTPServer) pb.ResourceServiceHookedBridger {
	return pb.WithResourceServiceHook(&ResourceServiceHookedBridge{
		log: log.NewHelper(r.WithLogger("module", "service/system")),
	})(client)
}

// NewResourceServiceBridge new a menu service.
func NewResourceServiceBridge(r runtime.Runtime, client *service.GRPCClient) pb.ResourceServiceServer {
	return pb.NewResourceServiceBridge(client)
}

func NewResourceServiceBridgeClient(r runtime.Runtime, clients map[string]*service.GRPCClient) pb.ResourceServiceServer {
	if c, ok := clients["system"]; ok {
		return pb.NewResourceServiceBridge(c)
	} else {
		return pb.UnimplementedResourceServiceServer{}
	}
}

// NewResourceServiceHTTPBridge new a menu service.
func NewResourceServiceHTTPBridge(r runtime.Runtime, client *service.HTTPClient) pb.ResourceServiceHTTPServer {
	return pb.NewResourceServiceHTTPBridge(client)
}

var _ pb.ResourceServiceHooker = (*ResourceServiceHookedBridge)(nil)
