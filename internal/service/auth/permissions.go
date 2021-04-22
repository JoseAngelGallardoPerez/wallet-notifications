package auth

import (
	"github.com/Confialink/wallet-notifications/internal/srvdiscovery"
	"context"
	"log"
	"net/http"

	permissionspb "github.com/Confialink/wallet-permissions/rpc/permissions"
	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
)

type Permission string

const (
	ViewSettings   = Permission("view_settings")
	ModifySettings = Permission("modify_settings")

	ViewUserProfiles    = Permission("view_user_profiles")
	ViewAdminProfiles   = Permission("view_admin_profiles")
	ModifyUserProfiles  = Permission("modify_user_profiles")
	ModifyAdminProfiles = Permission("modify_admin_profiles")
)

type PermissionsChecker struct {
}

func NewPermissionsChecker() *PermissionsChecker {
	return &PermissionsChecker{}
}

// CheckPermission calls permission service in order to check if user granted permission
func (p *PermissionsChecker) CheckPermission(permissionValue interface{}, user *userpb.User) bool {
	perm := permissionValue.(Permission)
	result, err := p.Check(user.UID, string(perm))
	if err != nil {
		log.Printf("permission policy failed to check permission: %s", err.Error())
		return false
	}
	return result
}

//Check checks if specified user is granted permission to perform some action
func (p *PermissionsChecker) Check(userId, actionKey string) (bool, error) {
	request := &permissionspb.PermissionReq{UserId: userId, ActionKey: actionKey}

	checker, err := p.checker()
	if nil != err {
		return false, err
	}

	response, err := checker.Check(context.Background(), request)
	if nil != err {
		return false, err
	}
	return response.IsAllowed, nil
}

func (p *PermissionsChecker) CanReadUserSettings(requestedUser interface{}, user *userpb.User) bool {
	reqUser := requestedUser.(*userpb.User)
	var actionKey Permission
	if isClientBasedRole[reqUser.RoleName] {
		actionKey = ViewUserProfiles
	} else {
		actionKey = ViewAdminProfiles
	}

	return p.CheckPermission(actionKey, user)
}

func (p *PermissionsChecker) CanUpdateUserSettings(requestedUser interface{}, user *userpb.User) bool {
	reqUser := requestedUser.(*userpb.User)
	var actionKey Permission
	if isClientBasedRole[reqUser.RoleName] {
		actionKey = ModifyUserProfiles
	} else {
		actionKey = ModifyAdminProfiles
	}

	return p.CheckPermission(actionKey, user)
}

func (p *PermissionsChecker) checker() (permissionspb.PermissionChecker, error) {
	permissionsUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNamePermissions)
	if nil != err {
		return nil, err
	}
	checker := permissionspb.NewPermissionCheckerProtobufClient(permissionsUrl.String(), http.DefaultClient)
	return checker, nil
}
