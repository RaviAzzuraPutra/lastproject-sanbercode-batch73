package models

import (
	"time"
)

type Barang struct {
	ID           *string `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Gudang_id    *string `json:"gudang_id"`
	Category_id  *string `json:"category_id"`
	Name         *string `json:"name"`
	Sku          *string `json:"sku"`
	Stock        *int    `json:"stock"`
	Need_restock *bool   `json:"need_restock"`
	Created_at   time.Time
	Updated_at   time.Time

	Gudang   Gudang   `gorm:"foreignKey:Gudang_id;references:ID;constraint:OnDelete:CASCADE"`
	Category Category `gorm:"foreignKey:Category_id;references:ID"`
}
