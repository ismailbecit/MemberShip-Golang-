package config

import (
	"app/api/modal"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Conn() *gorm.DB {
	dsn := "root:sifre@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Veri Tabanına Bağlanılamadı")
	}
	db.AutoMigrate(&modal.Users{})
	return db

}
