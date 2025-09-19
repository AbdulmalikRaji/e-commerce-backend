package models

import "time"

type Address struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID     string    `gorm:"type:uuid;not null" json:"user_id"`
	Line1      string    `gorm:"type:varchar(255);not null" json:"address_line1"`
	Line2      string    `gorm:"type:varchar(255)" json:"address_line2"`
	City       string    `gorm:"type:varchar(100);not null" json:"city"`
	State      string    `gorm:"type:varchar(100);not null" json:"state"`
	PostalCode string    `gorm:"type:varchar(20);not null" json:"postal_code"`
	Country    string    `gorm:"type:varchar(100);not null" json:"country"`
	IsDefault  bool      `gorm:"default:false" json:"is_default"`
	CreatedAt  time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:now()" json:"updated_at"`
	DelFlg     bool      `gorm:"default:false" json:"del_flg"`

	// Relations
	User User `gorm:"foreignKey:UserID"`
}

func (Address) TableName() string {
	return "public.addresses"
}
