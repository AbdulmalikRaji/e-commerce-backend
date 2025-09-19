package models

import "time"

type ProductCategory struct {
	ID          string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string    `gorm:"unique;not null" json:"name"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now()" json:"updated_at"`
	DelFlg      bool      `gorm:"default:false" json:"del_flg"`

	// Relations
	Products []Product `gorm:"foreignKey:CategoryID"`
}

func (ProductCategory) TableName() string {
	return "public.product_categories"
}
