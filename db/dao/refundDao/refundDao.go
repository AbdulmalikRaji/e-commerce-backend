package refundDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.Refund, error)
	FindById(id string) (models.Refund, error)
	FindByPaymentId(paymentId string) ([]models.Refund, error)
	FindByOrderId(orderId string) ([]models.Refund, error)
	Insert(item models.Refund) (models.Refund, error)
	Update(item models.Refund) error
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

func (d dataAccess) FindAll() ([]models.Refund, error) {
	var refunds []models.Refund
	result := d.db.Table(models.Refund{}.TableName()).
		Where("del_flg = ?", false).
		Preload("Payment").
		Find(&refunds)
	if result.Error != nil {
		return []models.Refund{}, result.Error
	}
	return refunds, nil
}

func (d dataAccess) FindById(id string) (models.Refund, error) {
	var refund models.Refund
	result := d.db.Table(models.Refund{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Payment").
		First(&refund)
	if result.Error != nil {
		return models.Refund{}, result.Error
	}
	return refund, nil
}

func (d dataAccess) FindByPaymentId(paymentId string) ([]models.Refund, error) {
	var refunds []models.Refund
	result := d.db.Table(models.Refund{}.TableName()).
		Where("payment_id = ? AND del_flg = ?", paymentId, false).
		Preload("Payment").
		Find(&refunds)
	if result.Error != nil {
		return []models.Refund{}, result.Error
	}
	return refunds, nil
}

func (d dataAccess) FindByOrderId(orderId string) ([]models.Refund, error) {
	var refunds []models.Refund
	result := d.db.Table(models.Refund{}.TableName()).
		Where("order_id = ? AND del_flg = ?", orderId, false).
		Preload("Payment").
		Find(&refunds)
	if result.Error != nil {
		return []models.Refund{}, result.Error
	}
	return refunds, nil
}

func (d dataAccess) Insert(item models.Refund) (models.Refund, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.Refund{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.Refund) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.Refund
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.Refund
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
