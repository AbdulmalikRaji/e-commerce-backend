package models

import (
	"time"

	"github.com/google/uuid"
)

type UserToken struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	RefreshToken string    `gorm:"type:text;not null" json:"refresh_token"`
	IsRevoked    bool      `gorm:"default:false" json:"is_revoked"`
	ExpiresAt    time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg       bool      `gorm:"default:false" json:"del_flg"`
}

func (UserToken) TableName() string {
	return "ecom.user_tokens"
}
