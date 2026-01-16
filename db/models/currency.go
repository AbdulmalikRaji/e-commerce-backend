package models

import (
	"time"

	"github.com/google/uuid"
)

type Currency struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Code          string    `gorm:"type:varchar(3);not null;unique" json:"code"`
	Name          string    `gorm:"type:varchar(100);not null" json:"name"`
	Symbol        *string   `gorm:"type:varchar(10)" json:"symbol,omitempty"`
	DecimalPlaces int       `gorm:"type:int;default:2" json:"decimal_places"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg        bool      `gorm:"default:false" json:"del_flg"`
}

func (Currency) TableName() string {
	return "ecom.currencies"
}
