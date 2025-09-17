package models

import "time"

type OrderItem struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrderID   string    `gorm:"type:uuid;not null" json:"order_id"`
	ProductID string    `gorm:"type:uuid;not null" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	UnitPrice float64   `gorm:"type:numeric(10,2);not null" json:"unit_price"`
	ReviewID  *string   `gorm:"type:uuid;unique" json:"review_id"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`

	// Relations
	Order   Order   `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product Product `gorm:"foreignKey:ProductID"`
	Review  *Review `gorm:"foreignKey:ReviewID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (OrderItem) TableName() string {
	return "public.order_items"
}
