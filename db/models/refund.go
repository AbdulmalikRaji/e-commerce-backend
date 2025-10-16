package models

import (
	"time"

	"github.com/google/uuid"
)

type Refund struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	PaymentID   uuid.UUID  `gorm:"type:uuid;index;not null" json:"payment_id"`
	Amount      float64    `gorm:"type:numeric(10,2);not null" json:"amount"`
	Reason      string     `gorm:"type:text" json:"reason"`
	Status      string     `gorm:"type:varchar(20);not null;default:'pending'" json:"status"` // pending | approved | rejected | processed
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	ProcessedAt *time.Time `json:"processed_at"`

	// Relations
	Payment Payment `gorm:"foreignKey:PaymentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"payment,omitempty"`
}

func (Refund) TableName() string {
	return "ecom.refunds"
}
