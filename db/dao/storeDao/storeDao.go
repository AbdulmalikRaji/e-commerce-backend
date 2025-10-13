package storeDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.Store, error)
	FindById(id string) (models.Store, error)
	FindByOwnerId(ownerId string) ([]models.Store, error)
	FindStoreProducts(storeId string) ([]models.Product, error)
	Insert(item models.Store) (models.Store, error)
	Update(item models.Store) error
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

func (d dataAccess) FindAll() ([]models.Store, error) {
	var stores []models.Store
	result := d.db.Table(models.Store{}.TableName()).
		Where("del_flg = ?", false).
		Preload("Owner").
		Preload("Products").
		Find(&stores)
	if result.Error != nil {
		return []models.Store{}, result.Error
	}
	return stores, nil
}

func (d dataAccess) FindById(id string) (models.Store, error) {
	var store models.Store
	result := d.db.Table(models.Store{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Owner").
		Preload("Products").
		First(&store)
	if result.Error != nil {
		return models.Store{}, result.Error
	}
	return store, nil
}

func (d dataAccess) FindByOwnerId(ownerId string) ([]models.Store, error) {
	var stores []models.Store
	result := d.db.Table(models.Store{}.TableName()).
		Where("owner_id = ? AND del_flg = ?", ownerId, false).
		Preload("Owner").
		Preload("Products").
		Find(&stores)
	if result.Error != nil {
		return []models.Store{}, result.Error
	}
	return stores, nil
}

func (d dataAccess) FindStoreProducts(storeId string) ([]models.Product, error) {
	var products []models.Product
	result := d.db.Table(models.Product{}.TableName()).
		Where("store_id = ? AND del_flg = ?", storeId, false).
		Preload("Category").
		Preload("Images").
		Preload("Variants").
		Preload("WarehouseStock").
		Preload("Tags").
		Preload("SubCategories").
		Find(&products)
	if result.Error != nil {
		return []models.Product{}, result.Error
	}
	return products, nil
}

func (d dataAccess) Insert(item models.Store) (models.Store, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.Store{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.Store) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.Store
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.Store
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
