package models

import "github.com/google/uuid"

type Warehouse struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	StoreID   uuid.UUID  `gorm:"type:uuid;index;not null" json:"store_id"` // FK to Store
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	Address   string     `gorm:"type:text;not null" json:"address"`
	CreatedAt int64      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64      `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy *uuid.UUID `gorm:"type:uuid;index" json:"created_by,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid;index" json:"updated_by,omitempty"`
	DelFlg    bool       `gorm:"default:false" json:"del_flg"`

	Store          Store            `gorm:"foreignKey:StoreID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"store,omitempty"`
	WarehouseStock []WarehouseStock `gorm:"foreignKey:WarehouseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"warehouse_stock,omitempty"`
}

func (Warehouse) TableName() string {
	return "ecom.warehouses"
}
