package storeUserDao

import (
	"errors"
	"log"

	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.StoreUser, error)
	FindById(id string) (models.StoreUser, error)
	FindByStoreID(storeId string) ([]models.StoreUser, error)
	FindByUserId(userId string) ([]models.StoreUser, error)
	FindByStoreAndUser(storeID, userID uuid.UUID) (models.StoreUser, error)
	Insert(item models.StoreUser) (models.StoreUser, error)
	Update(item models.StoreUser) error
	SoftDelete(id string) error
	Delete(id string) error
	HasPermission(storeID, userID uuid.UUID, action string) bool
}

type dataAccess struct {
	db *gorm.DB
}

func New(client connection.Client) DataAccess {
	return dataAccess{
		db: client.PostgresConnection,
	}
}

func (d dataAccess) FindAll() ([]models.StoreUser, error) {
	var items []models.StoreUser
	result := d.db.Table(models.StoreUser{}.TableName()).
		Where("del_flg = ?", false).
		Preload("User").
		Preload("Store").
		Find(&items)
	if result.Error != nil {
		return []models.StoreUser{}, result.Error
	}
	return items, nil
}

func (d dataAccess) FindById(id string) (models.StoreUser, error) {
	var item models.StoreUser
	result := d.db.Table(models.StoreUser{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("User").
		Preload("Store").
		First(&item)
	if result.Error != nil {
		return models.StoreUser{}, result.Error
	}
	return item, nil
}

func (d dataAccess) FindByStoreID(storeId string) ([]models.StoreUser, error) {
	var items []models.StoreUser
	result := d.db.Table(models.StoreUser{}.TableName()).
		Where("store_id = ? AND del_flg = ?", storeId, false).
		Preload("User").
		Find(&items)
	if result.Error != nil {
		return []models.StoreUser{}, result.Error
	}
	return items, nil
}

func (d dataAccess) FindByUserId(userId string) ([]models.StoreUser, error) {
	var items []models.StoreUser
	result := d.db.Table(models.StoreUser{}.TableName()).
		Where("user_id = ? AND del_flg = ?", userId, false).
		Preload("Store").
		Find(&items)
	if result.Error != nil {
		return []models.StoreUser{}, result.Error
	}
	return items, nil
}

func (d dataAccess) FindByStoreAndUser(storeID, userID uuid.UUID) (models.StoreUser, error) {
	var item models.StoreUser
	result := d.db.Table(models.StoreUser{}.TableName()).
		Where("store_id = ? AND user_id = ? AND del_flg = ?", storeID, userID, false).
		Preload("User").
		Preload("Store").
		First(&item)
	if result.Error != nil {
		return models.StoreUser{}, result.Error
	}
	return item, nil
}

func (d dataAccess) Insert(item models.StoreUser) (models.StoreUser, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.StoreUser{}, result.Error
	}
	return item, nil
}

const suIdWhere = "id = ? "

func (d dataAccess) Update(item models.StoreUser) error {
	result := d.db.Table(item.TableName()).
		Where(suIdWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.StoreUser
	result := d.db.Table(item.TableName()).
		Where(suIdWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.StoreUser
	result := d.db.Table(item.TableName()).
		Where(suIdWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Action constants
const (
	ActionAddProduct          = "add_product"
	ActionUpdateProduct       = "update_product"
	ActionDeleteProduct       = "delete_product"
	ActionManageOrders        = "manage_orders"
	ActionManageStoreSettings = "manage_store_settings"
)

func (d dataAccess) HasPermission(storeID, userID uuid.UUID, action string) bool {
	var su models.StoreUser
	res := d.db.Table(models.StoreUser{}.TableName()).
		Where("store_id = ? AND user_id = ? AND del_flg = ?", storeID, userID, false).
		First(&su)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false
		}
		log.Printf("storeUserDao: failed to load StoreUser store=%s user=%s err=%v", storeID.String(), userID.String(), res.Error)
		return false
	}

	// role defaults
	var rd map[string]bool
	switch su.Role {
	case models.RoleManager:
		rd = map[string]bool{
			ActionAddProduct:          true,
			ActionUpdateProduct:       true,
			ActionDeleteProduct:       true,
			ActionManageOrders:        true,
			ActionManageStoreSettings: true,
		}
	case models.RoleWorker:
		rd = map[string]bool{
			ActionManageOrders: true,
		}
	default:
		rd = map[string]bool{}
	}

	switch action {
	case ActionAddProduct:
		return su.CanAddProducts || rd[ActionAddProduct]
	case ActionUpdateProduct:
		return su.CanUpdateProducts || rd[ActionUpdateProduct]
	case ActionDeleteProduct:
		return su.CanDeleteProducts || rd[ActionDeleteProduct]
	case ActionManageOrders:
		return su.CanManageOrders || rd[ActionManageOrders]
	case ActionManageStoreSettings:
		return su.CanManageStoreSettings || rd[ActionManageStoreSettings]
	default:
		return false
	}
}
