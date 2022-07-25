package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/litonshil/crud_go_echo/config"
	"github.com/litonshil/crud_go_echo/pkg/models"
)

var (
	db *gorm.DB
)

func Connect() {

	dns := fmt.Sprintf("%s/%s?charset=utf8mb4&parseTime=True&loc=Local", config.GetConfig().SqlUri, config.GetConfig().SqlDb)
	// d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/user_crud?charset=utf8mb4&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Println("error connecting to db")
		panic(err)
	}
	db = d
}

func Migration() {
	db.AutoMigrate(&models.User{})

}

func GetDB() *gorm.DB {
	if db == nil {
		Connect()
		Migration()
	}

	return db
}
