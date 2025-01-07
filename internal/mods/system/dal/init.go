package dal

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/rand"

	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
)

// InitData 初始化所有数据
func InitData(ctx context.Context, client *ent.Client) error {
	// 按依赖顺序初始化
	initializers := []struct {
		name string
		fn   func(context.Context, *ent.Client) error
	}{
		{"department", initDepartments},
		{"position", initPositions},
		{"role", initRoles},
		{"resource", initResources},
		{"permission", initPermissions},
		{"user", initUsers},
	}

	for _, init := range initializers {
		if err := init.fn(ctx, client); err != nil {
			return fmt.Errorf("init %s: %w", init.name, err)
		}
	}

	return nil
}

func initResources(ctx context.Context, client *ent.Client) error {
	return loadJSON("resource.json", client.Resource)
}

// loadJSON 通用的 JSON 加载函数
func loadJSON(filename string, v interface{}) error {
	path := filepath.Join("internal", "mods", "system", "dal", "entity", "ent", "schema", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file %s: %w", filename, err)
	}
	return json.Unmarshal(data, v)
}

// DepartmentNode 部门树节点
type DepartmentNode struct {
	Keyword     string           `json:"keyword"`
	Name        string           `json:"name"`
	Sequence    int              `json:"sequence"`
	Status      int8             `json:"status"`
	Level       int              `json:"level"`
	Description string           `json:"description"`
	Children    []DepartmentNode `json:"children"`
}

// initDepartments 初始化部门数据
func initDepartments(ctx context.Context, client *ent.Client) error {
	var nodes []DepartmentNode
	if err := loadJSON("department.json", &nodes); err != nil {
		return err
	}

	return createDepartmentTree(ctx, client, nodes, nil)
}

// createDepartmentTree 递归创建部门树
func createDepartmentTree(ctx context.Context, client *ent.Client, nodes []DepartmentNode, parent *ent.Department) error {
	for _, node := range nodes {
		// 检查部门是否已存在
		exists, err := client.Department.Query().
			Where(department.Keyword(node.Keyword)).
			Exist(ctx)
		if err != nil {
			return err
		}
		if exists {
			continue
		}

		// 创建部门
		dept, err := client.Department.Create().
			SetKeyword(node.Keyword).
			SetName(node.Name).
			SetSequence(node.Sequence).
			SetStatus(node.Status).
			SetLevel(node.Level).
			SetDescription(node.Description).
			SetParent(parent).
			Save(ctx)
		if err != nil {
			return err
		}

		// 递归创建子部门
		if len(node.Children) > 0 {
			if err := createDepartmentTree(ctx, client, node.Children, dept); err != nil {
				return err
			}
		}
	}
	return nil
}

// initPositions 初始化岗位数据
func initPositions(ctx context.Context, client *ent.Client) error {
	var positions []struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		DepartmentID string `json:"department_id"`
	}
	if err := loadJSON("position.json", &positions); err != nil {
		return err
	}

	for _, pos := range positions {
		// 查找对应部门
		dept, err := client.Department.Query().
			Where(department.Keyword(pos.DepartmentID)).
			Only(ctx)
		if err != nil {
			return fmt.Errorf("find department %s: %w", pos.DepartmentID, err)
		}

		// 创建岗位
		_, err = client.Position.Create().
			SetName(pos.Name).
			SetDescription(pos.Description).
			SetDepartment(dept).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("create position %s: %w", pos.Name, err)
		}
	}
	return nil
}

// initRoles 初始化角色数据
func initRoles(ctx context.Context, client *ent.Client) error {
	var roles []struct {
		Keyword     string `json:"keyword"`
		Name        string `json:"name"`
		Type        int    `json:"type"`
		Sequence    int    `json:"sequence"`
		Status      int8   `json:"status"`
		IsSystem    bool   `json:"is_system"`
		Description string `json:"description"`
	}
	if err := loadJSON("role.json", &roles); err != nil {
		return err
	}

	for _, role := range roles {
		_, err := client.Role.Create().
			SetKeyword(role.Keyword).
			SetName(role.Name).
			SetType(int8(role.Type)).
			SetSequence(role.Sequence).
			SetStatus(role.Status).
			SetIsSystem(role.IsSystem).
			SetDescription(role.Description).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("create role %s: %w", role.Name, err)
		}
	}
	return nil
}

// initPermissions 初始化权限数据
func initPermissions(ctx context.Context, client *ent.Client) error {
	var permissions []struct {
		Name        string   `json:"name"`
		Keyword     string   `json:"keyword"`
		Type        string   `json:"type"`
		Description string   `json:"description"`
		DataScope   string   `json:"data_scope"`
		Resources   []string `json:"resources"`
	}
	if err := loadJSON("permission.json", &permissions); err != nil {
		return err
	}

	for _, perm := range permissions {
		// 创建权限
		p, err := client.Permission.Create().
			SetName(perm.Name).
			SetKeyword(perm.Keyword).
			//SetType(perm.Type).
			SetDescription(perm.Description).
			SetDataScope(perm.DataScope).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("create permission %s: %w", perm.Name, err)
		}

		// 关联资源
		for _, resKey := range perm.Resources {
			res, err := client.Resource.Query().
				Where(resource.Keyword(resKey)).
				Only(ctx)
			if err != nil {
				return fmt.Errorf("find resource %s: %w", resKey, err)
			}
			if err := p.Update().AddResources(res).Exec(ctx); err != nil {
				return fmt.Errorf("link resource %s to permission %s: %w", resKey, perm.Name, err)
			}
		}
	}
	return nil
}

// initUsers 初始化用户数据
func initUsers(ctx context.Context, client *ent.Client) error {
	var users []struct {
		Username    string   `json:"username"`
		Nickname    string   `json:"nickname"`
		Password    string   `json:"password"`
		Email       string   `json:"email"`
		Phone       string   `json:"phone"`
		Status      int8     `json:"status"`
		Roles       []string `json:"roles"`
		Departments []string `json:"departments"`
		Positions   []string `json:"positions"`
	}
	if err := loadJSON("user.json", &users); err != nil {
		return err
	}

	for _, u := range users {
		salt := rand.GenerateSalt()
		passwd, err := hash.Generate(u.Password, salt)
		if err != nil {
			return err
		}
		// 创建用户
		user, err := client.User.Create().
			SetUsername(u.Username).
			SetNickname(u.Nickname).
			SetPassword(passwd). // 加密密码
			SetSalt(salt).
			SetEmail(u.Email).
			SetPhone(u.Phone).
			SetStatus(u.Status).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("create user %s: %w", u.Username, err)
		}

		// 关联角色
		for _, roleKey := range u.Roles {
			role, err := client.Role.Query().
				Where(role.Keyword(roleKey)).
				Only(ctx)
			if err != nil {
				return fmt.Errorf("find role %s: %w", roleKey, err)
			}
			if err := user.Update().AddRoles(role).Exec(ctx); err != nil {
				return fmt.Errorf("link role %s to user %s: %w", roleKey, u.Username, err)
			}
		}

		// 关联部门
		for _, deptKey := range u.Departments {
			dept, err := client.Department.Query().
				Where(department.Keyword(deptKey)).
				Only(ctx)
			if err != nil {
				return fmt.Errorf("find department %s: %w", deptKey, err)
			}
			if err := user.Update().AddDepartments(dept).Exec(ctx); err != nil {
				return fmt.Errorf("link department %s to user %s: %w", deptKey, u.Username, err)
			}
		}

		// 关联岗位
		for _, posName := range u.Positions {
			pos, err := client.Position.Query().
				Where(position.Name(posName)).
				Only(ctx)
			if err != nil {
				return fmt.Errorf("find position %s: %w", posName, err)
			}
			if err := user.Update().AddPositions(pos).Exec(ctx); err != nil {
				return fmt.Errorf("link position %s to user %s: %w", posName, u.Username, err)
			}
		}
	}
	return nil
}
