package models

import (
	"time"
)

type Category struct {
	ID          *string `json:"ID" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Created_at  time.Time
	Updated_at  time.Time
}
