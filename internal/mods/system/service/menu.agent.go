/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// MenuAPIGINRPCService is a menu service.
type MenuAPIGINRPCService struct {
	resp.Response

	client pb.MenuAPIClient
}

func (s MenuAPIGINRPCService) CreateMenu(context transhttp.Context, request *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	response, err := s.client.CreateMenu(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Menu),
	})
	return nil, nil
}

func (s MenuAPIGINRPCService) DeleteMenu(context transhttp.Context, request *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error) {
	_, err := s.client.DeleteMenu(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Data{
		Success: true,
		Data:    nil,
	})
	return nil, nil
}

func (s MenuAPIGINRPCService) GetMenu(context transhttp.Context, request *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	response, err := s.client.GetMenu(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Menu),
	})
	return nil, nil
}

func (s MenuAPIGINRPCService) ListMenus(context transhttp.Context, request *pb.ListMenusRequest) (*pb.ListMenusResponse, error) {
	response, err := s.client.ListMenus(context, request)
	if err != nil {
		return nil, err
	}
	//bytes, err := protojson.Marshal(response)
	//if err != nil {
	//	return nil, err
	//}
	//var pbs structpb.Struct
	//_ = pbs.UnmarshalJSON(bytes)
	//marshal, err := json.Marshal(&pbs)
	//if err != nil {
	//	return nil, err
	//}
	//s.Bytes(context, http.StatusOK, bytes)

	s.JSON(context, http.StatusOK, &resp.Page{
		Success: true,
		Total:   response.TotalSize,
		Data:    resp.Proto2AnyPBArray(response.Menus...),
		//Struct: &pbs,
		//Struct:  &pbs,
	})
	return nil, nil
}

func (s MenuAPIGINRPCService) UpdateMenu(context transhttp.Context, request *pb.UpdateMenuRequest) (*pb.UpdateMenuResponse, error) {
	response, err := s.client.UpdateMenu(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Menu),
	})
	return nil, nil
}

// NewMenuAPIGINRPCService new a menu service.
func NewMenuAPIGINRPCService(client pb.MenuAPIClient) *MenuAPIGINRPCService {
	return &MenuAPIGINRPCService{client: client}
}

// NewMenuAPIAgent new a menu service.
func NewMenuAPIAgent(client pb.MenuAPIClient) pb.MenuAPIAgent {
	return &MenuAPIGINRPCService{client: client}
}
func NewMenuServerAgent(client *service.GRPCClient) pb.MenuAPIAgent {
	cli := pb.NewMenuAPIClient(client)
	return NewMenuAPIAgent(cli)
}

var _ pb.MenuAPIAgent = (*MenuAPIGINRPCService)(nil)
