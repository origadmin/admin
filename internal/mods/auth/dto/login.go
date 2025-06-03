/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	"github.com/google/uuid"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/rand"

	pb "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/helpers/id"
)

type (
	CaptchaIDRequest     = pb.CaptchaIdRequest
	CaptchaIDResponse    = pb.CaptchaIdResponse
	CaptchaImageRequest  = pb.CaptchaImageRequest
	CaptchaImageResponse = pb.CaptchaImageResponse
	CaptchaAudioRequest  = pb.CaptchaAudioRequest
	CaptchaAudioResponse = pb.CaptchaAudioResponse
	CaptchaRequest       = pb.CaptchaRequest
	CaptchaResponse      = pb.CaptchaResponse
	CaptchaData          = pb.CaptchaData
	LoginRequest         = pb.LoginRequest
	LoginResponse        = pb.LoginResponse
	LogoutRequest        = pb.LogoutRequest
	LogoutResponse       = pb.LogoutResponse
	TokenRefreshRequest  = pb.TokenRefreshRequest
	TokenRefreshResponse = pb.TokenRefreshResponse
	RegisterRequest      = pb.RegisterRequest
	RegisterResponse     = pb.RegisterResponse
	CurrentUserRequest   = pb.CurrentUserRequest
	CurrentUserResponse  = pb.CurrentUserResponse
)

type LoginRepo interface {
	CaptchaID(ctx context.Context, in *CaptchaIDRequest) (*CaptchaIDResponse, error)
	CaptchaImage(ctx context.Context, id string, reload bool) (*CaptchaImageResponse, error)
	CaptchaAudio(ctx context.Context, id string, reload bool) (*CaptchaAudioResponse, error)
	Captcha(ctx context.Context, in *CaptchaRequest) (*CaptchaResponse, error)
	CurrentUser(ctx context.Context, in *CurrentUserRequest) (*CurrentUserResponse, error)
	Login(ctx context.Context, in *LoginRequest) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest) (*LogoutResponse, error)
	Register(ctx context.Context, in *RegisterRequest) (*RegisterResponse, error)
	TokenRefresh(ctx context.Context, in *TokenRefreshRequest) (*TokenRefreshResponse, error)
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

// MakeCreateUser functions are used to create new users
func MakeCreateUser(user *UserPB, username, password string, option UserMutationOption) (*UserPB, string, error) {
	log.Debugf("Creating user with options: %+v", option)
	if !option.NoPasswd {
		log.Debugf("NoPasswd is false, checking for RandomPasswd")
		if option.RandomPasswd && (user.Email != "" || user.Phone != "") {
			log.Debugf("RandomPasswd is true and user has email or phone, generating random password")
			password = rand.GenerateRandom(8)
			log.Debugf("Generated random password: %s", password)
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
	registerID := id.Gen()
	user.Id = registerID
	user.Uuid = uuid.Must(uuid.NewRandom()).String()
	user.Username = username
	user.Name = "user_" + random.RandString(8)
	user.Status = 1
	return user, password, nil
}

var random = rand.NewRand(rand.KindDigit | rand.KindLowerCase | rand.KindUpperCase)
