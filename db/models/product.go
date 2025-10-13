package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	StoreID       uuid.UUID `gorm:"type:uuid;index;not null" json:"store_id"`
	CategoryID    uuid.UUID `gorm:"type:uuid;index;not null" json:"category_id"` // main category
	Name          string    `gorm:"type:varchar(100);not null" json:"name"`
	Description   string    `gorm:"type:text" json:"description"`
	Price         float64   `gorm:"type:numeric(10,2);not null" json:"price"`
	Stock         int       `gorm:"default:0" json:"stock"`
	HasVariants   bool      `gorm:"default:false" json:"has_variants"`
	IsDiscounted  bool      `gorm:"default:false" json:"is_discounted"`
	DiscountPct   float64   `gorm:"type:numeric(5,2);default:0" json:"discount_percent"`
	IsPopular     bool      `gorm:"default:false" json:"is_popular"`
	Barcode       string    `gorm:"type:varchar(64);uniqueIndex" json:"barcode"`
	RatingAverage float64   `gorm:"type:numeric(3,2);default:0" json:"rating_average"`
	RatingCount   int       `gorm:"default:0" json:"rating_count"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg        bool      `gorm:"default:false" json:"del_flg"`

	// Relations
	Store          Store            `gorm:"foreignKey:StoreID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"store,omitempty"`
	Category       Category         `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"`
	Images         []ProductImage   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"images,omitempty"`
	Variants       []ProductVariant `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"variants,omitempty"`
	SubCategories  []Category       `gorm:"many2many:product_subcategories;joinForeignKey:ProductID;joinReferences:SubcategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"subcategories,omitempty"`
	Reviews        []Review         `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"reviews,omitempty"`
	WarehouseStock []WarehouseStock `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"warehouse_stock,omitempty"`
	Tags           []Tag            `gorm:"many2many:product_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tags,omitempty"`
}

func (Product) TableName() string {
	return "public.products"
}
