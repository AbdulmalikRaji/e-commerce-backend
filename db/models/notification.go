package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	Type      string    `gorm:"type:varchar(50)" json:"type"` // order_update | payment_update | system_message
	Message   string    `gorm:"type:text" json:"message"`
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	CreatedBy *uuid.UUID `gorm:"type:uuid;index" json:"created_by,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid;index" json:"updated_by,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relations
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}

func (Notification) TableName() string {
	return "ecom.notifications"
}
