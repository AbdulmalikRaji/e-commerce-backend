package migration

import (
	"log"
	"sync"

	"gorm.io/gorm"
)

var onlyOnce sync.Once

func Migrate(connection *gorm.DB) {

	onlyOnce.Do(func() {

		log.Println("Migrating the database...")

		if err := connection.AutoMigrate(); err != nil {
			log.Fatalf("Could not migrate: %v", err)
		}

		log.Println("Database migration completed successfully.")
	})
}
