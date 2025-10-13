package tagDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.Tag, error)
	FindById(id string) (models.Tag, error)
	FindByName(name string) (models.Tag, error)
	FindTagProducts(tagId string) ([]models.Product, error)
	Insert(item models.Tag) (models.Tag, error)
	Update(item models.Tag) error
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

func (d dataAccess) FindAll() ([]models.Tag, error) {
	var tags []models.Tag
	result := d.db.Table(models.Tag{}.TableName()).
		Where("del_flg = ?", false).
		Preload("Products").
		Find(&tags)
	if result.Error != nil {
		return []models.Tag{}, result.Error
	}
	return tags, nil
}

func (d dataAccess) FindById(id string) (models.Tag, error) {
	var tag models.Tag
	result := d.db.Table(models.Tag{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Products").
		First(&tag)
	if result.Error != nil {
		return models.Tag{}, result.Error
	}
	return tag, nil
}

func (d dataAccess) FindByName(name string) (models.Tag, error) {
	var tag models.Tag
	result := d.db.Table(models.Tag{}.TableName()).
		Where("name = ? AND del_flg = ?", name, false).
		Preload("Products").
		First(&tag)
	if result.Error != nil {
		return models.Tag{}, result.Error
	}
	return tag, nil
}

func (d dataAccess) FindTagProducts(tagId string) ([]models.Product, error) {
	var tag models.Tag
	result := d.db.Table(models.Tag{}.TableName()).
		Where("id = ? AND del_flg = ?", tagId, false).
		Preload("Products.Category").
		Preload("Products.Images").
		Preload("Products.Variants").
		Preload("Products.WarehouseStock").
		Preload("Products.Tags").
		Preload("Products.SubCategories").
		First(&tag)
	if result.Error != nil {
		return nil, result.Error
	}
	return tag.Products, nil
}

func (d dataAccess) Insert(item models.Tag) (models.Tag, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.Tag{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.Tag) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.Tag
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.Tag
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
