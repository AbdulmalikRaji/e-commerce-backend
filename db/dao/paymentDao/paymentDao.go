package paymentDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.Payment, error)
	FindById(id string) (models.Payment, error)
	FindByOrderId(orderId string) ([]models.Payment, error)
	FindLatestByOrderId(orderId string) (models.Payment, error)
	Insert(item models.Payment) (models.Payment, error)
	Update(item models.Payment) error
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

func (d dataAccess) FindAll() ([]models.Payment, error) {
	var payments []models.Payment
	result := d.db.Table(models.Payment{}.TableName()).
		Where("del_flg = ?", false).
		Preload("Order").
		Find(&payments)
	if result.Error != nil {
		return []models.Payment{}, result.Error
	}
	return payments, nil
}

func (d dataAccess) FindById(id string) (models.Payment, error) {
	var payment models.Payment
	result := d.db.Table(models.Payment{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Order").
		First(&payment)
	if result.Error != nil {
		return models.Payment{}, result.Error
	}
	return payment, nil
}

func (d dataAccess) FindByOrderId(orderId string) ([]models.Payment, error) {
	var payments []models.Payment
	result := d.db.Table(models.Payment{}.TableName()).
		Where("order_id = ? AND del_flg = ?", orderId, false).
		Preload("Order").
		Find(&payments)
	if result.Error != nil {
		return []models.Payment{}, result.Error
	}
	return payments, nil
}

func (d dataAccess) FindLatestByOrderId(orderId string) (models.Payment, error) {
	var payment models.Payment
	result := d.db.Table(models.Payment{}.TableName()).
		Where("order_id = ? AND del_flg = ?", orderId, false).
		Order("created_at DESC").
		Preload("Order").
		First(&payment)
	if result.Error != nil {
		return models.Payment{}, result.Error
	}
	return payment, nil
}

func (d dataAccess) Insert(item models.Payment) (models.Payment, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.Payment{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.Payment) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.Payment
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.Payment
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
