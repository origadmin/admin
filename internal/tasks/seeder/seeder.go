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

	"github.com/origadmin/contrib/security"
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
	userUseCase     *biz.UserUseCase
	resourceUseCase *biz.ResourceUseCase
	viewUseCase     *biz.ViewUseCase
	rootUserCfg     *confpb.RootUser
	log             *log.Helper
}

// NewSeeder creates a new Seeder.
func NewSeeder(userUseCase *biz.UserUseCase, resourceUseCase *biz.ResourceUseCase, viewUseCase *biz.ViewUseCase, cfg *conf.Config, logger log.Logger) (*Seeder, error) {
	return &Seeder{
		userUseCase:     userUseCase,
		resourceUseCase: resourceUseCase,
		viewUseCase:     viewUseCase,
		rootUserCfg:     cfg.RootUser(),
		log:             log.NewHelper(log.With(logger, "module", "seeder")),
	}, nil
}

// Run executes all seeding tasks.
func (s *Seeder) Run() error {
	if err := s.createRootUser(); err != nil {
		return err
	}
	if err := s.createInitialResources(); err != nil {
		return err
	}
	if err := s.createInitialViews(); err != nil {
		return err
	}
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

func (s *Seeder) createInitialResources() error {
	ctx := context.Background()
	for _, policy := range security.RegisteredPolicies() {
		// Check if resource already exists by its operation, which should be unique.
		_, count, err := s.resourceUseCase.ListResources(ctx,
			&system.ListResourcesRequest{
				Operation: policy.ServiceMethod,
				OnlyCount: true,
			})
		if err == nil && count > 0 {
			s.log.Infof("Resource for operation '%s' already exists, skipping.", policy.ServiceMethod)
			continue
		}

		// If not exists, create it using the dedicated biz method.
		if _, err := s.resourceUseCase.CreateResourceFromPolicy(ctx, &policy); err != nil {
			s.log.Errorf("failed to create resource from policy '%s': %v", policy.ServiceMethod, err)
		} else {
			s.log.Infof("Successfully created resource from policy: %s", policy.ServiceMethod)
		}
	}
	return nil
}

func (s *Seeder) createInitialViews() error {
	ctx := context.Background()
	views := []*types.View{
		{Name: "Dashboard", Keyword: "dashboard", Path: "/dashboard", Component: "default"},
		{Name: "System", Keyword: "system", Path: "/system", Component: "default"},
	}

	for _, view := range views {
		_, _, err := s.viewUseCase.ListViews(ctx, &system.ListViewsRequest{Keyword: view.Keyword})
		if err == nil {
			s.log.Infof("View '%s' already exists, skipping.", view.Name)
			continue
		}
		if _, err := s.viewUseCase.CreateView(ctx, view); err != nil {
			s.log.Errorf("failed to create view %s: %v", view.Name, err)
		} else {
			s.log.Infof("Successfully created view: %s", view.Name)
		}
	}
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
