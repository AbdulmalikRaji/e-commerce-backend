package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductSubcategory struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProductID     uuid.UUID `gorm:"type:uuid;index;not null" json:"product_id"`
	SubcategoryID uuid.UUID `gorm:"type:uuid;index;not null" json:"subcategory_id"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relations
	Product     Product  `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
	Subcategory Category `gorm:"foreignKey:SubcategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"subcategory,omitempty"`
}

func (ProductSubcategory) TableName() string {
	return "public.product_subcategories"
}
