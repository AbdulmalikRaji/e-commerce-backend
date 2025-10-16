package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string     `gorm:"type:varchar(100);not null;unique" json:"name"`
	Description string     `gorm:"type:text" json:"description,omitempty"`
	ParentID    *uuid.UUID `gorm:"type:uuid;index" json:"parent_id,omitempty"`
	DelFlg      bool       `gorm:"default:false" json:"del_flg"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy   *uuid.UUID `gorm:"type:uuid;index" json:"created_by,omitempty"`
	UpdatedBy   *uuid.UUID `gorm:"type:uuid;index" json:"updated_by,omitempty"`

	// Relations
	Parent      *Category  `gorm:"foreignKey:ParentID;references:ID" json:"parent,omitempty"`
	Children    []Category `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Products    []Product  `gorm:"foreignKey:CategoryID" json:"products,omitempty"` // main category products
	SubProducts []Product  `gorm:"many2many:product_subcategories;joinForeignKey:SubcategoryID;joinReferences:ProductID" json:"sub_products,omitempty"`
}

func (Category) TableName() string {
	return "ecom.categories"
}
