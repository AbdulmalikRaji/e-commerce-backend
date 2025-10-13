package models

import (
	"time"

	"github.com/google/uuid"
)

// Analytics
type ProductView struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProductID uuid.UUID  `gorm:"type:uuid;index;not null" json:"product_id"`
	UserID    *uuid.UUID `gorm:"type:uuid;index" json:"user_id"`                     // nullable for guests
	SessionID *string    `gorm:"type:varchar(64);index" json:"session_id,omitempty"` // for guest tracking
	ViewedAt  time.Time  `gorm:"autoCreateTime;index" json:"viewed_at"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
	User    *User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}

func (ProductView) TableName() string {
	return "public.product_views"
}

type AddToCartEvent struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProductID uuid.UUID  `gorm:"type:uuid;index;not null" json:"product_id"`
	UserID    *uuid.UUID `gorm:"type:uuid;index" json:"user_id"`                     // nullable for guests
	SessionID *string    `gorm:"type:varchar(64);index" json:"session_id,omitempty"` // for guest tracking
	Quantity  int        `gorm:"not null" json:"quantity"`
	AddedAt   time.Time  `gorm:"autoCreateTime;index" json:"added_at"`
	CartID    *uuid.UUID `gorm:"type:uuid;index" json:"cart_id"` // nullable for guests

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
	User    *User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	Cart    *Cart   `gorm:"foreignKey:CartID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"cart,omitempty"`
}

func (AddToCartEvent) TableName() string {
	return "public.add_to_cart_events"
}

type SalesStat struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProductID uuid.UUID  `gorm:"type:uuid;index;not null" json:"product_id"`
	UserID    *uuid.UUID `gorm:"type:uuid;index" json:"user_id,omitempty"`           // optional, for user-specific stats
	SessionID *string    `gorm:"type:varchar(64);index" json:"session_id,omitempty"` // for guest tracking
	TotalSold int        `gorm:"default:0" json:"total_sold"`
	Revenue   float64    `gorm:"type:numeric(10,2);default:0" json:"revenue"`
	LastSold  *time.Time `json:"last_sold"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
}

func (SalesStat) TableName() string {
	return "public.sales_stats"
}

type AbandonedCart struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CartID        uuid.UUID  `gorm:"type:uuid;index;not null" json:"cart_id"`
	UserID        *uuid.UUID `gorm:"type:uuid;index" json:"user_id"`
	SessionID     *string    `gorm:"type:varchar(64);index" json:"session_id,omitempty"` // for guest tracking
	LastUpdatedAt time.Time  `gorm:"autoUpdateTime" json:"last_updated_at"`
	AbandonedAt   time.Time  `gorm:"index" json:"abandoned_at"`

	// Relations
	Cart Cart  `gorm:"foreignKey:CartID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cart,omitempty"`
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}

func (AbandonedCart) TableName() string {
	return "public.abandoned_carts"
}
