package userDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	// Postgres Data Access Object Methods
	FindAll() ([]models.User, error)
	FindById(id string) (models.User, error)
	FindByName(name string) (models.User, error)
	FindUserProducts(id string) (models.User, error)
	FindUserOrders(id string) (models.User, error)
	FindUserReviews(id string) (models.User, error)
	FindUserAddresses(id string) (models.User, error)
	Insert(item models.User) (models.User, error)
	Update(item models.User) error
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

func (d dataAccess) FindAll() ([]models.User, error) {

	var users []models.User
	result := d.db.Table(models.User{}.TableName()).
		Where("del_flg = ?", false).
		Find(&users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}
	return users, nil
}

func (d dataAccess) FindById(id string) (models.User, error) {

	var user models.User
	result := d.db.Table(models.User{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (d dataAccess) FindByName(name string) (models.User, error) {

	var user models.User
	result := d.db.Table(models.User{}.TableName()).
		Where("name = ? AND del_flg = ?", name, false).
		First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (d dataAccess) FindUserProducts(id string) (models.User, error) {

	var user models.User
	result := d.db.Table(models.User{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Products").
		First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (d dataAccess) FindUserOrders(id string) (models.User, error) {

	var user models.User
	result := d.db.Table(models.User{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Orders").
		First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (d dataAccess) FindUserReviews(id string) (models.User, error) {

	var user models.User
	result := d.db.Table(models.User{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Reviews").
		First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (d dataAccess) FindUserAddresses(id string) (models.User, error) {

	var user models.User
	result := d.db.Table(models.User{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Addresses").
		First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (d dataAccess) Insert(item models.User) (models.User, error) {

	result := d.db.Table(item.TableName()).
		Create(&item)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return item, nil
}

const idWhere = "id = ? "

func (d dataAccess) Update(item models.User) error {

	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d dataAccess) SoftDelete(id string) error {

	var item models.User

	result := d.db.Table(item.TableName()).Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d dataAccess) Delete(id string) error {

	var item models.User

	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
