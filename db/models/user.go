package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	AuthID      uuid.UUID `gorm:"type:text;uniqueIndex;not null" json:"auth_id"`
	Email       string    `gorm:"type:text;uniqueIndex;not null" json:"email"`
	FirstName   string    `gorm:"type:varchar(50)" json:"first_name"`
	LastName    string    `gorm:"type:varchar(50)" json:"last_name"`
	PhoneNumber string    `gorm:"type:varchar(20)" json:"phone_number"`
	Role        string    `gorm:"type:varchar(20);not null;default:'buyer'" json:"role"` // buyer | seller | admin
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg      bool      `gorm:"default:false" json:"del_flg"`

	// Relations (grouped)
	Addresses     []Address      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"addresses,omitempty"`
	Orders        []Order        `gorm:"foreignKey:BuyerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders,omitempty"`
	Reviews       []Review       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reviews,omitempty"`
	Notifications []Notification `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"notifications,omitempty"`
	Cart          Cart           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cart,omitempty"` // ONE cart per user
}

func (User) TableName() string {
	return "ecom.users"
}
