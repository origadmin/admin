// Copyright 2024 OrigAdmin. All rights reserved.

package fixtures

import (
	"path/filepath"
	"runtime"
)

// ProjectRoot returns the project root directory
func ProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../../..")
}

// CasbinModelPath returns the path to casbin model file
func CasbinModelPath() string {
	return filepath.Join(ProjectRoot(), "resources", "casbin_model.conf")
}

// TestDataPath returns the path to test data directory
func TestDataPath() string {
	return filepath.Join(ProjectRoot(), "tests", "fixtures", "data")
}

// BasePermissionsPath returns the path to base permissions YAML
func BasePermissionsPath() string {
	return filepath.Join(TestDataPath(), "base_permissions.yaml")
}

// BaseRolesPath returns the path to base roles YAML
func BaseRolesPath() string {
	return filepath.Join(TestDataPath(), "base_roles.yaml")
}

// CasbinPolicyPath returns the path to casbin policy CSV
func CasbinPolicyPath() string {
	return filepath.Join(TestDataPath(), "adapter", "rbac_policy.csv")
}

// CasbinModelFilePath returns the path to casbin model conf
func CasbinModelFilePath() string {
	return filepath.Join(TestDataPath(), "adapter", "rbac_model.conf")
}
