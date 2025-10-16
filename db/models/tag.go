package models

import "github.com/google/uuid"

type Tag struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	CreatedAt int64      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64      `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy *uuid.UUID `gorm:"type:uuid;index" json:"created_by,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid;index" json:"updated_by,omitempty"`
	DelFlg    bool       `gorm:"default:false" json:"del_flg"`

	Products []Product `gorm:"many2many:product_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"products,omitempty"`
}

func (Tag) TableName() string {
	return "ecom.tags"
}
