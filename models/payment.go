package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	IdPayment   int    `gorm:"primaryKey" json:"id_payment"`
	UserId      int    `json:"user_id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Status      string `json:"status"`
}
