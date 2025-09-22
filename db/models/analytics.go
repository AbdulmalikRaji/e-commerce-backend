package models

import (
	"time"

	"github.com/google/uuid"
)

//
// Analytics
//
type ProductView struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProductID uuid.UUID  `gorm:"type:uuid;index;not null" json:"product_id"`
	UserID    *uuid.UUID `gorm:"type:uuid;index" json:"user_id"` // nullable for guests
	ViewedAt  time.Time  `gorm:"autoCreateTime;index" json:"viewed_at"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
	User    *User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}

type SalesStat struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProductID uuid.UUID  `gorm:"type:uuid;index;not null" json:"product_id"`
	TotalSold int        `gorm:"default:0" json:"total_sold"`
	Revenue   float64    `gorm:"type:numeric(10,2);default:0" json:"revenue"`
	LastSold  *time.Time `json:"last_sold"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
}

type AbandonedCart struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CartID        uuid.UUID  `gorm:"type:uuid;index;not null" json:"cart_id"`
	UserID        *uuid.UUID `gorm:"type:uuid;index" json:"user_id"`
	LastUpdatedAt time.Time  `gorm:"autoUpdateTime" json:"last_updated_at"`
	AbandonedAt   time.Time  `gorm:"index" json:"abandoned_at"`

	// Relations
	Cart Cart  `gorm:"foreignKey:CartID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cart,omitempty"`
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}
