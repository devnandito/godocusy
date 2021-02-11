package controllers

import (
	"database/sql"
	"github.com/revel/revel"
	_ "github.com/lib/pq"
	"fmt"
)

var DB *sql.DB

func InitDB() {
	connstring := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", "tech", "F3rnand@21", "docusys")

	var err error
	DB, err := sql.Open("postgres", connstring)
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