package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderItem struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrderID   uuid.UUID  `gorm:"type:uuid;index;not null" json:"order_id"`
	ProductID uuid.UUID  `gorm:"type:uuid;index;not null" json:"product_id"`
	VariantID *uuid.UUID `gorm:"type:uuid;index" json:"variant_id"`
	Quantity  int        `gorm:"not null" json:"quantity"`
	UnitPrice float64    `gorm:"type:numeric(10,2);not null" json:"unit_price"` // snapshot at purchase
	ReviewID  *uuid.UUID `gorm:"type:uuid;index" json:"review_id"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	DelFlg    bool       `gorm:"default:false" json:"del_flg"`

	// Relations
	Order   Order           `gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order,omitempty"`
	Product Product         `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
	Review  *Review         `gorm:"foreignKey:ReviewID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Variant *ProductVariant `gorm:"foreignKey:VariantID;references:ID" json:"variant,omitempty"`
	// Review relation is optional and modeled in Review struct (Review.OrderItemID -> OrderItem.ID) if desired
}

func (OrderItem) TableName() string {
	return "public.order_items"
}
