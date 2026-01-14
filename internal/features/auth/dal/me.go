/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/data/entity/ent/view"
	"origadmin/application/admin/internal/features/auth/dto"
	systemDto "origadmin/application/admin/internal/features/system/dto"
)

type MeRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewMeRepo .
func NewMeRepo(db *ent.Database, logger log.Logger) dto.MeRepo {
	return &MeRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

func (r *MeRepo) GetProfile(ctx context.Context, userID int64) (*types.User, error) {
	u, err := r.db.User(ctx).Query().Where(user.ID(userID)).Only(ctx)
	if err != nil {
		// CORRECTED: Check for a "not found" error and return a specific, application-level error.
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		// For all other errors, return them as is.
		return nil, err
	}
	return systemDto.ConvertUserToUserPB(u), nil
}

// GetAllViewsByScope retrieves all views for a given scope and builds a tree.
func (r *MeRepo) GetAllViewsByScope(ctx context.Context, scope string) ([]*types.View, error) {
	views, err := r.db.View(ctx).Query().
		Where(view.ScopeEQ(scope)).
		Order(ent.Asc(view.FieldSequence)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return buildViewTree(systemDto.ConvertViewsToViewsPB(views)), nil
}

// GetPermissionKeywordsByUserID retrieves all permission keywords for a user.
func (r *MeRepo) GetPermissionKeywordsByUserID(ctx context.Context, userID string) ([]string, error) {
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
