/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package seeder contains the implementation for database seeding tasks.
package seeder

import (
	"cmp"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"unicode"

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
	"origadmin/application/admin/internal/helpers/contextutil"
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

	// 1. Check if user exists. If it does, we trust the infrastructure to have fixed its status.
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
		fmt.Printf("\n")
		fmt.Printf("\033[1;32m==================== [Root User Created] ===================\033[0m\n")
		fmt.Printf("\033[1;33mUsername: %s\033[0m\n", username)
		fmt.Printf("\033[1;33mPassword: %s\033[0m\n", password)
		fmt.Printf("\033[1;32m============================================================\033[0m\n")
		fmt.Printf("\n")
	}

	// 3. Create the user with System Context to ensure IsSystem=true in database.
	hashedPassword, err := s.hasher.Hash(password)
	if err != nil {
		return err
	}

	userIn := &types.User{
		Username: username,
		Email:    s.rootUserCfg.GetEmail(),
		Nickname: cmp.Or(s.rootUserCfg.GetNickname(), username),
		Status:   int32(enums.StatusEnabled),
	}

	systemCtx := contextutil.NewSystemUser(ctx)
	createdUser, err := s.userUseCase.CreateUser(systemCtx, userIn, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to create root user: %w", err)
	}

	if isRandom {
		s.log.Infof("Root user created successfully with random password.")
	} else {
		s.log.Infof("Root user '%s' created successfully.", username)
	}
	s.log.Infof("Successfully created root user with ID: %d", createdUser.Id)

	return nil
}

func (s *Seeder) createInitialResources() error {
	ctx := context.Background()
	ps := security.RegisteredPolicies()
	if len(ps) == 0 {
		return nil
	}

	seq := 1
	for _, policy := range ps {
		parts := strings.Split(policy.ServiceMethod, "/")
		if len(parts) < 3 {
			continue
		}
		fullService := parts[1]
		method := parts[2]

		serviceParts := strings.Split(fullService, ".")
		var moduleName string
		if len(serviceParts) >= 2 {
			moduleName = serviceParts[len(serviceParts)-2]
		} else {
			moduleName = "system"
		}

		serviceName := serviceParts[len(serviceParts)-1]
		resourceName := strings.TrimSuffix(serviceName, "Service")

		keyword := strings.Join([]string{strings.ToLower(moduleName), toSnakeCase(resourceName), toSnakeCase(method)}, ":")
		displayName := toTitleCase(resourceName) + " " + toTitleCase(method)
		i18nKey := "resource." + strings.ToLower(moduleName) + "." + toSnakeCase(resourceName) + "." + toSnakeCase(method)

		fullServiceName := moduleName
		if !strings.HasSuffix(fullServiceName, "-service") {
			fullServiceName += "-service"
		}

		existing, _, err := s.resourceUseCase.ListResources(ctx,
			&dto.ResourceQueryOption{
				Operation: policy.ServiceMethod,
			})

		if err == nil && len(existing) > 0 {
			res := existing[0]
			if res.VersionId == policy.VersionID && res.SyncStatus == "Synced" {
				continue
			}

			s.log.Infof("Updating existing resource '%s' (Version: %s -> %s)", keyword, res.VersionId, policy.VersionID)
			updateReq := &types.Resource{
				Id:          res.Id,
				Name:        displayName,
				Keyword:     keyword,
				I18N:        i18nKey,
				Operation:   policy.ServiceMethod,
				ServiceName: fullServiceName,
				VersionId:   policy.VersionID,
				SyncStatus:  "Synced",
				Sequence:    int32(seq),
			}
			if _, err := s.resourceUseCase.UpdateResource(ctx, updateReq); err != nil {
				s.log.Errorf("Failed to update resource from policy '%s': %v", policy.ServiceMethod, err)
			}
		} else {
			input := &dto.ResourceFromPolicyInput{
				Policy:      &policy,
				Keyword:     keyword,
				DisplayName: displayName,
				I18n:        i18nKey,
				Sequence:    seq,
				ServiceName: fullServiceName,
			}
			if _, err := s.resourceUseCase.CreateResourceFromPolicy(ctx, input); err != nil {
				s.log.Errorf("Failed to seed resource for policy %s/%s: %v", policy.Name, policy.ServiceMethod, err)
			}
		}
		seq++
	}
	return nil
}

func (s *Seeder) createInitialViews() error {
	jsonPath := filepath.Join("resources", "data", "views.json")
	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		return nil
	}
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("failed to read views.json: %w", err)
	}
	var views []*types.View
	if err := json.Unmarshal(data, &views); err != nil {
		return fmt.Errorf("failed to unmarshal views.json: %w", err)
	}
	ctx := context.Background()
	return s.createViewsRecursive(ctx, views, nil)
}

func (s *Seeder) createViewsRecursive(ctx context.Context, views []*types.View, parentID *int64) error {
	for _, view := range views {
		existing, total, err := s.viewUseCase.ListViews(ctx, &dto.ViewQueryOption{
			QueryOption: repo.QueryOption{
				Keyword:  view.Keyword,
				PageSize: 1,
			},
		})
		if err == nil && total > 0 && len(existing) > 0 {
			if len(view.Children) > 0 {
				existingID := existing[0].Id
				if err := s.createViewsRecursive(ctx, view.Children, &existingID); err != nil {
					return err
				}
			}
			continue
		}
		if parentID != nil {
			view.ParentId = *parentID
		}
		createdView, err := s.viewUseCase.CreateView(ctx, view)
		if err != nil {
			continue
		}
		if len(view.Children) > 0 {
			createdID := createdView.Id
			if err := s.createViewsRecursive(ctx, view.Children, &createdID); err != nil {
				return err
			}
		}
	}
	return nil
}

func toTitleCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune(' ')
		}
		result.WriteRune(r)
	}
	return result.String()
}

func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
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
