package models

import "time"

type Product struct {
	ID            string  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	SellerID      string  `gorm:"type:uuid;not null" json:"seller_id"`
	CategoryID    *string `gorm:"type:uuid" json:"category_id"`
	Name          string  `gorm:"type:varchar(100);not null" json:"name"`
	Description   string  `gorm:"type:text" json:"description"`
	Price         float64 `gorm:"type:numeric(10,2);not null" json:"price"`
	Stock         int     `gorm:"default:0" json:"stock"`
	Barcode       string  `gorm:"type:varchar(64);uniqueIndex" json:"barcode"`
	RatingAverage float64 `gorm:"type:numeric(3,2);default:0" json:"rating_average"`
	RatingCount   int     `gorm:"default:0" json:"rating_count"`

	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now()" json:"updated_at"`
	DelFlg        bool    `gorm:"default:false" json:"del_flg"`

	// Relations
	Seller   User             `gorm:"foreignKey:SellerID"`
	Category *ProductCategory `gorm:"foreignKey:CategoryID"`
	Reviews  []Review         `gorm:"foreignKey:ProductID"`
}

func (Product) TableName() string {
	return "public.products"
}
