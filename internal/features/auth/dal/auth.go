package dal

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/data/entity/ent/view"
	"origadmin/application/admin/internal/features/auth/dto"
)

type AuthRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(db *ent.Database, logger log.Logger) dto.AuthRepo {
	return &AuthRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

// GetUserByUsername retrieves a user's auth-specific data by their username.
func (r *AuthRepo) GetUserByUsername(ctx context.Context, username string) (*dto.AuthedUser, error) {
	u, err := r.db.User(ctx).Query().Where(user.UsernameEQ(username)).WithRoles().Only(ctx)
	if err != nil {
		return nil, err
	}
	return &dto.AuthedUser{
		User:              dto.ConvertUserToUserPB(u),
		EncryptedPassword: u.EncryptedPassword,
	}, nil
}

// UpdateLoginInfo updates the last login time, current login time, and last login IP for a user.
func (r *AuthRepo) UpdateLoginInfo(ctx context.Context, userID int64, loginIP string) error {
	// First, get the current user entity to perform the "shift change".
	currentUser, err := r.db.User(ctx).Get(ctx, userID)
	if err != nil {
		return err
	}

	// Perform the "shift change" and update to the new values.
	return r.db.User(ctx).
		UpdateOneID(userID).
		SetLastLoginTime(currentUser.LoginTime). // Previous login time becomes the last login time
		SetLoginTime(time.Now()).               // Set current login time
		SetLastLoginIP(currentUser.LoginIP).    // Previous login IP becomes the last login IP
		SetLoginIP(loginIP).                    // Set current login IP
		Exec(ctx)
}

// GetAllViewsByScope retrieves all views for a given scope and builds a tree.
func (r *AuthRepo) GetAllViewsByScope(ctx context.Context, scope string) ([]*types.View, error) {
	views, err := r.db.View(ctx).Query().
		Where(view.ScopeEQ(scope)).
		Order(ent.Asc(view.FieldSequence)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return buildViewTree(dto.ConvertViewsToViewsPB(views)), nil
}

// GetPermissionKeywordsByUserID retrieves all permission keywords for a user.
func (r *AuthRepo) GetPermissionKeywordsByUserID(ctx context.Context, userID string) ([]string, error) {
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, err
	}

	permissions, err := r.db.User(ctx).Query().
		Where(user.ID(id)).
		QueryRoles().
		QueryPermissions().
		All(ctx)
	if err != nil {
		return nil, err
	}

	keywords := make([]string, len(permissions))
	for i, p := range permissions {
		keywords[i] = p.Keyword
	}
	return keywords, nil
}

// buildViewTree converts a flat list of views into a tree structure.
func buildViewTree(views []*types.View) []*types.View {
	nodeMap := make(map[int64]*types.View)
	for _, v := range views {
		nodeMap[v.Id] = v
	}

	var tree []*types.View
	for _, v := range views {
		if parent, ok := nodeMap[v.ParentId]; ok {
			parent.Children = append(parent.Children, v)
		} else {
			// If a view has no parent in the map, it's a root node.
			tree = append(tree, v)
		}
	}
	return tree
}
