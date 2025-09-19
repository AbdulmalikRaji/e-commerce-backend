package addressDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	// Postgres Data Access Object Methods
	FindAll() ([]models.Address, error)
	FindById(id string) (models.Address, error)
	FindByUserID(userID string) ([]models.Address, error)
	FindDefaultAddress(userID string) (models.Address, error)
	Insert(item models.Address) (models.Address, error)
	Update(item models.Address) error
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

func (d dataAccess) FindAll() ([]models.Address, error) {

	var addresses []models.Address
	result := d.db.Table(models.Address{}.TableName()).
		Where("del_flg = ?", false).
		Preload("User").
		Find(&addresses)
	if result.Error != nil {
		return []models.Address{}, result.Error
	}
	return addresses, nil
}

func (d dataAccess) FindById(id string) (models.Address, error) {

	var address models.Address
	result := d.db.Table(models.Address{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("User").
		First(&address)
	if result.Error != nil {
		return models.Address{}, result.Error
	}
	return address, nil
}

func (d dataAccess) FindByUserID(userID string) ([]models.Address, error) {

	var addresses []models.Address
	result := d.db.Table(models.Address{}.TableName()).
		Where("user_id = ? AND del_flg = ?", userID, false).
		Preload("User").
		Find(&addresses)
	if result.Error != nil {
		return []models.Address{}, result.Error
	}
	return addresses, nil
}

func (d dataAccess) FindDefaultAddress(userID string) (models.Address, error) {

	var address models.Address
	result := d.db.Table(models.Address{}.TableName()).
		Where("user_id = ? AND is_default = ? AND del_flg = ?", userID, true, false).
		Preload("User").
		First(&address)
	if result.Error != nil {
		return models.Address{}, result.Error
	}
	return address, nil
}

func (d dataAccess) Insert(item models.Address) (models.Address, error) {

	result := d.db.Table(item.TableName()).Create(&item)

	if result.Error != nil {
		return models.Address{}, result.Error
	}

	return item, nil
}

func (d dataAccess) Update(item models.Address) error {

	result := d.db.Table(item.TableName()).
		Where("id = ? ", item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d dataAccess) SoftDelete(id string) error {

	var item models.Address

	result := d.db.Table(item.TableName()).
		Where("id = ? ", id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d dataAccess) Delete(id string) error {

	var item models.Address

	result := d.db.Table(item.TableName()).
		Where("id = ? ", id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
