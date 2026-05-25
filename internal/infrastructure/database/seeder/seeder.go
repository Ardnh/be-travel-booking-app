// database/seeder.go
package seeder

import (
	"fmt"
	"log"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	log.Println("Running seeders...")

	// Check if admin user already exists
	var count int64
	db.Model(&entities.Users{}).Where("email = ?", "platform-owner@example.com").Count(&count)
	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}

		admin := entities.Users{
			Name:         "Ardan",
			Email:        "platform-owner@example.com",
			Phone:        "0893748374",
			PasswordHash: string(hashedPassword),
		}

		if err := db.Create(&admin).Error; err != nil {
			return fmt.Errorf("failed to seed admin user: %w", err)
		}

		// userRole := entities.UserRoles{
		// 	UserID: admin.UserID,
		// 	Role:   constants.RolePlatformOwner,
		// }

		// if err := db.Create(&userRole).Error; err != nil {
		// 	return fmt.Errorf("failed to seed user role: %w", err)
		// }

		log.Println("Admin user seeded successfully")
		return nil
	} else {
		log.Println("Admin user already exists, skipping...")
	}

	log.Println("Seeding completed")
	return nil
}
