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
		log.Println("Starting database migration...")

		// Begin transaction
		tx := connection.Begin()
		if tx.Error != nil {
			log.Fatalf("Could not begin transaction: %v", tx.Error)
		}

		// Ensure rollback in case of error
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				log.Fatalf("Panic during migration, rolled back: %v", r)
			}
		}()

		log.Println("Creating tables...")

		if err := tx.AutoMigrate(
			&models.User{},
			&models.Category{},
			&models.Product{},
			&models.ProductImage{},
			&models.ProductVariant{},
			&models.Order{},
			&models.OrderItem{},
			&models.Payment{},
			&models.Address{},
			&models.Review{},
			&models.Store{},
			&models.Cart{},
			&models.CartItem{},
			&models.Coupon{},
			&models.Notification{},
			&models.Refund{},
			&models.Subcategory{},
			&models.Tag{},
			&models.Warehouse{},
			&models.WarehouseStock{},
			&models.UserToken{},
			&models.ProductView{},
			&models.AddToCartEvent{},
			&models.SearchAnalytics{},
			&models.StoreVisit{},
			&models.MarketingAnalytics{},
			&models.SalesStat{},
		); err != nil {
			tx.Rollback()
			log.Fatalf("Could not migrate, rolling back: %v", err)
		}

		// Commit transaction
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			log.Fatalf("Could not commit migration transaction: %v", err)
		}

		log.Println("Database migration completed successfully and committed.")
	})
}
