package orderItemDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.OrderItem, error)
	FindById(id string) (models.OrderItem, error)
	FindByOrderId(orderId string) ([]models.OrderItem, error)
	FindOrderItemReview(id string) (models.OrderItem, error)
	Insert(item models.OrderItem) (models.OrderItem, error)
	Update(item models.OrderItem) error
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

func (d dataAccess) FindAll() ([]models.OrderItem, error) {
	var items []models.OrderItem
	result := d.db.Table(models.OrderItem{}.TableName()).
		Where("del_flg = ?", false).
		Preload("Product").
		Find(&items)
	if result.Error != nil {
		return []models.OrderItem{}, result.Error
	}
	return items, nil
}

func (d dataAccess) FindById(id string) (models.OrderItem, error) {
	var item models.OrderItem
	result := d.db.Table(models.OrderItem{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Product").
		First(&item)
	if result.Error != nil {
		return models.OrderItem{}, result.Error
	}
	return item, nil
}

func (d dataAccess) FindByOrderId(orderId string) ([]models.OrderItem, error) {
	var items []models.OrderItem
	result := d.db.Table(models.OrderItem{}.TableName()).
		Where("order_id = ? AND del_flg = ?", orderId, false).
		Preload("Product").
		Find(&items)
	if result.Error != nil {
		return []models.OrderItem{}, result.Error
	}
	return items, nil
}

func (d dataAccess) FindOrderItemReview(id string) (models.OrderItem, error) {
	var item models.OrderItem
	result := d.db.Table(models.OrderItem{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Product").
		Preload("Review").
		First(&item)
	if result.Error != nil {
		return models.OrderItem{}, result.Error
	}
	return item, nil
}

func (d dataAccess) Insert(item models.OrderItem) (models.OrderItem, error) {
	result := d.db.Table(item.TableName()).
	Create(&item)
	if result.Error != nil {
		return models.OrderItem{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.OrderItem) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.OrderItem
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.OrderItem
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
