package productimagedao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.ProductImage, error)
	FindById(id string) (models.ProductImage, error)
	FindByProductId(productId string) ([]models.ProductImage, error)
	Insert(item models.ProductImage) (models.ProductImage, error)
	Update(item models.ProductImage) error
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

func (d dataAccess) FindAll() ([]models.ProductImage, error) {
	var items []models.ProductImage
	result := d.db.Table(models.ProductImage{}.TableName()).
		Where("del_flg = ?", false).
		Preload("Product").
		Find(&items)
	if result.Error != nil {
		return []models.ProductImage{}, result.Error
	}
	return items, nil
}

func (d dataAccess) FindById(id string) (models.ProductImage, error) {
	var item models.ProductImage
	result := d.db.Table(models.ProductImage{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Product").
		First(&item)
	if result.Error != nil {
		return models.ProductImage{}, result.Error
	}
	return item, nil
}

func (d dataAccess) FindByProductId(productId string) ([]models.ProductImage, error) {
	var items []models.ProductImage
	result := d.db.Table(models.ProductImage{}.TableName()).
		Where("product_id = ? AND del_flg = ?", productId, false).
		Preload("Product").
		Find(&items)
	if result.Error != nil {
		return []models.ProductImage{}, result.Error
	}
	return items, nil
}

func (d dataAccess) Insert(item models.ProductImage) (models.ProductImage, error) {
	result := d.db.Table(item.TableName()).
		Create(&item)
	if result.Error != nil {
		return models.ProductImage{}, result.Error
	}
	return item, nil
}

func (d dataAccess) Update(item models.ProductImage) error {
	result := d.db.Table(item.TableName()).
		Where("id = ?", item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.ProductImage
	result := d.db.Table(item.TableName()).
		Where("id = ?", id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.ProductImage
	result := d.db.Table(item.TableName()).
		Where("id = ?", id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
