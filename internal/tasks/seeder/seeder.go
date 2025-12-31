/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package seeder contains the implementation for database seeding tasks.
package seeder

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/features/system/biz"
)

// ProviderSet is for wire injection.
var ProviderSet = wire.NewSet(NewSeeder)

const (
	passwordCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	passwordLength  = 16
)

// Seeder is the container for initialization tasks.
type Seeder struct {
	userUseCase *biz.UserUseCase
	rootUserCfg *confpb.RootUser
	log         *log.Helper
}

// NewSeeder creates a new Seeder.
func NewSeeder(userUseCase *biz.UserUseCase, cfg *conf.Config, logger log.Logger) (*Seeder, error) {
	return &Seeder{
		userUseCase: userUseCase,
		rootUserCfg: cfg.RootUser(),
		log:         log.NewHelper(log.With(logger, "module", "seeder")),
	}, nil
}

// Run executes all seeding tasks.
func (s *Seeder) Run() error {
	if err := s.createRootUser(); err != nil {
		return err
	}
	// Add other seeding tasks here, e.g., s.createInitialMenus()
	return nil
}

// createRootUser creates the initial administrator user if it does not exist.
// If the password in the config is empty, a random one will be generated and printed.
func (s *Seeder) createRootUser() error {
	if s.rootUserCfg == nil || !s.rootUserCfg.Enabled {
		s.log.Info("Root user seeding is disabled in config.")
		return nil
	}

	ctx := context.Background()
	username := s.rootUserCfg.GetUsername()

	// 1. Check if the user already exists by trying to list them.
	listReq := &system.ListUsersRequest{
		Keyword:  username,
		PageSize: 1,
	}
	existingUsers, total, err := s.userUseCase.ListUsers(ctx, listReq)
	if err != nil {
		s.log.Errorf("Failed to check for root user: %v", err)
		return err
	}
	if total > 0 || len(existingUsers) > 0 {
		s.log.Infof("Root user '%s' already exists, skipping creation.", username)
		return nil
	}

	// 2. If not found, create the new user.
	s.log.Infof("Root user '%s' not found, creating...", username)

	password := s.rootUserCfg.GetPassword()
	if password == "" {
		var err error
		password, err = generateRandomPassword(passwordLength)
		if err != nil {
			s.log.Errorf("Failed to generate random password: %v", err)
			return err
		}
		s.log.Infof("Generated random password for user '%s': %s", username, password)
	}

	newUser := &types.User{
		Username: username,
		Nickname: s.rootUserCfg.GetNickname(),
		Email:    s.rootUserCfg.GetEmail(),
	}

	createdUser, err := s.userUseCase.CreateUser(ctx, newUser, password)
	if err != nil {
		s.log.Errorf("Failed to create root user '%s': %v", username, err)
		return err
	}

	s.log.Infof("Successfully created root user '%s' with ID: %d", createdUser.Username, createdUser.Id)
	return nil
}

// generateRandomPassword creates a random string of a given length.
func generateRandomPassword(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(passwordCharset))))
		if err != nil {
			return "", err
		}
		b[i] = passwordCharset[num.Int64()]
	}
	return string(b), nil
}
