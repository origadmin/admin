package biz

import (
	"context"
	"strconv"

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

// ListMyViews retrieves the view tree for the currently authenticated user, filtered by their permissions.
func (uc *MeUseCase) ListMyViews(ctx context.Context, p security.Principal, scope string) ([]*types.View, error) {
	// 1. Get all views for the given scope.
	allViews, err := uc.repo.GetAllViewsByScope(ctx, scope)
	if err != nil {
		return nil, err
	}

	// If the user is the system user, return all views without filtering.
	// Ensure SystemUserID is not 0 to prevent accidental privilege escalation for user ID 0.
	if data.SystemUserID != 0 {
		currentUserID, _ := strconv.ParseInt(p.GetID(), 10, 64)
		if currentUserID == data.SystemUserID {
			return allViews, nil
		}
	}

	// 2. Get all permission keywords for the current user.
	permissionKeywords, err := uc.repo.GetPermissionKeywordsByUserID(ctx, p.GetID())
	if err != nil {
		return nil, err
	}

	// 3. Create a set for quick lookup.
	userPermissions := make(map[string]struct{})
	for _, keyword := range permissionKeywords {
		userPermissions[keyword] = struct{}{}
	}

	// 4. Filter the view tree based on permissions.
	return uc.filterViews(allViews, userPermissions), nil
}

// filterViews recursively filters a slice of views and their children.
func (uc *MeUseCase) filterViews(views []*types.View, userPermissions map[string]struct{}) []*types.View {
	var filtered []*types.View

	for _, view := range views {
		// Recursively filter children first.
		if len(view.Children) > 0 {
			view.Children = uc.filterViews(view.Children, userPermissions)
		}

		// A view is kept if:
		// 1. It has visible children after filtering.
		// 2. Or, it's a view that doesn't require a specific permission (e.g., a group or a link).
		// 3. Or, the user has direct permission for this view's keyword.
		if len(view.Children) > 0 || view.Keyword == "" {
			filtered = append(filtered, view)
		} else {
			if _, hasPerm := userPermissions[view.Keyword]; hasPerm {
				filtered = append(filtered, view)
			}
		}
	}

	return filtered
}
