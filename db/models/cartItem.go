package models

import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CartID    uuid.UUID  `gorm:"type:uuid;index;not null" json:"cart_id"`
	ProductID uuid.UUID  `gorm:"type:uuid;index;not null" json:"product_id"`
	VariantID *uuid.UUID `gorm:"type:uuid;index" json:"variant_id"` // nullable if product has no variants
	Quantity  int        `gorm:"not null;default:1" json:"quantity"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg    bool       `gorm:"default:false" json:"del_flg"`

	// Relations
	Cart    Cart    `gorm:"foreignKey:CartID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cart,omitempty"`
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
}

func (CartItem) TableName() string {
	return "ecom.cart_items"
}
