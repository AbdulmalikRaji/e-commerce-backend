package models

import "time"

type Payment struct {
	ID             string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrderID        string    `gorm:"type:uuid;not null" json:"order_id"`
	Provider       string    `gorm:"type:varchar(50);not null" json:"provider"` // Stripe, Paystack, etc.
	Amount         float64   `gorm:"type:numeric(10,2);not null" json:"amount"`
	Status         string    `gorm:"type:varchar(20);not null" json:"status"` // initiated, successful, failed
	TransactionRef string    `gorm:"unique;not null" json:"transaction_ref"`
	CreatedAt      time.Time `gorm:"default:now()" json:"created_at"`
	DelFlg         bool      `gorm:"default:false" json:"del_flg"`

	// Relations
	Order Order `gorm:"foreignKey:OrderID"`
}

func (Payment) TableName() string {
	return "public.payments"
}
