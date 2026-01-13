package models

import (
	"time"
)

type Barang struct {
	ID             *string   `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Gudang_id      *string   `json:"gudang_id"`
	Category_id    *string   `json:"category_id"`
	Name           *string   `json:"name"`
	Sku            *string   `json:"sku"`
	Image_url      *string   `json:"image_url"`
	Stock          *int      `json:"stock"`
	Safety_stock   *int      `json:"safety_stock"`
	Need_restock   *bool     `json:"need_restock"`
	Lead_time_days *int      `json:"lead_time_days"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`

	Gudang   Gudang   `gorm:"foreignKey:Gudang_id;references:ID;constraint:OnDelete:CASCADE"`
	Category Category `gorm:"foreignKey:Category_id;references:ID"`
}

func (Barang) TableName() string {
	return "barang"
}
