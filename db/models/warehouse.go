package models

import "github.com/google/uuid"

type Warehouse struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Address   string    `gorm:"type:text;not null" json:"address"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64     `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg    bool      `gorm:"default:false" json:"del_flg"`

	WarehouseStock []WarehouseStock `gorm:"foreignKey:WarehouseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"warehouse_stock,omitempty"`
}

func (Warehouse) TableName() string {
	return "public.warehouses"
}
