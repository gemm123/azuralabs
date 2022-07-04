package main

import (
	"fmt"
	"log"
	"os"
	"test-azuralabs/api-product/config"
	"test-azuralabs/api-product/controller"
	"test-azuralabs/api-product/repository"
	"test-azuralabs/api-product/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("cant get .env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName)
	config.InitDB(dsn)
	config.MigrateDB()
	defer config.CloseDB()

	r := gin.Default()
	repository := repository.NewRepository(config.DB)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	api := r.Group("/api")
	{
		api.GET("/product", controller.GetAllProduct)
		api.POST("/product", controller.PostProduct)
		api.GET("/product/:id", controller.GetProductByID)
	}

	r.Run()
}
