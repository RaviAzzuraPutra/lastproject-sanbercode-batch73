package models

import "time"

type TrxLogType string

const (
	TrxIn  TrxLogType = "in"
	TrxOut TrxLogType = "out"
)

type Trx_Log struct {
	ID        *string    `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	BarangID  *string    `json:"barang_id"`
	GudangID  *string    `json:"gudang_id"`
	Qty       *int       `json:"qty"`
	Type      TrxLogType `gorm:"type:trx_type" json:"type"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`

	Barang Barang `gorm:"foreignKey:BarangID;references:ID;constraint:OnDelete:CASCADE"`
	Gudang Gudang `gorm:"foreignKey:GudangID;references:ID;constraint:OnDelete:CASCADE"`
}

func (Trx_Log) TableName() string {
	return "trx_log"
}
