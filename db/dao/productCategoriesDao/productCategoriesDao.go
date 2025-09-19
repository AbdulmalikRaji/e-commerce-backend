package productCategoriesDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.ProductCategory, error)
	FindById(id string) (models.ProductCategory, error)
	FindByName(name string) (models.ProductCategory, error)
	FindCategoryProducts(id string) (models.ProductCategory, error)
	Insert(item models.ProductCategory) (models.ProductCategory, error)
	Update(item models.ProductCategory) error
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

func (d dataAccess) FindAll() ([]models.ProductCategory, error) {
	var categories []models.ProductCategory
	result := d.db.Table(models.ProductCategory{}.TableName()).
		Where("del_flg = ?", false).
		Find(&categories)
	if result.Error != nil {
		return []models.ProductCategory{}, result.Error
	}
	return categories, nil
}

func (d dataAccess) FindById(id string) (models.ProductCategory, error) {
	var category models.ProductCategory
	result := d.db.Table(models.ProductCategory{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		First(&category)
	if result.Error != nil {
		return models.ProductCategory{}, result.Error
	}
	return category, nil
}

func (d dataAccess) FindByName(name string) (models.ProductCategory, error) {
	var category models.ProductCategory
	result := d.db.Table(models.ProductCategory{}.TableName()).
		Where("name = ? AND del_flg = ?", name, false).
		First(&category)
	if result.Error != nil {
		return models.ProductCategory{}, result.Error
	}
	return category, nil
}

func (d dataAccess) FindCategoryProducts(id string) (models.ProductCategory, error) {
	var category models.ProductCategory
	result := d.db.Table(models.ProductCategory{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Products").
		First(&category)
	if result.Error != nil {
		return models.ProductCategory{}, result.Error
	}
	return category, nil
}

func (d dataAccess) Insert(item models.ProductCategory) (models.ProductCategory, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.ProductCategory{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.ProductCategory) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.ProductCategory
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.ProductCategory
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
