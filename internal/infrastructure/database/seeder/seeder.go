// database/seeder.go
package seeder

import (
	"fmt"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/pkg/constants"
	"github.com/casbin/casbin/v3"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB, casbinEnforcer *casbin.Enforcer, log *logrus.Logger) error {
	log.Println("Running seeders...")

	// Check if admin user already exists
	var count int64
	db.Model(&entities.Users{}).Where("email = ?", "platform-owner@example.com").Count(&count)
	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}

		// platform
		admin := entities.Users{
			Name:         "Ardan",
			Email:        "platform-owner@example.com",
			Phone:        "0893748374",
			PasswordHash: string(hashedPassword),
		}

		if err := db.Create(&admin).Error; err != nil {
			return fmt.Errorf("failed to seed admin user: %w", err)
		}

		_, err = casbinEnforcer.AddGroupingPolicy(admin.UserID.String(), constants.RolePlatformOwner)
		if err != nil {
			log.WithFields(logrus.Fields{
				"email":  admin.Email,
				"userID": admin.UserID.String(),
				"error":  err,
			}).Error("failed to add casbin grouping policy")
			return fmt.Errorf("failed to seed admin user: %w", err)
		}

		casbinEnforcer.LoadPolicy()

		log.Println("Admin user seeded successfully")
		return nil
	} else {
		log.Println("Admin user already exists, skipping...")
	}

	log.Println("Seeding completed")
	return nil
}
