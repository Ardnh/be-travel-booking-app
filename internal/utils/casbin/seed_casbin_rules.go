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

		// ─── daily_user ───────────────────────────────────────────────
		{constants.RoleDailyUser, constants.ResourceProfile, constants.ActionRead},
		{constants.RoleDailyUser, constants.ResourceProfile, constants.ActionUpdate},

		{constants.RoleDailyUser, constants.ResourceServiceTypes, constants.ActionRead},

		{constants.RoleDailyUser, constants.ResourceVendors, constants.ActionRead},
		{constants.RoleDailyUser, constants.ResourceVendors, constants.ActionCreate}, // register vendor

		{constants.RoleDailyUser, constants.ResourceSchedules, constants.ActionRead},

		{constants.RoleDailyUser, constants.ResourcePoolPoints, constants.ActionRead},

		{constants.RoleDailyUser, constants.ResourceLayouts, constants.ActionRead},

		{constants.RoleDailyUser, constants.ResourceLayoutPositions, constants.ActionRead},

		{constants.RoleDailyUser, constants.ResourceBookings, constants.ActionCreate},
		{constants.RoleDailyUser, constants.ResourceBookings, constants.ActionRead},

		// ─── business_owner ──────────────────────────────────────────
		{constants.RoleBusinessOwner, constants.ResourceProfile, constants.ActionRead},

		{constants.RoleBusinessOwner, constants.ResourceVendors, constants.ActionRead},
		{constants.RoleBusinessOwner, constants.ResourceVendors, constants.ActionUpdate},
		{constants.RoleBusinessOwner, constants.ResourceVendors, constants.ActionDelete},

		{constants.RoleBusinessOwner, constants.ResourcePoolPoints, constants.ActionRead},
		{constants.RoleBusinessOwner, constants.ResourcePoolPoints, constants.ActionCreate},
		{constants.RoleBusinessOwner, constants.ResourcePoolPoints, constants.ActionUpdate},
		{constants.RoleBusinessOwner, constants.ResourcePoolPoints, constants.ActionDelete},

		{constants.RoleBusinessOwner, constants.ResourceSchedules, constants.ActionRead},
		{constants.RoleBusinessOwner, constants.ResourceSchedules, constants.ActionCreate},
		{constants.RoleBusinessOwner, constants.ResourceSchedules, constants.ActionUpdate},
		{constants.RoleBusinessOwner, constants.ResourceSchedules, constants.ActionDelete},

		{constants.RoleBusinessOwner, constants.ResourceBookings, constants.ActionRead},

		{constants.RoleBusinessOwner, constants.ResourceLayouts, constants.ActionRead},
		{constants.RoleBusinessOwner, constants.ResourceLayouts, constants.ActionCreate},
		{constants.RoleBusinessOwner, constants.ResourceLayouts, constants.ActionUpdate},
		{constants.RoleBusinessOwner, constants.ResourceLayouts, constants.ActionDelete},

		{constants.RoleBusinessOwner, constants.ResourceLayoutPositions, constants.ActionRead},
		{constants.RoleBusinessOwner, constants.ResourceLayoutPositions, constants.ActionCreate},
		{constants.RoleBusinessOwner, constants.ResourceLayoutPositions, constants.ActionUpdate},
		{constants.RoleBusinessOwner, constants.ResourceLayoutPositions, constants.ActionDelete},

		{constants.RoleBusinessOwner, constants.ResourceUserRoles, constants.ActionRead},
		{constants.RoleBusinessOwner, constants.ResourceUserRoles, constants.ActionCreate},
		{constants.RoleBusinessOwner, constants.ResourceUserRoles, constants.ActionUpdate},
		{constants.RoleBusinessOwner, constants.ResourceUserRoles, constants.ActionDelete},

		// ─── admin_owner ─────────────────────────────────────────────
		{constants.RoleAdminBusiness, constants.ResourceProfile, constants.ActionRead},

		{constants.RoleAdminBusiness, constants.ResourcePoolPoints, constants.ActionRead},
		{constants.RoleAdminBusiness, constants.ResourcePoolPoints, constants.ActionCreate},
		{constants.RoleAdminBusiness, constants.ResourcePoolPoints, constants.ActionUpdate},

		{constants.RoleAdminBusiness, constants.ResourceSchedules, constants.ActionRead},
		{constants.RoleAdminBusiness, constants.ResourceSchedules, constants.ActionCreate},
		{constants.RoleAdminBusiness, constants.ResourceSchedules, constants.ActionUpdate},

		{constants.RoleAdminBusiness, constants.ResourceBookings, constants.ActionRead},

		{constants.RoleAdminBusiness, constants.ResourceUserRoles, constants.ActionRead},
		{constants.RoleAdminBusiness, constants.ResourceUserRoles, constants.ActionCreate},
		{constants.RoleAdminBusiness, constants.ResourceUserRoles, constants.ActionUpdate},
		{constants.RoleAdminBusiness, constants.ResourceUserRoles, constants.ActionDelete},

		// ─── admin_pool ──────────────────────────────────────────────
		{constants.RoleAdminPool, constants.ResourceProfile, constants.ActionRead},

		{constants.RoleAdminPool, constants.ResourcePoolPoints, constants.ActionRead},

		{constants.RoleAdminPool, constants.ResourceSchedules, constants.ActionRead},
		{constants.RoleAdminPool, constants.ResourceSchedules, constants.ActionCreate},
		{constants.RoleAdminPool, constants.ResourceSchedules, constants.ActionUpdate},

		{constants.RoleAdminPool, constants.ResourceBookings, constants.ActionRead},
		{constants.RoleAdminPool, constants.ResourceBookings, constants.ActionCreate},
		{constants.RoleAdminPool, constants.ResourceBookings, constants.ActionUpdate},

		// ─── admin_platform ──────────────────────────────────────────
		{constants.RoleAdminPlatform, constants.ResourceProfile, constants.ActionRead},

		{constants.RoleAdminPlatform, constants.ResourceUsers, constants.ActionCreate},
		{constants.RoleAdminPlatform, constants.ResourceUsers, constants.ActionUpdate},
		{constants.RoleAdminPlatform, constants.ResourceUsers, constants.ActionDelete},

		{constants.RoleAdminPlatform, constants.ResourceUserRoles, constants.ActionRead},
		{constants.RoleAdminPlatform, constants.ResourceUserRoles, constants.ActionCreate},
		{constants.RoleAdminPlatform, constants.ResourceUserRoles, constants.ActionUpdate},
		{constants.RoleAdminPlatform, constants.ResourceUserRoles, constants.ActionDelete},

		{constants.RoleAdminPlatform, constants.ResourceServiceTypes, constants.ActionRead},
		{constants.RoleAdminPlatform, constants.ResourceServiceTypes, constants.ActionCreate},
		{constants.RoleAdminPlatform, constants.ResourceServiceTypes, constants.ActionUpdate},
		{constants.RoleAdminPlatform, constants.ResourceServiceTypes, constants.ActionDelete},

		{constants.RoleAdminPlatform, constants.ResourceVendors, constants.ActionRead},
		{constants.RoleAdminPlatform, constants.ResourceVendors, constants.ActionUpdate},

		{constants.RoleAdminPlatform, constants.ResourceSchedules, constants.ActionRead},

		{constants.RoleAdminPlatform, constants.ResourceBookings, constants.ActionRead},

		{constants.RoleAdminPlatform, constants.ResourceLayouts, constants.ActionRead},
		{constants.RoleAdminPlatform, constants.ResourceLayouts, constants.ActionCreate},
		{constants.RoleAdminPlatform, constants.ResourceLayouts, constants.ActionUpdate},
		{constants.RoleAdminPlatform, constants.ResourceLayouts, constants.ActionDelete},

		{constants.RoleAdminPlatform, constants.ResourceLayoutPositions, constants.ActionRead},
		{constants.RoleAdminPlatform, constants.ResourceLayoutPositions, constants.ActionCreate},
		{constants.RoleAdminPlatform, constants.ResourceLayoutPositions, constants.ActionUpdate},
		{constants.RoleAdminPlatform, constants.ResourceLayoutPositions, constants.ActionDelete},

		{constants.RoleAdminPlatform, constants.ResourcePoolPoints, constants.ActionRead},

		// ─── platform_owner ──────────────────────────────────────────
		{constants.RolePlatformOwner, constants.ResourcePlatformDashboard, constants.ActionRead},
		{constants.RolePlatformOwner, constants.ResourcePlatformSettings, constants.ActionRead},
		{constants.RolePlatformOwner, constants.ResourcePlatformSettings, constants.ActionUpdate},

		{constants.RolePlatformOwner, constants.ResourceProfile, constants.ActionRead},

		{constants.RolePlatformOwner, constants.ResourceServiceTypes, constants.ActionRead},
		{constants.RolePlatformOwner, constants.ResourceServiceTypes, constants.ActionCreate},
		{constants.RolePlatformOwner, constants.ResourceServiceTypes, constants.ActionUpdate},
		{constants.RolePlatformOwner, constants.ResourceServiceTypes, constants.ActionDelete},

		{constants.RolePlatformOwner, constants.ResourceVendors, constants.ActionRead},
		{constants.RolePlatformOwner, constants.ResourceVendors, constants.ActionUpdate},
		{constants.RolePlatformOwner, constants.ResourceVendors, constants.ActionDelete},

		{constants.RolePlatformOwner, constants.ResourceSchedules, constants.ActionRead},

		{constants.RolePlatformOwner, constants.ResourceBookings, constants.ActionRead},

		{constants.RolePlatformOwner, constants.ResourceUserRoles, constants.ActionRead},
		{constants.RolePlatformOwner, constants.ResourceUserRoles, constants.ActionCreate},
		{constants.RolePlatformOwner, constants.ResourceUserRoles, constants.ActionUpdate},
		{constants.RolePlatformOwner, constants.ResourceUserRoles, constants.ActionDelete},

		{constants.RolePlatformOwner, constants.ResourceLayouts, constants.ActionRead},
		{constants.RolePlatformOwner, constants.ResourceLayouts, constants.ActionCreate},
		{constants.RolePlatformOwner, constants.ResourceLayouts, constants.ActionUpdate},
		{constants.RolePlatformOwner, constants.ResourceLayouts, constants.ActionDelete},

		{constants.RolePlatformOwner, constants.ResourceLayoutPositions, constants.ActionRead},
		{constants.RolePlatformOwner, constants.ResourceLayoutPositions, constants.ActionCreate},
		{constants.RolePlatformOwner, constants.ResourceLayoutPositions, constants.ActionUpdate},
		{constants.RolePlatformOwner, constants.ResourceLayoutPositions, constants.ActionDelete},

		{constants.RolePlatformOwner, constants.ResourcePoolPoints, constants.ActionRead},

		{constants.RolePlatformOwner, constants.ResourceUsers, constants.ActionCreate},
		{constants.RolePlatformOwner, constants.ResourceUsers, constants.ActionUpdate},
		{constants.RolePlatformOwner, constants.ResourceUsers, constants.ActionDelete},
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
