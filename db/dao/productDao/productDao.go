package productDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	// Postgres Data Access Object Methods
	FindAll() ([]models.Product, error)
	FindById(id string) (models.Product, error)
	FindByName(name string) (models.Product, error)
	FindProductReviews(id string) (models.Product, error)
	Insert(item models.Product) (models.Product, error)
	Update(item models.Product) error
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

func (d dataAccess) FindAll() ([]models.Product, error) {

	var products []models.Product
	result := d.db.Table(models.Product{}.TableName()).
	Where("del_flg = ?", false).
	Preload("Category").
	Find(&products)
	if result.Error != nil {
		return []models.Product{}, result.Error
	}
	return products, nil
}

func (d dataAccess) FindById(id string) (models.Product, error) {

	var product models.Product
	result := d.db.Table(models.Product{}.TableName()).
	Where("id = ? AND del_flg = ?", id, false).
	Preload("Category").
	First(&product)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}

func (d dataAccess) FindByName(name string) (models.Product, error) {

	var product models.Product
	result := d.db.Table(models.Product{}.TableName()).
	Where("name = ? AND del_flg = ?", name, false).
	Preload("Category").
	First(&product)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}

func (d dataAccess) FindProductReviews(id string) (models.Product, error) {

	var product models.Product
	result := d.db.Table(models.Product{}.TableName()).
	Where("id = ? AND del_flg = ?", id, false).
	Preload("Category").
	Preload("Reviews").
	First(&product)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}

func (d dataAccess) Insert(item models.Product) (models.Product, error) {

	result := d.db.Table(item.TableName()).Create(&item)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return item, nil
}

func (d dataAccess) Update(item models.Product) error {

	result := d.db.Table(item.TableName()).
	Where("id = ? ", item.ID).
	Updates(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d dataAccess) SoftDelete(id string) error {

	var item models.Product

	result := d.db.Table(item.TableName()).
	Where("id = ? ", id).
	Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d dataAccess) Delete(id string) error {

	var item models.Product

	result := d.db.Table(item.TableName()).
	Where("id = ? ", id).
	Delete(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
