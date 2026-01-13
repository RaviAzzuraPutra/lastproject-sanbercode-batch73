package models

import (
	"time"
)

type User struct {
	ID        *string   `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      *string   `json:"name"`
	Email     *string   `json:"email"`
	Password  *string   `json:"password"`
	No_Telp   *string   `json:"no_telp"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Toko *Toko `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "user"
}
