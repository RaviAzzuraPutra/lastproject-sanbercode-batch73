package models

import (
	"time"
)

type Smart_Log struct {
	ID                     *string   `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Barang_id              *string   `json:"barang_id"`
	Gudang_id              *string   `json:"gudang_id"`
	Period_month           *int      `json:"period_month"`
	Period_year            *int      `json:"period_year"`
	EOQ_calculation_result *int      `json:"eoq_calculation_result"`
	ROP_value              *int      `json:"rop_value"`
	AI_Insight             *string   `json:"ai_insight"`
	CreatedAt              time.Time `gorm:"column:created_at"`
	UpdatedAt              time.Time `gorm:"column:updated_at"`

	Barang Barang `gorm:"foreignKey:Barang_id;references:ID;constraint:OnDelete:CASCADE"`
	Gudang Gudang `gorm:"foreignKey:Gudang_id;references:ID"`
}

func (Smart_Log) TableName() string {
	return "smart_log"
}
