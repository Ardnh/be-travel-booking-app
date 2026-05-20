package casbin_utils

import (
	"github.com/casbin/casbin/v3"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitCasbin(modelPath string, db *gorm.DB) (*casbin.Enforcer, error) {
	// Buat adapter dari koneksi GORM yang sudah ada
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	// Buat enforcer dengan model dan adapter
	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		return nil, err
	}

	// Load policy dari database
	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}

	return enforcer, nil
}

func GetUserPermissions(enforcer *casbin.Enforcer, userID string) []string {
	permissions, _ := enforcer.GetImplicitPermissionsForUser(userID)
	result := []string{}
	for _, p := range permissions {
		result = append(result, p[1]+":"+p[2])
	}
	return result
}
