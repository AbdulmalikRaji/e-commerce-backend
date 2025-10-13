package analyticsDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	LogProductView(view models.ProductView) error
	LogAddToCart(event models.AddToCartEvent) error
	LogAbandonedCart(event models.AbandonedCart) error
	LogSalesStat(stat models.SalesStat) error

	FindProductViews(productId string, sessionId *string, userId *string) ([]models.ProductView, error)
	FindAddToCartEvents(productId string, sessionId *string, userId *string) ([]models.AddToCartEvent, error)
	FindAbandonedCarts(userId *string, sessionId *string) ([]models.AbandonedCart, error)
	FindSalesStats(productId string, sessionId *string, userId *string) ([]models.SalesStat, error)
}

type dataAccess struct {
	db *gorm.DB
}

func New(client connection.Client) DataAccess {
	return dataAccess{
		db: client.PostgresConnection,
	}
}

func (d dataAccess) LogProductView(view models.ProductView) error {
	return d.db.Table(view.TableName()).Create(&view).Error
}
func (d dataAccess) LogAddToCart(event models.AddToCartEvent) error {
	return d.db.Table(event.TableName()).Create(&event).Error
}
func (d dataAccess) LogAbandonedCart(event models.AbandonedCart) error {
	return d.db.Table(event.TableName()).Create(&event).Error
}
func (d dataAccess) LogSalesStat(stat models.SalesStat) error {
	return d.db.Table(stat.TableName()).Create(&stat).Error
}

func (d dataAccess) FindProductViews(productId string, sessionId *string, userId *string) ([]models.ProductView, error) {
	var views []models.ProductView
	query := d.db.Table(models.ProductView{}.TableName()).Where("product_id = ?", productId)
	if sessionId != nil {
		query = query.Where("session_id = ?", *sessionId)
	}
	if userId != nil {
		query = query.Where("user_id = ?", *userId)
	}
	err := query.Find(&views).Error
	return views, err
}

func (d dataAccess) FindAddToCartEvents(productId string, sessionId *string, userId *string) ([]models.AddToCartEvent, error) {
	var events []models.AddToCartEvent
	query := d.db.Table(models.AddToCartEvent{}.TableName()).Where("product_id = ?", productId)
	if sessionId != nil {
		query = query.Where("session_id = ?", *sessionId)
	}
	if userId != nil {
		query = query.Where("user_id = ?", *userId)
	}
	err := query.Find(&events).Error
	return events, err
}

func (d dataAccess) FindAbandonedCarts(userId *string, sessionId *string) ([]models.AbandonedCart, error) {
	var carts []models.AbandonedCart
	query := d.db.Table(models.AbandonedCart{}.TableName())
	if userId != nil {
		query = query.Where("user_id = ?", *userId)
	}
	if sessionId != nil {
		query = query.Where("session_id = ?", *sessionId)
	}
	err := query.Find(&carts).Error
	return carts, err
}

func (d dataAccess) FindSalesStats(productId string, sessionId *string, userId *string) ([]models.SalesStat, error) {
	var stats []models.SalesStat
	query := d.db.Table(models.SalesStat{}.TableName()).Where("product_id = ?", productId)
	if sessionId != nil {
		query = query.Where("session_id = ?", *sessionId)
	}
	if userId != nil {
		query = query.Where("user_id = ?", *userId)
	}
	err := query.Find(&stats).Error
	return stats, err
}
