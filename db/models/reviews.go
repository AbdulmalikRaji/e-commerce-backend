package models

import "time"

type Review struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProductID string    `gorm:"type:uuid;not null;index" json:"product_id"`
	UserID    string    `gorm:"type:uuid;not null;index" json:"user_id"`
	Score     int       `gorm:"type:int;not null;check:score >= 1 AND score <= 5" json:"score"`
	Comment   string    `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now()" json:"updated_at"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User    User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Review) TableName() string {
	return "public.reviews"
}
