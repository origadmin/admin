/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	system "origadmin/application/admin/api/v1/system"
	"origadmin/application/admin/internal/features/system/biz"
)

func toRoleDO(dto *system.Role) *biz.Role {
	if dto == nil {
		return nil
	}
	return &biz.Role{
		ID:      int(dto.Id),
		Name:    dto.Name,
		Keyword: dto.Keyword,
	}
}

func toRoleDTO(do *biz.Role) *system.Role {
	if do == nil {
		return nil
	}
	return &system.Role{
		Id:      int32(do.ID),
		Name:    do.Name,
		Keyword: do.Keyword,
	}
}

func toRoleDTOs(dos []*biz.Role) []*system.Role {
	dtos := make([]*system.Role, len(dos))
	for i, do := range dos {
		dtos[i] = toRoleDTO(do)
	}
	return dtos
}

func toUserDO(dto *system.User) *biz.User {
	if dto == nil {
		return nil
	}
	return &biz.User{
		ID:       int(dto.Id),
		Username: dto.Username,
	}
}

func toUserDTO(do *biz.User) *system.User {
	if do == nil {
		return nil
	}
	return &system.User{
		Id:       int32(do.ID),
		Username: do.Username,
	}
}

func toUserDTOs(dos []*biz.User) []*system.User {
	dtos := make([]*system.User, len(dos))
	for i, do := range dos {
		dtos[i] = toUserDTO(do)
	}
	return dtos
}

func toResourceDO(dto *system.Resource) *biz.Resource {
	if dto == nil {
		return nil
	}
	return &biz.Resource{
		ID:       int(dto.Id),
		Name:     dto.Name,
		ParentID: int(dto.ParentId),
	}
}

func toResourceDTO(do *biz.Resource) *system.Resource {
	if do == nil {
		return nil
	}
	return &system.Resource{
		Id:       int32(do.ID),
		Name:     do.Name,
		ParentId: int32(do.ParentID),
	}
}

func toResourceDTOs(dos []*biz.Resource) []*system.Resource {
	dtos := make([]*system.Resource, len(dos))
	for i, do := range dos {
		dtos[i] = toResourceDTO(do)
	}
	return dtos
}

func toPermissionDO(dto *system.Permission) *biz.Permission {
	if dto == nil {
		return nil
	}
	return &biz.Permission{
		ID:   int(dto.Id),
		Name: dto.Name,
	}
}

func toPermissionDTO(do *biz.Permission) *system.Permission {
	if do == nil {
		return nil
	}
	return &system.Permission{
		Id:   int32(do.ID),
		Name: do.Name,
	}
}

func toPermissionDTOs(dos []*biz.Permission) []*system.Permission {
	dtos := make([]*system.Permission, len(dos))
	for i, do := range dos {
		dtos[i] = toPermissionDTO(do)
	}
	return dtos
}
