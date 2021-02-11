package controllers

import (
	"database/sql"
	"fmt"
	"github.com/revel/revel"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"os"
)

var DB *sql.DB

func InitDB() {
	load := godotenv.Load(".env")
	
	if load != nil {
		panic(load)
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PWD := os.Getenv("DB_PWD")
	DB_NAME := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PWD, DB_NAME)

	var err error
	DB, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	defer DB.Close()

	err = DB.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfull connected")
}

func init() {
	revel.OnAppStart(InitDB)
}