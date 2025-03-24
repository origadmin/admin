/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	pb "origadmin/application/admin/api/v1/services/system"
)

// CasbinServiceHTTPServer is a login service.
type CasbinServiceHTTPServer struct {
	pb.UnimplementedCasbinSourceServiceServer

	client pb.CasbinSourceServiceHTTPClient
}

// NewCasbinServiceHTTPServer new a login service.
func NewCasbinServiceHTTPServer(client pb.CasbinSourceServiceHTTPClient) *CasbinServiceHTTPServer {
	return &CasbinServiceHTTPServer{client: client}
}

// NewCasbinServiceHTTPServerPB new a login service.
func NewCasbinServiceHTTPServerPB(client pb.CasbinSourceServiceHTTPClient) pb.CasbinSourceServiceServer {
	return &CasbinServiceHTTPServer{client: client}
}

var _ pb.CasbinSourceServiceHTTPServer = (*CasbinServiceHTTPServer)(nil)
