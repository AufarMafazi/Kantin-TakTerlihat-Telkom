package main

import (
	"auth-service/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/tugasbesariae?parseTime=true"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	return db
}