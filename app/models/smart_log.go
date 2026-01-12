package models

import (
	"time"
)

type Smart_Log struct {
	ID                     *string `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Barang_id              *string `json:"barang_id"`
	Gudang_id              *string `json:"gudang_id"`
	EOQ_calculation_result *int    `json:"eoq_calculation_result"`
	AI_Insight             *string `json:"ai_insight"`
	Created_at             time.Time
	Updated_at             time.Time

	Barang Barang `gorm:"foreignKey:Barang_id;references:ID;constraint:OnDelete:CASCADE"`
	Gudang Gudang `gorm:"foreignKey:Gudang_id;references:ID"`
}
