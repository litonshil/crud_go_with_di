package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id       int    `gorm:"primary_key;AUTO_INCREMENT"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Type     string `json:"type"`
	Password string `json:"password"`
}
