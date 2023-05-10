package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	BillId int    `json:"bill_id"`
	UserId int    `json:"user_id"`
	Amount int    `json:"amount"`
	Status string `json:"status"`
}
