package controllers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"github.com/revel/revel"
	"os"
	models "docusys/app/models"
)


var DB *gorm.DB

func InitDB () {
	load := godotenv.Load(".env")

	if load != nil {
		panic(load)
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PWD := os.Getenv("DB_PWD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PWD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.DB()
	db.AutoMigrate(&models.Client{})
  DB = db
}

func init () {
	revel.OnAppStart(InitDB)
}