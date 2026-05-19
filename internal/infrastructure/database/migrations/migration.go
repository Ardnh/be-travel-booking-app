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
		// 1. Tabel tanpa FK (independen)
		&entities.Users{},
		&entities.Vendors{},
		&entities.ServiceTypes{},
		&entities.Layouts{},

		// 2. Tabel yang bergantung pada level 1
		&entities.Pools{},
		&entities.LayoutPositions{},
		&entities.UserRoles{},

		// 3. Tabel yang bergantung pada level 2
		&entities.Schedules{},

		// 4. Tabel yang bergantung pada level 3
		&entities.Bookings{},
		&entities.SeatHolds{},

		// 5. Tabel yang bergantung pada level 4
		&entities.BookingSeats{},
		&entities.Payments{},
	)

	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Migration completed successfully")
	return nil
}
