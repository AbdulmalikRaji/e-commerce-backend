package notificationDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	FindAll() ([]models.Notification, error)
	FindById(id string) (models.Notification, error)
	FindByUserId(userId string) ([]models.Notification, error)
	FindUserUnreadNotifications(userId string) ([]models.Notification, error)
	Insert(item models.Notification) (models.Notification, error)
	Update(item models.Notification) error
	MarkAsRead(id string) error
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

func (d dataAccess) FindAll() ([]models.Notification, error) {
	var notifications []models.Notification
	result := d.db.Table(models.Notification{}.TableName()).
		Find(&notifications)
	if result.Error != nil {
		return []models.Notification{}, result.Error
	}
	return notifications, nil
}

func (d dataAccess) FindById(id string) (models.Notification, error) {
	var notification models.Notification
	result := d.db.Table(models.Notification{}.TableName()).
		Where("id = ?", id).
		First(&notification)
	if result.Error != nil {
		return models.Notification{}, result.Error
	}
	return notification, nil
}

func (d dataAccess) FindByUserId(userId string) ([]models.Notification, error) {
	var notifications []models.Notification
	result := d.db.Table(models.Notification{}.TableName()).
		Where("user_id = ?", userId).
		Find(&notifications)
	if result.Error != nil {
		return []models.Notification{}, result.Error
	}
	return notifications, nil
}

func (d dataAccess) FindUserUnreadNotifications(userId string) ([]models.Notification, error) {
	var notifications []models.Notification
	result := d.db.Table(models.Notification{}.TableName()).
		Where("user_id = ? AND is_read = ?", userId, false).
		Find(&notifications)
	if result.Error != nil {
		return []models.Notification{}, result.Error
	}
	return notifications, nil
}

func (d dataAccess) Insert(item models.Notification) (models.Notification, error) {
	result := d.db.Table(item.TableName()).Create(&item)
	if result.Error != nil {
		return models.Notification{}, result.Error
	}
	return item, nil
}

func (d dataAccess) Update(item models.Notification) error {
	result := d.db.Table(item.TableName()).
		Where("id = ?", item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) MarkAsRead(id string) error {
	result := d.db.Table(models.Notification{}.TableName()).
		Where("id = ?", id).
		Update("is_read", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.Notification
	result := d.db.Table(item.TableName()).
		Where("id = ?", id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
