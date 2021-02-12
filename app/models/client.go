package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	First_name string `sql:"size:255" json:"first_name"`
	Last_name string `sql:"size:255" json:"last_name"`
	Ci string `sql:"size:255" json:"ci"`
}