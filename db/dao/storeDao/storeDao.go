package storeDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.Store, error)
	FindById(id string) (models.Store, error)
	FindByOwnerID(ownerId string) ([]models.Store, error)
	FindByName(name string) ([]models.Store, error)
	FindStoreProducts(storeId string) (models.Store, error)
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

func (d dataAccess) FindByOwnerID(ownerId string) ([]models.Store, error) {
	//todo: figure if owner can have multiple stores
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

func (d dataAccess) FindByName(name string) ([]models.Store, error) {
	var stores []models.Store
	result := d.db.Table(models.Store{}.TableName()).
		Where("name ILIKE ? AND del_flg = ?", "%"+name+"%", false).
		Preload("Owner").
		Preload("Products").
		Find(&stores)
	if result.Error != nil {
		return []models.Store{}, result.Error
	}
	return stores, nil
}

func (d dataAccess) FindStoreProducts(storeId string) (models.Store, error) {
	var store models.Store
	result := d.db.Table(models.Store{}.TableName()).
		Where("id = ? AND del_flg = ?", storeId, false).
		Preload("Products", "del_flg = ?", false).
		First(&store)
	if result.Error != nil {
		return models.Store{}, result.Error
	}
	return store, nil
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
