package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() {
	// PostgreSQL bağlantı bilgileri
	dsn := "user=postgres password=adem dbname=godatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %s", err)
	}

	fmt.Println("Veritabanına başarıyla bağlanıldı.")
}

// Database bağlantısını geri döndüren bir fonksiyon
func GetDB() *gorm.DB {
	return db
}
