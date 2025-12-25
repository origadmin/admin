package enums

// RoleType defines the type of a role.
type RoleType int8

const (
	// RoleTypeSystem indicates a system-level role (e.g., Super Admin) that cannot be deleted.
	RoleTypeSystem RoleType = 1
	// RoleTypeUser indicates a user-defined role (e.g., general user, operator).
	RoleTypeUser RoleType = 2
)

// String returns the string representation of the role type.
func (rt RoleType) String() string {
	switch rt {
	case RoleTypeSystem:
		return "system"
	case RoleTypeUser:
		return "user"
	default:
		return "unknown"
	}
}
