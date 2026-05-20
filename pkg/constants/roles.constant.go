package constants

const (
	RoleDailyUser     = "daily_user"
	RolePlatformOwner = "platform_owner"
	RoleAdminPlatform = "admin_platform"
	RoleBusinessOwner = "business_owner"
	RoleAdminBusiness = "admin_business"
	RoleAdminPool     = "admin_pool"
)

// Role anak → role induk
// Artinya role anak mewarisi semua permission role induk
var RoleHierarchy = map[string]string{
	RoleAdminPlatform: RolePlatformOwner, // admin_platform mewarisi platform_owner
	RoleAdminBusiness: RoleBusinessOwner, // admin_business mewarisi business_owner
	RolePlatformOwner: RoleDailyUser,     // platform_owner mewarisi daily_user
	RoleBusinessOwner: RoleDailyUser,     // business_owner mewarisi daily_user
	RoleAdminPool:     RoleDailyUser,     // admin_pool mewarisi daily_user
}

var ValidRoles = map[string]bool{
	RoleDailyUser:     true,
	RolePlatformOwner: true,
	RoleAdminPlatform: true,
	RoleBusinessOwner: true,
	RoleAdminBusiness: true,
	RoleAdminPool:     true,
}

func IsValidRole(role string) bool {
	return ValidRoles[role]
}
