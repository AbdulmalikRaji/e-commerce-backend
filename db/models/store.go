package models

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null;unique" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	OwnerID     uuid.UUID `gorm:"type:uuid;index;not null" json:"owner_id"` // FK to User
	Settings    string    `gorm:"type:jsonb" json:"settings"`
	Image       *string   `gorm:"type:text" json:"image,omitempty"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg      bool      `gorm:"default:false" json:"del_flg"`

	Owner    User        `gorm:"foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"owner,omitempty"`
	Products []Product   `gorm:"foreignKey:StoreID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products,omitempty"`
	Users    []StoreUser `gorm:"foreignKey:StoreID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"users,omitempty"`
}

func (Store) TableName() string {
	return "ecom.stores"
}
