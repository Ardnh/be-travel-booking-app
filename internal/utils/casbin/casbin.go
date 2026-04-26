package casbin_utils

import "github.com/casbin/casbin/v3"

func GetUserPermissions(enforcer *casbin.Enforcer, userID string) []string {
	permissions, _ := enforcer.GetImplicitPermissionsForUser(userID)
	// flatten jadi ["articles:read", "articles:create", ...]
	result := []string{}
	for _, p := range permissions {
		result = append(result, p[1]+":"+p[2]) // resource:action
	}
	return result
}
