package orderDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	// Postgres Data Access Object Methods
	FindAll() ([]models.Order, error)
	FindById(id string) (models.Order, error)
	FindByBuyerId(userId string) ([]models.Order, error)
	FindOrderItems(id string) (models.Order, error)
	Insert(item models.Order) (models.Order, error)
	Update(item models.Order) error
	SoftDelete(id string) error
	Delete(id string) error
}

type dataAccess struct {
	db *gorm.DB
}

func New(client connection.Client) DataAccess {
	return dataAccess{
		db: client.PostgresConnection,
	}
}

func (d dataAccess) FindAll() ([]models.Order, error) {
	var orders []models.Order
	result := d.db.Table(models.Order{}.TableName()).
		Where("del_flg = ?", false).
		Preload("Buyer").
		Find(&orders)
	if result.Error != nil {
		return []models.Order{}, result.Error
	}
	return orders, nil
}

func (d dataAccess) FindById(id string) (models.Order, error) {
	var order models.Order
	result := d.db.Table(models.Order{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Buyer").
		Preload("Items").
		Preload("ShippingAddress").
		First(&order)
	if result.Error != nil {
		return models.Order{}, result.Error
	}
	return order, nil
}

func (d dataAccess) FindByBuyerId(userId string) ([]models.Order, error) {
	var orders []models.Order
	result := d.db.Table(models.Order{}.TableName()).
		Where("buyer_id = ? AND del_flg = ?", userId, false).
		Preload("Buyer").
		Preload("Items").
		Preload("ShippingAddress").
		Find(&orders)
	if result.Error != nil {
		return []models.Order{}, result.Error
	}
	return orders, nil
}

func (d dataAccess) FindOrderItems(id string) (models.Order, error) {
	var order models.Order
	result := d.db.Table(models.Order{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Items").
		First(&order)
	if result.Error != nil {
		return models.Order{}, result.Error
	}
	return order, nil
}

func (d dataAccess) Insert(item models.Order) (models.Order, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.Order{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.Order) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.Order
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.Order
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
