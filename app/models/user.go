package models

import (
	"time"
)

type User struct {
	ID         *string `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Password   *string `json:"password"`
	No_Telp    *string `json:"no_telp"`
	Created_at time.Time
	Updated_at time.Time
}
