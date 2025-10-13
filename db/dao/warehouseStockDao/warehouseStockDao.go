package warehouseStockDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.WarehouseStock, error)
	FindById(id string) (models.WarehouseStock, error)
	FindByWarehouseId(warehouseId string) ([]models.WarehouseStock, error)
	FindByProductId(productId string) ([]models.WarehouseStock, error)
	Insert(item models.WarehouseStock) (models.WarehouseStock, error)
	Update(item models.WarehouseStock) error
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

func (d dataAccess) FindAll() ([]models.WarehouseStock, error) {
	var stocks []models.WarehouseStock
	result := d.db.Table(models.WarehouseStock{}.TableName()).
		Where("del_flg = ?", false).
		Preload("Warehouse").
		Preload("Product").
		Find(&stocks)
	if result.Error != nil {
		return []models.WarehouseStock{}, result.Error
	}
	return stocks, nil
}

func (d dataAccess) FindById(id string) (models.WarehouseStock, error) {
	var stock models.WarehouseStock
	result := d.db.Table(models.WarehouseStock{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Warehouse").
		Preload("Product").
		First(&stock)
	if result.Error != nil {
		return models.WarehouseStock{}, result.Error
	}
	return stock, nil
}

func (d dataAccess) FindByWarehouseId(warehouseId string) ([]models.WarehouseStock, error) {
	var stocks []models.WarehouseStock
	result := d.db.Table(models.WarehouseStock{}.TableName()).
		Where("warehouse_id = ? AND del_flg = ?", warehouseId, false).
		Preload("Warehouse").
		Preload("Product").
		Find(&stocks)
	if result.Error != nil {
		return []models.WarehouseStock{}, result.Error
	}
	return stocks, nil
}

func (d dataAccess) FindByProductId(productId string) ([]models.WarehouseStock, error) {
	var stocks []models.WarehouseStock
	result := d.db.Table(models.WarehouseStock{}.TableName()).
		Where("product_id = ? AND del_flg = ?", productId, false).
		Preload("Warehouse").
		Preload("Product").
		Find(&stocks)
	if result.Error != nil {
		return []models.WarehouseStock{}, result.Error
	}
	return stocks, nil
}

func (d dataAccess) Insert(item models.WarehouseStock) (models.WarehouseStock, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.WarehouseStock{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.WarehouseStock) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.WarehouseStock
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.WarehouseStock
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
