package models

import (
	"time"
)

type Toko struct {
	ID        *string   `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    *string   `json:"user_id"`
	Name      *string   `json:"name"`
	Address   *string   `json:"address"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (Toko) TableName() string {
	return "toko"
}
