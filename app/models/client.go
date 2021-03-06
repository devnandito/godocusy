package models

import (
	// "fmt"
	"github.com/revel/revel"
	"gorm.io/gorm"
	"time"
)

type Client struct {
	gorm.Model
	First_name string `sql:"size:255" json:"first_name"`
	Last_name string `sql:"size:255" json:"last_name"`
	Ci string `sql:"size:20" json:"ci"`
	Birthday time.Time `sql:"date" json:"birthday`
	Sex string `sql:size:1 json:"sex"`
	Nationality string `sql:size:140 json:"nationality"`
	Des_type string `sql:size:140 json: "des_type"`
	Code1 string `sql:size:20 json:"code1"`
	Code2 string `sql:size:20 json:"code2"`
	Code3 string `sql:size:20 json:"code3"`
	Direction string `sql:size:20 json:"direction"`
	Phone string `sql:size:10 json:"phone"`
}

func (c Client) BirthdayDateStr() string {
	return c.Birthday.Format("2006-01-02")
}

func (client *Client) Validate(v *revel.Validation) {
	v.Required(client.First_name).Message("* El nombre es requerido")
	v.Required(client.Last_name).Message("* El apellido es requerido")
	v.Required(client.Ci).Message("* La cedula es requerido")
	v.Required(client.Birthday).Message("* La fecha de nacimiento es requerido")
	v.Required(client.Sex).Message("* El sexo es requerido")
	v.MaxSize(client.Sex, 1).Message("* Ingrese solo la inicial")
}

// const (
// 	DATA_FORMAT = "Jan _2, 2006"
// 	SQL_DATE_FORMAT = "2006-01-02"
// )

// func (c Client) formtDate() string {
// 	return fmt.Sprintf("%s", c.Birthday.Format(SQL_DATE_FORMAT))
// }