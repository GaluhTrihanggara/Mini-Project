package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nis          int    `json:"nis"`
	Username     string `json:"username" form:"username"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Name         string `json:"name" form:"name"`
	IdKelas      string `json:"id_kelas" form:"id_kelas"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin"`
	TahunAjaran  string `json:"tahun_ajaran" form:"tahun_ajaran"`
}

type Token struct {
	Token string `json:"token" form:"token"`
}
