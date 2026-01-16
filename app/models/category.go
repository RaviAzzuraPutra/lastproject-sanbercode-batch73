package models

import (
	"time"
)

type Category struct {
	ID          *string   `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	TokoID      *string   `json:"toko_id"`
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`

	Toko *Toko `gorm:"foreignKey:TokoID;references:ID;constraint:OnDelete:CASCADE"`
}

func (Category) TableName() string {
	return "category"
}
