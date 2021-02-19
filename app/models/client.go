package models

import (
	// "fmt"
	"gorm.io/gorm"
	"time"
)

type Client struct {
	gorm.Model
	First_name string `sql:"size:255" json:"first_name"`
	Last_name string `sql:"size:255" json:"last_name"`
	Ci string `sql:"size:20" json:"ci"`
	Birthday time.Time `sql:"timestamptz" json:"birthday"`
	Sex string `sql:size:1 json:"sex"`
	Nationality string `sql:size:140 json:"nationality"`
	Des_type string `sql:size:140 json: "des_type"`
	Code1 string `sql:size:20 json:"code1"`
	Code2 string `sql:size:20 json:"code2"`
	Code3 string `sql:size:20 json:"code3"`
	Direction string `sql:size:20 json:"direction"`
	Phone string `sql:size:10 json:"phone"`
}

// const (
// 	DATA_FORMAT = "Jan _2, 2006"
// 	SQL_DATE_FORMAT = "2006-01-02"
// )

// func (c Client) formtDate() string {
// 	return fmt.Sprintf("%s", c.Birthday.Format(SQL_DATE_FORMAT))
// }

func (c Client) BirthdayDateStr() string {
	return c.Birthday.Format("2006-01-02")
}