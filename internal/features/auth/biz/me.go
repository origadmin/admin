package biz

import (
	"context"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/origadmin/contrib/security"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/features/auth/dto"
)

// MeUseCase is a user profile use case.
type MeUseCase struct {
	repo dto.MeRepo
	log  *log.Helper
}

// NewMeUseCase new a user profile use case.
func NewMeUseCase(repo dto.MeRepo, logger log.Logger) *MeUseCase {
	return &MeUseCase{repo: repo, log: log.NewHelper(logger)}
}

// GetProfile gets the user profile.
func (uc *MeUseCase) GetProfile(ctx context.Context, userID int64) (*types.User, error) {
	return uc.repo.GetProfile(ctx, userID)
}

// ListMyViews retrieves all views available to the current user in a flat list.
func (uc *MeUseCase) ListMyViews(ctx context.Context, p security.Principal) ([]*types.View, error) {
	// 1. Get all active views from the database in a flat list.
	allViews, err := uc.repo.ListActiveViews(ctx)
	if err != nil {
		return nil, err
	}

	// 2. Bypass filtering for system user or admin user.
	userID, _ := strconv.ParseInt(p.GetID(), 10, 64)
	if userID == data.SystemUserID {
		return allViews, nil
	}

	isSystemRole, err := uc.repo.HasSystemRole(ctx, userID)
	if err != nil {
		return nil, err
	}
	if isSystemRole {
		return allViews, nil
	}

	// 3. For regular users, get their specific permission keywords.
	permissionKeywords, err := uc.repo.GetPermissionKeywordsByUserID(ctx, p.GetID())
	if err != nil {
		return nil, err
	}
	userPermissions := make(map[string]struct{})
	for _, keyword := range permissionKeywords {
		userPermissions[keyword] = struct{}{}
	}

	// 4. First pass: Identify all views that should be denied due to lack of direct permission.
	deniedIDs := make(map[int64]struct{})
	for _, view := range allViews {
		if view.Keyword != "" {
			if _, hasPerm := userPermissions[view.Keyword]; !hasPerm {
				deniedIDs[view.Id] = struct{}{}
			}
		}
	}

	// 5. Second pass: Build the final list, filtering out denied views and all their descendants.
	var filteredViews []*types.View
	for _, view := range allViews {
		if _, isDenied := deniedIDs[view.Id]; isDenied {
			continue
		}

		isDescendantOfDenied := false
		for deniedID := range deniedIDs {
			if strings.Contains(view.TreePath, ","+strconv.FormatInt(deniedID, 10)+",") {
				isDescendantOfDenied = true
				break
			}
		}

		if !isDescendantOfDenied {
			filteredViews = append(filteredViews, view)
		}
	}

	return filteredViews, nil
}
