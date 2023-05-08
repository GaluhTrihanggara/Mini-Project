package models

import (
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	ID  int `json:"id" gorm:"primaryKey"`
	Nis int `json:"nis"`

	UserID     int     `json:"user_id" gorm:"not null"`
	PaymentID  int     `json:"payment_id" gorm:"not null"`
	TotalPrice float64 `json:"total_price" gorm:"not null"`
}

func (h *History) TableName() string {
	return "history"
}
