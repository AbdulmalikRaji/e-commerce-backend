package models

import (
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Code           string     `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"`
	DiscountPct    *float64   `gorm:"type:numeric(5,2)" json:"discount_percent"` // nullable: either percent or amount
	DiscountAmount *float64   `gorm:"type:numeric(10,2)" json:"discount_amount"`
	MaxUses        *int       `json:"max_uses"`
	UsedCount      int        `gorm:"default:0" json:"used_count"`
	ValidFrom      *time.Time `json:"valid_from"`
	ValidUntil     *time.Time `json:"valid_until"`
	CreatedBy *uuid.UUID `gorm:"type:uuid;index" json:"created_by,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid;index" json:"updated_by,omitempty"`
	CreatedAt      time.Time  `gorm:"autoCreateTime" json:"created_at"`
	DelFlg         bool       `gorm:"default:false" json:"del_flg"`

	// Relations (if you want coupon -> orders many2many, create join table order_coupons separately)
}

func (Coupon) TableName() string {
	return "ecom.coupons"
}
