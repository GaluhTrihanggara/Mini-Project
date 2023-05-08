package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	ID       int    `json:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
}
