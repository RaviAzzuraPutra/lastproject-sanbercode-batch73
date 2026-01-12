package models

import (
	"time"
)

type Toko struct {
	ID         *string `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	User_id    *string `json:"user_id"`
	Name       *string `json:"name"`
	Address    *string `json:"address"`
	Created_at time.Time
	Updated_at time.Time

	User *User `gorm:"foreignKey:User_id;references:ID;constraint:OnDelete:CASCADE"`
}
