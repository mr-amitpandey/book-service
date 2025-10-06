package constants

type Role string

const (
	RoleSuperAdmin Role = "super_admin" // Super Admin - Full system access
	RoleAdmin      Role = "admin"       // Admin - Full system access
	RoleManager    Role = "manager"     // Manager - Manages teams and projects
	RoleHRManager  Role = "hr_manager"  // HR Manager - HR related tasks
	RoleEmployee   Role = "employee"    // Employee
)

// Role groups for easier middleware usage
var (
	// AdminRoles includes administrative roles
	AdminRoles = []Role{RoleSuperAdmin, RoleAdmin}

	// ManagementRoles includes admin and management roles
	ManagementRoles = []Role{RoleAdmin, RoleManager, RoleSuperAdmin}

	// ContentManagementRoles includes roles that can manage content
	ContentManagementRoles = []Role{RoleAdmin, RoleManager, RoleSuperAdmin}

	// AllRoles includes all defined roles
	AllRoles = []Role{RoleAdmin, RoleManager, RoleSuperAdmin, RoleHRManager, RoleEmployee}
)

// Helper functions for role checking
func (r Role) String() string {
	return string(r)
}

func (r Role) IsAdmin() bool {
	return r == RoleAdmin
}

func (r Role) IsManagement() bool {
	return r == RoleAdmin || r == RoleManager || r == RoleHRManager || r == RoleSuperAdmin
}

func (r Role) CanManageContent() bool {
	return r == RoleAdmin || r == RoleManager || r == RoleSuperAdmin
}

func (r Role) IsReadOnly() bool {
	return r == RoleEmployee
}

// Convert string slice to Role slice
func ToRoles(roleStrings []string) []Role {
	roles := make([]Role, len(roleStrings))
	for i, roleStr := range roleStrings {
		roles[i] = Role(roleStr)
	}
	return roles
}

// Convert Role slice to string slice
func ToStrings(roles []Role) []string {
	strings := make([]string, len(roles))
	for i, role := range roles {
		strings[i] = role.String()
	}
	return strings
}
