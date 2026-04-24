// cmd/migrate/main.go
package main

import (
	"flag"
	"log"

	"github.com/ardnh/be-travel-booking-app/internal/config"
	"github.com/ardnh/be-travel-booking-app/internal/infrastructure/database/migrations"
	"github.com/ardnh/be-travel-booking-app/internal/infrastructure/database/postgresql"
	"github.com/ardnh/be-travel-booking-app/internal/infrastructure/database/seeder"
)

func main() {
	seed := flag.Bool("seed", false, "run seeder after migration")
	flag.Parse()

	cfg := config.LoadConfig()

	db, err := postgresql.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer postgresql.CloseDB(db)

	log.Println("Running migration...")
	if err := migrations.Migrate(db); err != nil {
		log.Fatal(err)
	}

	if *seed {
		log.Println("Running seeder...")
		seeder.Seed(db)
	}

	log.Println("Done!")
}
