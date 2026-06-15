package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"order-service/models" // Sesuaikan nama modul
)

func InitDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/tugasbesariae"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal konek ke database")
	}

	// Auto Migration: Membuat tabel otomatis di database
	db.AutoMigrate(&models.Order{}, &models.OrderItem{})
	return db
}