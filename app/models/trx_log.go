package models

import "time"

type TrxLogType string

const (
	TrxIn  TrxLogType = "in"
	TrxOut TrxLogType = "out"
)

type Trx_Log struct {
	ID         *string    `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Barang_id  *string    `json:"barang_id"`
	Qty        *string    `json:"qty"`
	Type       TrxLogType `sql:"type:trx_type" json:"type"`
	Created_at time.Time
	Updated_at time.Time

	Barang Barang `gorm:"foreignKey:Barang_id;references:ID;constraint:OnDelete:CASCADE"`
}
