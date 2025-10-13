package productvariantdao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.ProductVariant, error)
	FindById(id string) (models.ProductVariant, error)
	FindByProductId(productId string) ([]models.ProductVariant, error)
	Insert(item models.ProductVariant) (models.ProductVariant, error)
	Update(item models.ProductVariant) error
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

func (d dataAccess) FindAll() ([]models.ProductVariant, error) {
	var variants []models.ProductVariant
	result := d.db.Table(models.ProductVariant{}.TableName()).
		Where("del_flg = ?", false).
		Preload("Product").
		Find(&variants)
	if result.Error != nil {
		return []models.ProductVariant{}, result.Error
	}
	return variants, nil
}

func (d dataAccess) FindById(id string) (models.ProductVariant, error) {
	var variant models.ProductVariant
	result := d.db.Table(models.ProductVariant{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Product").
		First(&variant)
	if result.Error != nil {
		return models.ProductVariant{}, result.Error
	}
	return variant, nil
}

func (d dataAccess) FindByProductId(productId string) ([]models.ProductVariant, error) {
	var variants []models.ProductVariant
	result := d.db.Table(models.ProductVariant{}.TableName()).
		Where("product_id = ? AND del_flg = ?", productId, false).
		Preload("Product").
		Find(&variants)
	if result.Error != nil {
		return []models.ProductVariant{}, result.Error
	}
	return variants, nil
}

func (d dataAccess) Insert(item models.ProductVariant) (models.ProductVariant, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.ProductVariant{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.ProductVariant) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.ProductVariant
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.ProductVariant
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
