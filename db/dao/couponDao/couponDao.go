package couponDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.Coupon, error)
	FindById(id string) (models.Coupon, error)
	FindByCode(code string) (models.Coupon, error)
	Insert(item models.Coupon) (models.Coupon, error)
	Update(item models.Coupon) error
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

func (d dataAccess) FindAll() ([]models.Coupon, error) {
	var coupons []models.Coupon
	result := d.db.Table(models.Coupon{}.TableName()).
		Where("del_flg = ?", false).
		Find(&coupons)
	if result.Error != nil {
		return []models.Coupon{}, result.Error
	}
	return coupons, nil
}

func (d dataAccess) FindById(id string) (models.Coupon, error) {
	var coupon models.Coupon
	result := d.db.Table(models.Coupon{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		First(&coupon)
	if result.Error != nil {
		return models.Coupon{}, result.Error
	}
	return coupon, nil
}

func (d dataAccess) FindByCode(code string) (models.Coupon, error) {
	var coupon models.Coupon
	result := d.db.Table(models.Coupon{}.TableName()).
		Where("code = ? AND del_flg = ?", code, false).
		First(&coupon)
	if result.Error != nil {
		return models.Coupon{}, result.Error
	}
	return coupon, nil
}

func (d dataAccess) Insert(item models.Coupon) (models.Coupon, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.Coupon{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.Coupon) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.Coupon
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.Coupon
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
