package cartdao

import (
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.Cart, error)
	FindById(id string) (models.Cart, error)
	FindByUserId(userId string) (models.Cart, error)
	Insert(item models.Cart) (models.Cart, error)
	Update(item models.Cart) error
	SoftDelete(id string) error
	Delete(id string) error
	// Cart item management
	FindItems(cartId string) ([]models.CartItem, error)
	AddItem(item models.CartItem) (models.CartItem, error)
	UpdateItem(item models.CartItem) error
	RemoveItem(itemId string) error
	ClearCart(cartId string) error
}

type dataAccess struct {
	db *gorm.DB
}

func New(db *gorm.DB) DataAccess {
	return dataAccess{
		db: db,
	}
}
func (d dataAccess) FindAll() ([]models.Cart, error) {
	var carts []models.Cart
	result := d.db.Table(models.Cart{}.TableName()).
		Where("del_flg = ?", false).
		Preload("User").
		Preload("Items").
		Find(&carts)
	if result.Error != nil {
		return []models.Cart{}, result.Error
	}
	return carts, nil
}

func (d dataAccess) FindById(id string) (models.Cart, error) {
	var cart models.Cart
	result := d.db.Table(models.Cart{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("User").
		Preload("Items.Product").
		First(&cart)
	if result.Error != nil {
		return models.Cart{}, result.Error
	}
	return cart, nil
}

func (d dataAccess) FindByUserId(userId string) (models.Cart, error) {
	var cart models.Cart
	result := d.db.Table(models.Cart{}.TableName()).
		Where("user_id = ? AND del_flg = ?", userId, false).
		Preload("User").
		Preload("Items.Product").
		First(&cart)
	if result.Error != nil {
		return models.Cart{}, result.Error
	}
	return cart, nil
}

func (d dataAccess) Insert(item models.Cart) (models.Cart, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.Cart{}, result.Error
	}
	return item, nil
}

func (d dataAccess) Update(item models.Cart) error {
	result := d.db.Table(item.TableName()).
		Where("id = ?", item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.Cart
	result := d.db.Table(item.TableName()).
		Where("id = ?", id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.Cart
	result := d.db.Table(item.TableName()).
		Where("id = ?", id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Cart item management
func (d dataAccess) FindItems(cartId string) ([]models.CartItem, error) {
	var items []models.CartItem
	result := d.db.Table(models.CartItem{}.TableName()).
		Where("cart_id = ? AND del_flg = ?", cartId, false).
		Preload("Product").
		Find(&items)
	if result.Error != nil {
		return []models.CartItem{}, result.Error
	}
	return items, nil
}

func (d dataAccess) AddItem(item models.CartItem) (models.CartItem, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.CartItem{}, result.Error
	}
	return item, nil
}

func (d dataAccess) UpdateItem(item models.CartItem) error {
	result := d.db.Table(item.TableName()).
		Where("id = ?", item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) RemoveItem(itemId string) error {
	var item models.CartItem
	result := d.db.Table(item.TableName()).
		Where("id = ?", itemId).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) ClearCart(cartId string) error {
	result := d.db.Table(models.CartItem{}.TableName()).
		Where("cart_id = ?", cartId).
		Delete(&models.CartItem{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
