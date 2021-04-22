package auth

import (
	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
)

const (
	ResourceUserSettings = "user_settings"
	Resources            = "resources"
)

const (
	ActionCreate   = "create"
	ActionUpdate   = "update"
	ActionRead     = "read"
	ActionReadList = "read_list"
	ActionDelete   = "delete"
	ActionHas      = "has"
)

const (
	RoleRoot   = "root"
	RoleAdmin  = "admin"
	RoleClient = "client"
)

var isClientBasedRole = map[string]bool{
	RoleClient: true,
}

type PermissionMap map[string]map[string]map[string]Policy

type Policy func(interface{}, *userpb.User) bool

type Service struct {
	dynamicPermissions PermissionMap
}

func NewService(permissionsChecker *PermissionsChecker) *Service {
	auth := Service{}
	auth.dynamicPermissions = PermissionMap{
		RoleAdmin: {
			ResourceUserSettings: {
				ActionUpdate: permissionsChecker.CanUpdateUserSettings,
				ActionRead:   permissionsChecker.CanReadUserSettings,
			},
			Resources: {
				ActionUpdate: permissionsChecker.CheckPermission,
			},
		},
	}

	return &auth
}

// CanDynamic checks action is allowed by calling associated function
func (auth *Service) CanDynamic(user *userpb.User, action string, resourceName string, resource interface{}) bool {
	if user.RoleName == RoleRoot {
		return true
	}

	function := auth.getPermissionFunc(user.RoleName, action, resourceName)
	return function(resource, user)
}

// allowFunc always allows access
func allowFunc(_ interface{}, _ *userpb.User) bool {
	return true
}

// blockFunc always block access
func blockFunc(_ interface{}, _ *userpb.User) bool {
	return false
}

// getPermissionFunc returns function by role, action and resourceName.
// Returns blockFunc if proposed func not found
func (auth *Service) getPermissionFunc(role string, action string, resourceName string) Policy {
	if rolePermission, ok := auth.dynamicPermissions[role]; ok {
		if resourcePermission, ok := rolePermission[resourceName]; ok {
			if actionPermission, ok := resourcePermission[action]; ok {
				return actionPermission
			}
		}
	}
	return blockFunc
}
