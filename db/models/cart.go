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
	return "ecom.carts"
}
