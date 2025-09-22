package models

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    *uuid.UUID `gorm:"type:uuid;index;uniqueIndex:idx_carts_userid" json:"user_id"`         // nullable for guest
	SessionID *string    `gorm:"type:varchar(255);uniqueIndex:idx_carts_sessionid" json:"session_id"` // guest session id
	IsActive  bool       `gorm:"default:true;index" json:"is_active"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg    bool       `gorm:"default:false" json:"del_flg"`

	// Relations
	User  *User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Items []CartItem `gorm:"foreignKey:CartID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items,omitempty"`
}

func (Cart) TableName() string {
	return "public.carts"
}

type CartItem struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CartID    uuid.UUID  `gorm:"type:uuid;index;not null" json:"cart_id"`
	ProductID uuid.UUID  `gorm:"type:uuid;index;not null" json:"product_id"`
	VariantID *uuid.UUID `gorm:"type:uuid;index" json:"variant_id"` // nullable if product has no variants
	Quantity  int        `gorm:"not null;default:1" json:"quantity"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg    bool       `gorm:"default:false" json:"del_flg"`

	// Relations
	Cart    Cart    `gorm:"foreignKey:CartID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cart,omitempty"`
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
}

func (CartItem) TableName() string {
	return "public.cart_items"
}
