package config

import (
	"fmt"
	"log"

	"FinalTaskEvermoss/internal/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:mieAYAM123@tcp(127.0.0.1:3306)/evermos_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal koneksi database: ", err)
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Toko{},
		&entity.Address{},
		&entity.Product{},
		&entity.Transaction{},
	)
	if err != nil {
		log.Fatal("❌ Gagal migrate database: ", err)
	}

	DB = db
	fmt.Println("✅ Koneksi & Migrasi Database Berhasil!")
}
