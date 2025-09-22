package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID                uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	BuyerID           uuid.UUID  `gorm:"type:uuid;index;not null" json:"buyer_id"`
	ShippingAddressID *uuid.UUID `gorm:"type:uuid;index" json:"shipping_address_id"`
	Status            string     `gorm:"type:varchar(20);not null;default:'pending'" json:"status"` // pending | paid | shipped | delivered | cancelled
	TotalAmount       float64    `gorm:"type:numeric(10,2)" json:"total_amount"`
	CreatedAt         time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg            bool       `gorm:"default:false" json:"del_flg"`

	// Relations
	Buyer           User        `gorm:"foreignKey:BuyerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"buyer,omitempty"`
	ShippingAddress *Address    `gorm:"foreignKey:ShippingAddressID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"shipping_address,omitempty"`
	Items           []OrderItem `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items,omitempty"`
	Payments        []Payment   `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"payments,omitempty"`
}

func (Order) TableName() string {
	return "public.orders"
}
