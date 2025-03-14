package service

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/casbin/casbin/v2"

	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
)

// CasbinRuleService 处理casbin规则的服务
type CasbinRuleService struct {
	client        *ent.Client
	enforcer      *casbin.Enforcer
	userRoleCache sync.Map // map[int64][]int64
}

// NewCasbinRuleService 创建casbin规则服务实例
func NewCasbinRuleService(client *ent.Client, enforcer *casbin.Enforcer) *CasbinRuleService {
	return &CasbinRuleService{
		client:   client,
		enforcer: enforcer,
	}
}

// LoadPolicyFromDB 从数据库加载策略
func (s *CasbinRuleService) LoadPolicyFromDB(ctx context.Context) error {
	// 清除现有策略
	s.enforcer.ClearPolicy()

	// 1. 加载权限-资源策略
	permissions, err := s.client.Permission.Query().
		WithResources().
		All(ctx)
	if err != nil {
		return fmt.Errorf("查询权限资源失败: %w", err)
	}

	// 添加权限-资源策略
	for _, p := range permissions {
		permID := strconv.FormatInt(p.ID, 10)
		for _, res := range p.Edges.Resources {
			// 添加权限策略规则: p, permission_id, resource_path, method
			_, err = s.enforcer.AddPolicy(permID, res.Path, res.Method)
			if err != nil {
				return fmt.Errorf("添加权限策略规则失败: %w", err)
			}
		}
	}

	// 2. 加载角色-权限关系
	roles, err := s.client.Role.Query().
		Where(role.StatusEQ(1)). // 只加载激活状态的角色
		WithPermissions().
		All(ctx)
	if err != nil {
		return fmt.Errorf("查询角色权限关系失败: %w", err)
	}

	// 添加角色-权限映射
	for _, r := range roles {
		roleID := strconv.FormatInt(r.ID, 10)
		for _, p := range r.Edges.Permissions {
			permID := strconv.FormatInt(p.ID, 10)
			// 添加角色-权限映射规则: g, role_id, permission_id, permission
			_, err = s.enforcer.AddNamedGroupingPolicy("g", roleID, permID, "permission")
			if err != nil {
				return fmt.Errorf("添加角色权限映射失败: %w", err)
			}
		}
	}

	// 3. 加载用户角色关系到缓存
	return s.LoadUserRoleCache(ctx)
}

// LoadUserRoleCache 加载用户角色关系到缓存
func (s *CasbinRuleService) LoadUserRoleCache(ctx context.Context) error {
	// 清除现有缓存
	s.userRoleCache = sync.Map{}

	// 查询所有激活用户的角色
	users, err := s.client.User.Query().
		Where(user.StatusEQ(1)). // 只加载激活状态的用户
		WithRoles(func(q *ent.RoleQuery) {
			q.Where(role.StatusEQ(1)) // 只加载激活状态的角色
		}).
		All(ctx)
	if err != nil {
		return fmt.Errorf("查询用户角色关系失败: %w", err)
	}

	// 缓存用户角色关系
	for _, u := range users {
		var roleIDs []int64
		for _, r := range u.Edges.Roles {
			roleIDs = append(roleIDs, r.ID)
		}
		if len(roleIDs) > 0 {
			s.userRoleCache.Store(u.ID, roleIDs)
		}
	}

	return nil
}

// GetUserRoles 从缓存获取用户角色
func (s *CasbinRuleService) GetUserRoles(userID int64) []int64 {
	if roles, ok := s.userRoleCache.Load(userID); ok {
		if roleIDs, ok := roles.([]int64); ok {
			return roleIDs
		}
	}
	return nil
}

// GetUserPermissions 获取用户的所有权限
func (s *CasbinRuleService) GetUserPermissions(ctx context.Context, userID int64) ([]string, error) {
	// 从缓存获取用户角色
	roleIDs := s.GetUserRoles(userID)
	if len(roleIDs) == 0 {
		return nil, nil
	}

	var permissions []string
	// 获取每个角色的权限
	for _, roleID := range roleIDs {
		// 获取角色的所有权限ID，使用命名更准确的方法
		permIDs, err := s.enforcer.GetFilteredGroupingPolicy(0, strconv.FormatInt(roleID, 10))
		if err != nil {
			return nil, fmt.Errorf("获取角色权限失败: %w", err)
		}
		// 获取每个权限的资源访问规则
		for _, policy := range permIDs {
			if len(policy) < 2 {
				continue
			}
			permID := policy[1] // 获取权限ID
			// 获取权限对应的资源访问规则
			resources, err := s.enforcer.GetFilteredPolicy(0, permID)
			if err != nil {
				return nil, fmt.Errorf("获取权限资源失败: %w", err)
			}
			for _, res := range resources {
				if len(res) >= 3 {
					perm := fmt.Sprintf("%s:%s", res[1], res[2]) // resource_path:method
					permissions = append(permissions, perm)
				}
			}
		}
	}

	return permissions, nil
}

// CheckPermission 检查用户是否有权限访问特定资源
func (s *CasbinRuleService) CheckPermission(ctx context.Context, userID int64, path, method string) (bool, error) {
	// 从缓存获取用户角色
	roleIDs := s.GetUserRoles(userID)
	if len(roleIDs) == 0 {
		return false, nil
	}

	// 检查用户的每个角色是否有权限
	for _, roleID := range roleIDs {
		// 获取角色的所有权限
		permIDs, err := s.enforcer.GetFilteredGroupingPolicy(0, strconv.FormatInt(roleID, 10))
		if err != nil {
			return false, fmt.Errorf("获取角色权限失败: %w", err)
		}
		// 检查每个权限是否可以访问请求的资源
		for _, policy := range permIDs {
			if len(policy) < 2 {
				continue
			}
			permID := policy[1]
			// 使用权限ID检查是否有访问权限
			hasPermission, err := s.enforcer.Enforce(permID, path, method)
			if err != nil {
				return false, fmt.Errorf("检查权限失败: %w", err)
			}
			if hasPermission {
				return true, nil
			}
		}
	}

	return false, nil
}

// ReloadPolicy 重新加载权限策略
func (s *CasbinRuleService) ReloadPolicy(ctx context.Context) error {
	return s.LoadPolicyFromDB(ctx)
}

// UpdateUserRoleCache 更新指定用户的角色缓存
func (s *CasbinRuleService) UpdateUserRoleCache(ctx context.Context, userID int64) error {
	// 查询用户最新角色
	userRoles, err := s.client.User.Query().
		Where(
			user.StatusEQ(1),
			user.ID(userID),
		).
		WithRoles(func(q *ent.RoleQuery) {
			q.Where(role.StatusEQ(1))
		}).
		First(ctx)
	if err != nil {
		return fmt.Errorf("查询用户角色失败: %w", err)
	}

	// 更新角色缓存
	var roleIDs []int64
	for _, r := range userRoles.Edges.Roles {
		roleIDs = append(roleIDs, r.ID)
	}

	if len(roleIDs) > 0 {
		s.userRoleCache.Store(userID, roleIDs)
	} else {
		s.userRoleCache.Delete(userID)
	}

	return nil
}
