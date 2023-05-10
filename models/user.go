package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nis          int    `json:"nis"`
	Name         string `json:"name" form:"name"`
	Username     string `json:"username" form:"username"`
	Password     string `json:"password" form:"password"`
	Email        string `json:"email" form:"email"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin"`
	TahunAjaran  string `json:"tahun_ajaran" form:"tahun_ajaran"`
}

type Token struct {
	Token string `json:"token" form:"token"`
}
