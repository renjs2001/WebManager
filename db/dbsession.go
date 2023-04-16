package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

// var dsn string

func Init() {
	// dsn = "root:123456@tcp(127.0.0.1:3306)/finalProject?charset=utf8mb4&parseTime=True&loc=local"
	var err error
	db, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/finalProject?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Panicf("failed to open mysql , error: %v\n", err)
	}
}

func GetDBConn() *gorm.DB {
	return db
}
