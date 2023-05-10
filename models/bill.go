package models

import (
	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	UserId      int    `json:"user_id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Status      string `json:"status"`
}
