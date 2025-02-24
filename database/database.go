package database

import (
	"github.com/tfajarama/go-struct-pos/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "user:password@tcp(localhost:3306)/yourdb?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	err = database.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Could not migrate `Product` database: ", err)
	}
	err = database.AutoMigrate(&models.Customer{})
	if err != nil {
		log.Fatal("Could not migrate `Customer` database: ", err)
	}

	DB = database
}
