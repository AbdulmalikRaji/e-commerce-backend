package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrderID        uuid.UUID `gorm:"type:uuid;index;not null" json:"order_id"`
	Provider       string    `gorm:"type:varchar(50);not null" json:"provider"` // e.g., Stripe, Paystack
	Method         string    `gorm:"type:varchar(50);not null" json:"method"`   // card, bank_transfer, wallet
	Amount         float64   `gorm:"type:numeric(10,2);not null" json:"amount"`
	Status         string    `gorm:"type:varchar(20);not null;default:'initiated'" json:"status"` // initiated | successful | failed
	TransactionRef string    `gorm:"type:text;uniqueIndex;not null" json:"transaction_ref"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	DelFlg         bool      `gorm:"default:false" json:"del_flg"`

	// Relations
	Order   Order    `gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order,omitempty"`
	Refunds []Refund `gorm:"foreignKey:PaymentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"refunds,omitempty"`
}

func (Payment) TableName() string {
	return "public.payments"
}
