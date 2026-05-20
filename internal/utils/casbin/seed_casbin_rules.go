package casbin_utils

import (
	"fmt"

	constants "github.com/ardnh/be-travel-booking-app/pkg/constants"
	"github.com/casbin/casbin/v3"
)

func SeedCasbinRules(enforcer *casbin.Enforcer) error {

	existingPolicies, _ := enforcer.GetPolicy()
	if len(existingPolicies) > 0 {
		return nil
	}

	// === Policies per role ===
	policies := [][]string{
		// daily_user
		{constants.RoleDailyUser, constants.ResourceProfile, constants.ActionRead},
		{constants.RoleDailyUser, constants.ResourceProfile, constants.ActionUpdate},
		{constants.RoleDailyUser, constants.ResourceServiceTypes, constants.ActionRead},
		{constants.RoleDailyUser, constants.ResourceVendors, constants.ActionRead},
		{constants.RoleDailyUser, constants.ResourceSchedules, constants.ActionRead},
		{constants.RoleDailyUser, constants.ResourcePoolPoints, constants.ActionRead},
		{constants.RoleDailyUser, constants.ResourceLayouts, constants.ActionRead},
		{constants.RoleDailyUser, constants.ResourceLayoutPositions, constants.ActionRead},

		// business_owner
		{constants.RoleBusinessOwner, constants.ResourceVendors, constants.ActionCreate},
		{constants.RoleBusinessOwner, constants.ResourceVendors, constants.ActionUpdate},
		{constants.RoleBusinessOwner, constants.ResourceVendors, constants.ActionDelete},
		{constants.RoleBusinessOwner, constants.ResourceSchedules, constants.ActionCreate},
		{constants.RoleBusinessOwner, constants.ResourceSchedules, constants.ActionUpdate},
		{constants.RoleBusinessOwner, constants.ResourceSchedules, constants.ActionDelete},
		{constants.RoleBusinessOwner, constants.ResourceLayouts, constants.ActionCreate},
		{constants.RoleBusinessOwner, constants.ResourceLayouts, constants.ActionUpdate},
		{constants.RoleBusinessOwner, constants.ResourceLayouts, constants.ActionDelete},
		{constants.RoleBusinessOwner, constants.ResourceLayoutPositions, constants.ActionCreate},
		{constants.RoleBusinessOwner, constants.ResourceLayoutPositions, constants.ActionUpdate},
		{constants.RoleBusinessOwner, constants.ResourceLayoutPositions, constants.ActionDelete},
		{constants.RoleBusinessOwner, constants.ResourcePoolPoints, constants.ActionCreate},
		{constants.RoleBusinessOwner, constants.ResourcePoolPoints, constants.ActionUpdate},
		{constants.RoleBusinessOwner, constants.ResourcePoolPoints, constants.ActionDelete},

		// admin_platform
		{constants.RoleAdminPlatform, constants.ResourceUsers, constants.ActionCreate},
		{constants.RoleAdminPlatform, constants.ResourceUsers, constants.ActionUpdate},
		{constants.RoleAdminPlatform, constants.ResourceUsers, constants.ActionDelete},
		{constants.RoleAdminPlatform, constants.ResourceUserRoles, constants.ActionRead},
		{constants.RoleAdminPlatform, constants.ResourceUserRoles, constants.ActionCreate},
		{constants.RoleAdminPlatform, constants.ResourceUserRoles, constants.ActionUpdate},
		{constants.RoleAdminPlatform, constants.ResourceUserRoles, constants.ActionDelete},
		{constants.RoleAdminPlatform, constants.ResourceServiceTypes, constants.ActionCreate},
		{constants.RoleAdminPlatform, constants.ResourceServiceTypes, constants.ActionUpdate},
		{constants.RoleAdminPlatform, constants.ResourceServiceTypes, constants.ActionDelete},

		// platform_owner
		{constants.RolePlatformOwner, constants.ResourcePlatformDashboard, constants.ActionRead},
		{constants.RolePlatformOwner, constants.ResourcePlatformSettings, constants.ActionRead},
		{constants.RolePlatformOwner, constants.ResourcePlatformSettings, constants.ActionUpdate},
	}

	_, err := enforcer.AddPolicies(policies)
	if err != nil {
		return fmt.Errorf("gagal seed policies: %w", err)
	}

	// === Role Hierarchy ===
	for childRole, parentRole := range constants.RoleHierarchy {
		_, err := enforcer.AddGroupingPolicy(childRole, parentRole)
		if err != nil {
			return fmt.Errorf("gagal seed hierarchy %s -> %s: %w", childRole, parentRole, err)
		}
	}
	return nil
}
