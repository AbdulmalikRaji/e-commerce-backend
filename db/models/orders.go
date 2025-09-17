package models

import "time"

type Order struct {
	ID                string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	BuyerID           string    `gorm:"type:uuid;not null" json:"buyer_id"`
	ShippingAddressID *string   `gorm:"type:uuid" json:"shipping_address_id"`
	Status            string    `gorm:"type:varchar(20);default:'pending'" json:"status"`
	TotalAmount       float64   `gorm:"type:numeric(10,2);default:0" json:"total_amount"`
	CreatedAt         time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt         time.Time `gorm:"default:now()" json:"updated_at"`

	// Relations
	Buyer           User        `gorm:"foreignKey:BuyerID"`
	ShippingAddress *Address    `gorm:"foreignKey:ShippingAddressID"`
	Items           []OrderItem `gorm:"foreignKey:OrderID"`
	Payments        []Payment   `gorm:"foreignKey:OrderID"`
}

func (Order) TableName() string {
	return "public.orders"
}