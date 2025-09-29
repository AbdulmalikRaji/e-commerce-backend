package categoryDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.Category, error)
	FindById(id string) (models.Category, error)
	FindByName(name string) (models.Category, error)
	FindChildren(parentId string) ([]models.Category, error)
	FindParent(id string) (models.Category, error)
	Insert(item models.Category) (models.Category, error)
	Update(item models.Category) error
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

func (d dataAccess) FindAll() ([]models.Category, error) {
	var categories []models.Category
	result := d.db.Table(models.Category{}.TableName()).
		Where("del_flg = ?", false).
		Find(&categories)
	if result.Error != nil {
		return []models.Category{}, result.Error
	}
	return categories, nil
}

func (d dataAccess) FindById(id string) (models.Category, error) {
	var category models.Category
	result := d.db.Table(models.Category{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		First(&category)
	if result.Error != nil {
		return models.Category{}, result.Error
	}
	return category, nil
}

func (d dataAccess) FindByName(name string) (models.Category, error) {
	var category models.Category
	result := d.db.Table(models.Category{}.TableName()).
		Where("name = ? AND del_flg = ?", name, false).
		First(&category)
	if result.Error != nil {
		return models.Category{}, result.Error
	}
	return category, nil
}

func (d dataAccess) FindChildren(parentId string) ([]models.Category, error) {
	var categories []models.Category
	result := d.db.Table(models.Category{}.TableName()).
		Where("parent_id = ? AND del_flg = ?", parentId, false).
		Find(&categories)
	if result.Error != nil {
		return []models.Category{}, result.Error
	}
	return categories, nil
}

func (d dataAccess) FindParent(id string) (models.Category, error) {
	var category models.Category
	result := d.db.Table(models.Category{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Parent").
		First(&category)
	if result.Error != nil {
		return models.Category{}, result.Error
	}
	if category.Parent == nil {
		return models.Category{}, gorm.ErrRecordNotFound
	}
	return *category.Parent, nil
}

func (d dataAccess) Insert(item models.Category) (models.Category, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.Category{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.Category) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.Category
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.Category
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
