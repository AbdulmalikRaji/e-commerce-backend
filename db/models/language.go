package models

import (
	"time"

	"github.com/google/uuid"
)

type Language struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Code      string    `gorm:"type:varchar(10);not null;unique" json:"code"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Locale    *string   `gorm:"type:varchar(20)" json:"locale,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg    bool      `gorm:"default:false" json:"del_flg"`
}

func (Language) TableName() string {
	return "ecom.languages"
}
