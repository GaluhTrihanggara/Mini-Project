package models

import (
	"gorm.io/gorm"
)

type Histories struct {
	gorm.Model
	PaymentID   int    `json:"payment_id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func (h *Histories) TableName() string {
	return "histories"
}
