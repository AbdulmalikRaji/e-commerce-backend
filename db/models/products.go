package models

import "time"

type Product struct {
	ID          string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	SellerID    string    `gorm:"type:uuid;not null" json:"seller_id"`
	CategoryID  *string   `gorm:"type:uuid" json:"category_id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"type:numeric(10,2);not null" json:"price"`
	Stock       int       `gorm:"default:0" json:"stock"`
	CreatedAt   time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now()" json:"updated_at"`

	// Relations
	Seller   User             `gorm:"foreignKey:SellerID"`
	Category *ProductCategory `gorm:"foreignKey:CategoryID"`
}

func (Product) TableName() string {
	return "public.products"
}