package models

import (
	"time"

	"github.com/google/uuid"
)

type WarehouseStock struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProductID   uuid.UUID `gorm:"type:uuid;index;not null" json:"product_id"`
	WarehouseID uuid.UUID `gorm:"type:uuid;index;not null" json:"warehouse_id"`
	Stock       int       `gorm:"not null" json:"stock"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg      bool      `gorm:"default:false" json:"del_flg"`

	Product   Product   `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"warehouse,omitempty"`
}

func (WarehouseStock) TableName() string {
	return "public.warehouse_stock"
}
