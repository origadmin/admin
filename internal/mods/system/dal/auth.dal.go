/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"sync"

	"entgo.io/ent/dialect/sql"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/security"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/securityx"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"origadmin/application/admin/internal/mods/system/dto"
)

type authRepo struct {
	DB         *Data
	BufPool    *sync.Pool
	Tokenizer  security.Tokenizer
	Authorizer security.Authorizer
}

func (repo authRepo) CreateToken(ctx context.Context, request *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	claims := security.ClaimsFromContext(ctx)
	token, err := repo.Tokenizer.CreateToken(ctx, claims)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTokenResponse{
		Token: token,
	}, nil
}

func (repo authRepo) VerifyToken(ctx context.Context, request *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	valid, err := repo.Tokenizer.Verify(ctx, request.Token)
	if err != nil {
		return nil, err
	}
	return &pb.VerifyTokenResponse{
		IsValid: valid,
	}, nil
}

func (repo authRepo) DestroyToken(ctx context.Context, request *pb.DestroyTokenRequest) (*pb.DestroyTokenResponse, error) {
	err := repo.Tokenizer.DestroyToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}
	return &pb.DestroyTokenResponse{}, nil
}

func (repo authRepo) Authenticate(ctx context.Context, request *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	claims, err := repo.Tokenizer.Authenticate(ctx, request.Token)
	if err != nil {
		return nil, err
	}
	authorized, err := repo.Authorizer.Authorized(ctx, fromClaims(claims, request.Method, request.Path))
	if err != nil {
		return nil, err
	}
	return &pb.AuthenticateResponse{
		IsValid: authorized,
		//Claims: fromClaims(claims),
	}, nil
}

func (repo authRepo) ListAuthResources(ctx context.Context, in *dto.ListAuthResourcesRequest, options ...dto.AuthResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	var option dto.AuthResourceQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := repo.DB.Resource(ctx).Query()
	return authResourcePageQuery(ctx, query, in, option)
}

func fromClaims(claims security.Claims, method, path string) security.UserClaims {
	return &securityx.CasbinPolicy{
		Subject: claims.GetSubject(),
		Method:  method,
		Path:    path,
		Claims:  claims,
	}
}

// NewAuthRepo .
func NewAuthRepo(db *Data, logger log.KLogger) dto.AuthRepo {
	return &authRepo{
		DB:      db,
		BufPool: BufPool(),
	}
}

func authResourcePageQuery(ctx context.Context, query *ent.ResourceQuery, in *pb.ListAuthResourcesRequest, option dto.AuthResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	query = authResourceQueryPage(query, in)
	query = authResourceQueryOptions(query, option)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	result, err := query.All(ctx)
	return dto.ConvertResources(result), int32(count), err
}

func authResourceQueryPage(query *ent.ResourceQuery, in *pb.ListAuthResourcesRequest) *ent.ResourceQuery {
	if in.NoPaging {
		pageSize := in.PageSize
		if pageSize > 0 {
			query = query.Limit(int(pageSize))
		}
		return query
	}

	pageSize := in.PageSize
	if pageSize > 0 {
		query = query.Limit(int(pageSize))
	}
	current := in.Current
	if current > 0 {
		query = query.Offset(int((current - 1) * pageSize))
	}
	return query
}

func authResourceQueryOptions(query *ent.ResourceQuery, option dto.AuthResourceQueryOption) *ent.ResourceQuery {
	if len(option.SelectFields) > 0 {
		query = query.Select(option.SelectFields...).ResourceQuery
	}
	if len(option.OmitFields) > 0 {
		query = query.Omit(option.OmitFields...).ResourceQuery
	}
	if len(option.OrderFields) > 0 {
		query = query.Order(resourceOrderBy(option.OrderFields)...)
	}
	return query
}

func authResourceOrderBy(fields []string, opts ...sql.OrderTermOption) []resource.OrderOption {
	var orders []resource.OrderOption
	for _, field := range fields {
		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
	}
	return orders
}
