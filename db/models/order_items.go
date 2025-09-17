package models

import "time"

type OrderItem struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrderID   string    `gorm:"type:uuid;not null" json:"order_id"`
	ProductID string    `gorm:"type:uuid;not null" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	UnitPrice float64   `gorm:"type:numeric(10,2);not null" json:"unit_price"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`

	// Relations
	Order   Order   `gorm:"foreignKey:OrderID"`
	Product Product `gorm:"foreignKey:ProductID"`
}

func (OrderItem) TableName() string {
	return "public.order_items"
}
