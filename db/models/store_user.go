package models

import (
	"time"

	"github.com/google/uuid"
)

// StoreUser represents a user granted explicit permissions on a store.
// The store has one owner (Store.OwnerID) but additional users can be
// granted fine-grained rights (add/update/delete products, manage orders, etc).
type StoreUser struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	StoreID uuid.UUID `gorm:"type:uuid;index;not null" json:"store_id"`
	UserID  uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`

	// Role is a higher-level role that implies default permissions.
	// Use values from the constants: RoleManager, RoleWorker
	Role string `gorm:"type:text;default:'worker'" json:"role"`

	// Explicit permission flags.
	CanAddProducts         bool `gorm:"default:false" json:"can_add_products"`
	CanUpdateProducts      bool `gorm:"default:false" json:"can_update_products"`
	CanDeleteProducts      bool `gorm:"default:false" json:"can_delete_products"`
	CanManageOrders        bool `gorm:"default:false" json:"can_manage_orders"`
	CanManageStoreSettings bool `gorm:"default:false" json:"can_manage_store_settings"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DelFlg    bool      `gorm:"default:false" json:"del_flg"`

	// Relations
	User  User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Store Store `gorm:"foreignKey:StoreID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"store,omitempty"`
}

func (StoreUser) TableName() string {
	return "ecom.store_users"
}

// Role constants for StoreUser.Role
const (
	RoleManager = "manager"
	RoleWorker  = "worker"
)
