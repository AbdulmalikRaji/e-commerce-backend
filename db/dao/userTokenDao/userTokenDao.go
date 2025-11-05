package userTokenDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	// Postgres Data Access Object Methods
	FindById(id string) (models.UserToken, error)
	FindByUserID(userID string) ([]models.UserToken, error)
	FindValidTokens(userID string) ([]models.UserToken, error)
	Insert(item models.UserToken) (models.UserToken, error)
	Update(item models.UserToken) error
	FindByRefreshToken(refreshToken string) (models.UserToken, error)
	RevokeToken(id string) error
	DeleteExpiredTokens() error
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

func (d dataAccess) FindById(id string) (models.UserToken, error) {
	var token models.UserToken
	result := d.db.Table(models.UserToken{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		First(&token)
	if result.Error != nil {
		return models.UserToken{}, result.Error
	}
	return token, nil
}

func (d dataAccess) FindByUserID(userID string) ([]models.UserToken, error) {
	var tokens []models.UserToken
	result := d.db.Table(models.UserToken{}.TableName()).
		Where("user_id = ? AND del_flg = ?", userID, false).
		Find(&tokens)
	if result.Error != nil {
		return nil, result.Error
	}
	return tokens, nil
}

func (d dataAccess) FindValidTokens(userID string) ([]models.UserToken, error) {
	var tokens []models.UserToken
	result := d.db.Table(models.UserToken{}.TableName()).
		Where("user_id = ? AND del_flg = ? AND is_revoked = ? AND expires_at > NOW()",
			userID, false, false).
		Find(&tokens)
	if result.Error != nil {
		return nil, result.Error
	}
	return tokens, nil
}

func (d dataAccess) Insert(item models.UserToken) (models.UserToken, error) {
	result := d.db.Table(item.TableName()).
		Create(&item)
	if result.Error != nil {
		return models.UserToken{}, result.Error
	}
	return item, nil
}

const idWhere = "id = ?"

func (d dataAccess) Update(item models.UserToken) error {
	result := d.db.Table(item.TableName()).
		Where(idWhere, item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) FindByRefreshToken(refreshToken string) (models.UserToken, error) {
	var token models.UserToken
	result := d.db.Table(models.UserToken{}.TableName()).
		Where("refresh_token = ? AND del_flg = ?", refreshToken, false).
		First(&token)
	if result.Error != nil {
		return models.UserToken{}, result.Error
	}
	return token, nil
}

func (d dataAccess) RevokeToken(id string) error {
	result := d.db.Table(models.UserToken{}.TableName()).
		Where(idWhere, id).
		Update("is_revoked", true)
	if result.Error != nil {
		return result.Error
	}
		return nil
	}

func (d dataAccess) DeleteExpiredTokens() error {
	result := d.db.Table(models.UserToken{}.TableName()).
		Where("expires_at < NOW() OR (is_revoked = ? AND updated_at < NOW() - INTERVAL '24 hours')", true).
		Delete(&models.UserToken{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) SoftDelete(id string) error {
	result := d.db.Table(models.UserToken{}.TableName()).
		Where(idWhere, id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d dataAccess) Delete(id string) error {
	var item models.UserToken
	result := d.db.Table(item.TableName()).
		Where(idWhere, id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
