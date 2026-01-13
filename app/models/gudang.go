package models

import (
	"time"
)

type Gudang struct {
	ID        *string   `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Toko_id   *string   `json:"toko_id"`
	Name      *string   `json:"name"`
	Address   *string   `json:"address"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Toko Toko `gorm:"foreignKey:Toko_id;references:ID;constraint:OnDelete:CASCADE"`
}

func (Gudang) TableName() string {
	return "gudang"
}
