// database/migration.go
package migrations

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Println("Running migrations...")

	err := db.AutoMigrate()

	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Migration completed successfully")
	return nil
}
