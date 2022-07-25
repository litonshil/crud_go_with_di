package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id       	int			`gorm:"primary_key;AUTO_INCREMENT"`
	Email		string		`json:"email" validate:"required,email"`
	Name     	string      `json:"name" validate:"required,min=2,max=15"`
	Address		string		`json:"address"`
	Type		string		`json:"type" validate:"required"`	
	Password  string		`json:"password" validate:"required"`
}