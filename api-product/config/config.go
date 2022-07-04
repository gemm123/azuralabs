package config

import (
	"log"
	"test-azuralabs/api-product/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB(dsn string) {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cant connect database!")
	}

	log.Println("connected to database")
}

func MigrateDB() {
	DB.AutoMigrate(&models.Product{})
}

func CloseDB() {
	dbSQL, _ := DB.DB()
	dbSQL.Close()
}
