package models

import (
	"time"
)

type Gudang struct {
	ID         *string `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Toko_id    *string `json:"toko_id"`
	Name       *string `json:"name"`
	Address    *string `json:"address"`
	Created_at time.Time
	Updated_at time.Time

	Toko Toko `gorm:"foreignKey:Toko_id;references:ID;constraint:OnDelete:CASCADE"`
}
