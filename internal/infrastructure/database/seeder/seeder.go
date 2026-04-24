// database/seeder.go
package seeder

import (
	"log"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	log.Println("Running seeders...")

	// Seed roles
	// roles := []models.Role{
	//     {Name: "admin"},
	//     {Name: "user"},
	//     {Name: "editor"},
	// }
	// for _, role := range roles {
	//     // FirstOrCreate → hanya insert kalau belum ada
	//     db.FirstOrCreate(&role, models.Role{Name: role.Name})
	// }

	// // Seed categories
	// categories := []models.Category{
	//     {Name: "Electronics"},
	//     {Name: "Clothing"},
	//     {Name: "Books"},
	// }
	// for _, cat := range categories {
	//     db.FirstOrCreate(&cat, models.Category{Name: cat.Name})
	// }

	log.Println("Seeding completed")
	return nil
}
