package models

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	Line1      string    `gorm:"type:varchar(255);not null" json:"line1"`
	Line2      string    `gorm:"type:varchar(255)" json:"line2"`
	City       string    `gorm:"type:varchar(100);not null" json:"city"`
	State      string    `gorm:"type:varchar(100);not null" json:"state"`
	PostalCode string    `gorm:"type:varchar(20);not null" json:"postal_code"`
	Country    string    `gorm:"type:varchar(100);not null" json:"country"`
	IsDefault  bool      `gorm:"default:false" json:"is_default"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg     bool      `gorm:"default:false" json:"del_flg"`

	// Relations
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}

func (Address) TableName() string {
	return "public.addresses"
}
