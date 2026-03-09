/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package seeder contains the implementation for database seeding tasks.
package seeder

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/origadmin/runtime/security"
	"github.com/origadmin/toolkits/crypto/hash"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/conf"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/repo"
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
	hasher          hash.Crypto
	rootUserCfg     *confpb.RootUser
	log             *log.Helper
}

// NewSeeder creates a new Seeder.
func NewSeeder(userUseCase *biz.UserUseCase, resourceUseCase *biz.ResourceUseCase, viewUseCase *biz.ViewUseCase, hasher hash.Crypto, cfg *conf.Config, logger log.Logger) (*Seeder, error) {
	return &Seeder{
		userUseCase:     userUseCase,
		resourceUseCase: resourceUseCase,
		viewUseCase:     viewUseCase,
		hasher:          hasher,
		rootUserCfg:     cfg.GetRootUser(),
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
func (s *Seeder) createRootUser() error {
	if s.rootUserCfg == nil || !s.rootUserCfg.Enabled {
		s.log.Info("Root user seeding is disabled in config.")
		return nil
	}

	ctx := context.Background()
	username := s.rootUserCfg.GetUsername()

	// 1. Check if the user already exists
	qo := &dto.UserQueryOption{
		QueryOption: repo.QueryOption{
			Keyword:  username,
			PageSize: 1,
		},
	}
	existingUsers, total, err := s.userUseCase.ListUsers(ctx, qo)
	if err != nil {
		s.log.Errorf("Failed to check for root user: %v", err)
		return err
	}
	if total > 0 || len(existingUsers) > 0 {
		s.log.Infof("Root user '%s' already exists, skipping creation.", username)
		return nil
	}

	// 2. Generate password
	password := s.rootUserCfg.GetPassword()
	isRandom := false
	if password == "" {
		password, _ = generateRandomPassword(passwordLength)
		isRandom = true
	}

	// 3. Create the user
	hashedPassword, _ := s.hasher.Hash(password)
	userIn := &types.User{
		Username: username,
		Nickname: "Administrator",
		Status:   int32(enums.StatusEnabled),
	}

	_, err = s.userUseCase.CreateUser(ctx, userIn, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to create root user: %w", err)
	}

	if isRandom {
		s.log.Infof("Root user created successfully.")
		s.log.Infof("Username: %s", username)
		s.log.Infof("Password: %s (PLEASE SAVE THIS PASSWORD!)", password)
	} else {
		s.log.Infof("Root user '%s' created successfully.", username)
	}

	return nil
}

func (s *Seeder) createInitialResources() error {
	ctx := context.Background()
	ps := security.RegisteredPolicies()
	if len(ps) == 0 {
		return nil
	}

	for _, p := range ps {
		policy := p
		// Generate a unique and descriptive keyword that satisfies constraints
		keyword := fmt.Sprintf("%s:%s", policy.Name, policy.ServiceMethod)

		input := &dto.ResourceFromPolicyInput{
			Policy:      &policy,
			Keyword:     keyword,
			DisplayName: policy.ServiceMethod,
			ServiceName: policy.Name,
		}
		_, err := s.resourceUseCase.CreateResourceFromPolicy(ctx, input)
		if err != nil {
			s.log.Errorf("Failed to seed resource for policy %s/%s: %v", p.Name, p.ServiceMethod, err)
		}
	}
	return nil
}

func (s *Seeder) createInitialViews() error {
	// TODO: Implement initial view seeding
	return nil
}

func generateRandomPassword(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(passwordCharset))))
		if err != nil {
			return "", err
		}
		result[i] = passwordCharset[num.Int64()]
	}
	return string(result), nil
}
