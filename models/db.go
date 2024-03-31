package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	var err error
	dsn := "enter your postgres dsn"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err.Error())
	}
	fmt.Println("Connected to database")

	db.AutoMigrate(&User{}, &GeoData{}) 
}
