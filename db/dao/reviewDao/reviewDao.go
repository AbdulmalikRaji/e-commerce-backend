package reviewDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.Review, error)
	FindById(id string) (models.Review, error)
	FindByProductId(productId string) ([]models.Review, error)
	FindByUserId(userId string) ([]models.Review, error)
	FindByOrderItemId(orderItemId string) (models.Review, error)
	Insert(item models.Review) (models.Review, error)
	Update(item models.Review) error
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

func (d dataAccess) FindAll() ([]models.Review, error) {
	var reviews []models.Review
	result := d.db.Table(models.Review{}.TableName()).
		Where("del_flg = ?", false).
		Preload("User").
		Preload("Product").
		Find(&reviews)
	if result.Error != nil {
		return []models.Review{}, result.Error
	}
	return reviews, nil
}

func (d dataAccess) FindById(id string) (models.Review, error) {
	var review models.Review
	result := d.db.Table(models.Review{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("User").
		Preload("Product").
		First(&review)
	if result.Error != nil {
		return models.Review{}, result.Error
	}
	return review, nil
}

func (d dataAccess) FindByProductId(productId string) ([]models.Review, error) {
	var reviews []models.Review
	result := d.db.Table(models.Review{}.TableName()).
		Where("product_id = ? AND del_flg = ?", productId, false).
		Preload("User").
		Preload("Product").
		Find(&reviews)
	if result.Error != nil {
		return []models.Review{}, result.Error
	}
	return reviews, nil
}

func (d dataAccess) FindByUserId(userId string) ([]models.Review, error) {
	var reviews []models.Review
	result := d.db.Table(models.Review{}.TableName()).
		Where("user_id = ? AND del_flg = ?", userId, false).
		Preload("User").
		Preload("Product").
		Find(&reviews)
	if result.Error != nil {
		return []models.Review{}, result.Error
	}
	return reviews, nil
}

func (d dataAccess) FindByOrderItemId(orderItemId string) (models.Review, error) {
	var review models.Review
	result := d.db.Table(models.Review{}.TableName()).
		Where("order_item_id = ? AND del_flg = ?", orderItemId, false).
		Preload("User").
		Preload("Product").
		First(&review)
	if result.Error != nil {
		return models.Review{}, result.Error
	}
	return review, nil
}

func (d dataAccess) Insert(item models.Review) (models.Review, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.Review{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.Review) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	var item models.Review
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.Review
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
