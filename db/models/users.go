package models

import "time"

type User struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Auth0ID   string    `gorm:"unique;not null" json:"auth0_id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Role      string    `gorm:"type:varchar(20);not null" json:"role"` // buyer, seller, admin
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now()" json:"updated_at"`

	// Relations
	Products  []Product `gorm:"foreignKey:SellerID"`
	Orders    []Order   `gorm:"foreignKey:BuyerID"`
	Addresses []Address `gorm:"foreignKey:UserID"`
	Reviews   []Review  `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "public.users"
}
