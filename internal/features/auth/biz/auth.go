package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/origadmin/contrib/security"
	"github.com/origadmin/toolkits/crypto/hash"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/auth/dto"
)

// AuthUseCase is a authentication use case.
type AuthUseCase struct {
	repo   dto.AuthRepo
	hasher hash.Crypto
	log    *log.Helper
}

// NewAuthUseCase new a authentication use case.
func NewAuthUseCase(repo dto.AuthRepo, hasher hash.Crypto, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		repo:   repo,
		hasher: hasher,
		log:    log.NewHelper(logger),
	}
}

// ListMyViews retrieves the view tree for the currently authenticated user, filtered by their permissions.
func (uc *AuthUseCase) ListMyViews(ctx context.Context, p security.Principal, scope string) ([]*types.View, error) {
	// 1. Get all views for the given scope.
	allViews, err := uc.repo.GetAllViewsByScope(ctx, scope)
	if err != nil {
		return nil, err
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
func (uc *AuthUseCase) filterViews(views []*types.View, userPermissions map[string]struct{}) []*types.View {
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

// VerifyUser verifies the user's credentials and returns the secure DTO if successful.
func (uc *AuthUseCase) VerifyUser(ctx context.Context, username, password string) (*types.User, error) {
	// 1. Get the internal AuthedUser DTO from the AuthRepo.
	authedUser, err := uc.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	// 2. Compare the provided password with the stored hash.
	if err := uc.hasher.Verify(authedUser.EncryptedPassword, password); err != nil {
		return nil, errors.New("invalid username or password")
	}

	// 3. On successful verification, return the safe User object from the DTO.
	return authedUser.User, nil
}

// UpdateLoginInfo delegates the update of login-related information to the repository.
func (uc *AuthUseCase) UpdateLoginInfo(ctx context.Context, userID int64, loginIP string) error {
	return uc.repo.UpdateLoginInfo(ctx, userID, loginIP)
}
