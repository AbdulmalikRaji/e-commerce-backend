package migration

import (
	"log"
	"sync"

	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

var onlyOnce sync.Once

func Migrate(connection *gorm.DB) {

	onlyOnce.Do(func() {

		log.Println("Migrating the database...")

		if err := connection.AutoMigrate(
			&models.User{},
			&models.Category{},
			&models.Product{},
			&models.Order{},
			&models.OrderItem{},
			&models.Payment{},
			&models.Address{},
			&models.Review{},
		); err != nil {
			log.Fatalf("Could not migrate: %v", err)
		}

		log.Println("Database migration completed successfully.")
	})
}
