/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	"github.com/goexts/generic/must"
	"github.com/google/uuid"

	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/rand"
	"github.com/origadmin/toolkits/identifier"
	"origadmin/application/admin/internal/helpers/pagination" // Corrected import path

	pb "origadmin/application/admin/api/v1/services/system"
)

type (
	ListUsersRequest  = pb.ListUsersRequest
	ListUsersResponse = pb.ListUsersResponse
)

type UserNode struct {
	UserPB
	IsSystem          bool     `json:"is_system"`
	RoleKeywords      []string `json:"role_keywords"`
	EncryptedPassword string   `json:"encrypted_password"`
}

// UserRepo is a UserPB repository interface.
type UserRepo interface {
	Get(context.Context, int64, ...UserQueryOption) (*UserPB, error)
	Create(context.Context, *UserPB, ...UserMutationOption) (*UserPB, error)
	Delete(context.Context, int64) error
	Update(context.Context, *UserPB, ...UserMutationOption) (*UserPB, error)
	List(context.Context, *ListUsersRequest, ...UserQueryOption) ([]*UserPB, int32, error)
	AddRoleIDs(context.Context, int64, []int64, ...UserMutationOption) error
	GetByUsername(context.Context, string, ...string) (*UserNode, error)
	GetRoleIDs(context.Context, int64) ([]int64, error)
	ListResourceByUserID(context.Context, int64, ...UserQueryOption) ([]*ResourcePB, error)
	Current(context.Context, int64) (*UserPB, error)
	UpdateUserStatus(ctx context.Context, id int64, status int8, options ...UserQueryOption) error
}

type UserMutationOption struct {
	RandomPasswd bool
	NoPasswd     bool
	Fields       []string
}

type UserQueryOption struct {
	IncludeRoles bool
	IsSystem     bool
	NoPasswd     bool
	RandomPasswd bool
	Status       int8 `form:"status" json:"status,omitempty"`
	SelectFields []string
	OmitFields   []string
	OrderFields  []string
	Fields       []string
}

func (o *UserQueryOption) FromListRequest(in *ListUsersRequest, limiter pagination.PageLimiter) error { // Updated usage
	in.Current = limiter.Current(in.Current)
	in.PageSize = limiter.PerPage(in.PageSize)
	return nil
}

func (o *UserQueryOption) FromGetRequest(in *pb.GetUserRequest, limiter pagination.PageLimiter) error { // Updated usage
	return nil
}

func (o *UserMutationOption) FromCreateRequest(in *pb.CreateUserRequest, limiter pagination.PageLimiter) error { // Updated usage
	o.RandomPasswd = in.RandomPassword
	return nil
}

// MakeCreateUser functions are used to create new users
func MakeCreateUser(user *UserPB, username, password string, option UserMutationOption) (*UserPB, string, error) {
	log.Debugf("Creating user with options: %+v", option)
	if !option.NoPasswd {
		log.Debugf("NoPasswd is false, checking for RandomPasswd")
		if option.RandomPasswd && (user.Email != "" || user.Phone != "") {
			log.Debugf("RandomPasswd is true and user has email or phone, generating random password")
			pwd, err := rand.RandomString(8)
			if err != nil {
				log.Errorf("Error generating random password: %v", err)
				return nil, "", err
			}
			log.Debugf("Generated random password: %s", pwd)
			password = pwd
		} else {
			log.Debugf("RandomPasswd is false or user has no email or phone")
		}
	} else {
		log.Debugf("NoPasswd is true, setting password to empty string")
		password = ""
	}
	var err error
	if password != "" {
		log.Debugf("Password is not empty, generating salt")
		//user.Salt = rand.GenerateSalt()
		//log.Debugf("Generated salt: %s", user.Salt)
		user.Password, err = hash.Generate(password)
		if err != nil {
			log.Errorf("Error generating password hash: %v", err)
			return nil, "", err
		}
		log.Debugf("Generated password hash: %s", user.Password)
	}
	registerID := identifier.GenerateNumber()
	user.Id = registerID
	user.Uuid = uuid.Must(uuid.NewRandom()).String()
	user.Username = username
	user.Name = "user_" + must.Do(rand.RandomString(8))
	user.Status = 1
	return user, password, nil
}
