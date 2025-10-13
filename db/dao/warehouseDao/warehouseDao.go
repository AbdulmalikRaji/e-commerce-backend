package warehouseDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.Warehouse, error)
	FindById(id string) (models.Warehouse, error)
	FindByAddress(address string) ([]models.Warehouse, error)
	FindWarehouseStocks(warehouseId string) (models.Warehouse, error)
	Insert(item models.Warehouse) (models.Warehouse, error)
	Update(item models.Warehouse) error
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

func (d dataAccess) FindAll() ([]models.Warehouse, error) {
	var warehouses []models.Warehouse
	result := d.db.Table(models.Warehouse{}.TableName()).
		Where("del_flg = ?", false).
		Preload("Stocks").
		Find(&warehouses)
	if result.Error != nil {
		return []models.Warehouse{}, result.Error
	}
	return warehouses, nil
}

func (d dataAccess) FindById(id string) (models.Warehouse, error) {
	var warehouse models.Warehouse
	result := d.db.Table(models.Warehouse{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Stocks").
		First(&warehouse)
	if result.Error != nil {
		return models.Warehouse{}, result.Error
	}
	return warehouse, nil
}

func (d dataAccess) FindByAddress(address string) ([]models.Warehouse, error) {
	var warehouses []models.Warehouse
	result := d.db.Table(models.Warehouse{}.TableName()).
		Where("address = ? AND del_flg = ?", address, false).
		Preload("Stocks").
		Find(&warehouses)
	if result.Error != nil {
		return []models.Warehouse{}, result.Error
	}
	return warehouses, nil
}

func (d dataAccess) FindWarehouseStocks(warehouseId string) (models.Warehouse, error) {
	var warehouse models.Warehouse
	result := d.db.Table(models.Warehouse{}.TableName()).
		Where("id = ? AND del_flg = ?", warehouseId, false).
		Preload("Stocks.Product").
		First(&warehouse)
	if result.Error != nil {
		return models.Warehouse{}, result.Error
	}
	return warehouse, nil
}

func (d dataAccess) Insert(item models.Warehouse) (models.Warehouse, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.Warehouse{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.Warehouse) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.Warehouse
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.Warehouse
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
