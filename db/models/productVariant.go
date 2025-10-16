package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductVariant struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProductID      uuid.UUID  `gorm:"type:uuid;index;not null" json:"product_id"`
	SKU            string     `gorm:"type:varchar(64);uniqueIndex;not null" json:"sku"`
	AttributeName  string     `gorm:"type:varchar(50)" json:"attribute_name"` // e.g., size, color, material
	AttributeValue string     `gorm:"type:varchar(50)" json:"attribute_value"`
	PriceOverride  *float64   `gorm:"type:numeric(10,2)" json:"price_override"` // nullable
	Stock          int        `gorm:"default:0" json:"stock"`
	CreatedAt      time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy      *uuid.UUID `gorm:"type:uuid;index" json:"created_by,omitempty"`
	UpdatedBy      *uuid.UUID `gorm:"type:uuid;index" json:"updated_by,omitempty"`
	DelFlg         bool       `gorm:"default:false" json:"del_flg"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
}

func (ProductVariant) TableName() string {
	return "ecom.product_variants"
}
