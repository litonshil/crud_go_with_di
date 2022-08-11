package connection

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
	"github.com/litonshil/crud_go_echo/config"
	"github.com/litonshil/crud_go_echo/pkg/models"
)

var (
	db *gorm.DB
)

func Connect() {
	var dns string
	var d *gorm.DB
	var err error
	if config.GetConfig().Mode == "Test" {
		dns = config.GetConfig().POSTGRES_URL
		d, err = gorm.Open("postgres", dns)

	} else {
		dns = fmt.Sprintf("%s/%s?charset=utf8mb4&parseTime=True&loc=Local", config.GetConfig().SqlUri, config.GetConfig().SqlDb)
		d, err = gorm.Open("mysql", dns)

	}
	// d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/user_crud?charset=utf8mb4&parseTime=True&loc=Local")
	// d, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Println("error connecting to db")
		panic(err)
	}
	db = d
	Migration()
}

func Migration() {
	db.AutoMigrate(&models.User{})

}

func GetDB() *gorm.DB {
	return db
}
