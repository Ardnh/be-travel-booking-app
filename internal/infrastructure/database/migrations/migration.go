// database/migration.go
package migrations

import (
	"fmt"
	"log"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Println("Running migrations...")

	err := db.AutoMigrate(
		entities.Users{},
		entities.UsersRole{},
		entities.ServiceType{},
		entities.Layout{},
		entities.LayoutPosition{},
	)

	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Migration completed successfully")
	return nil
}
