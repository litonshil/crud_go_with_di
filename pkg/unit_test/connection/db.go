package connection

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/unit_test/testconfig"
)

var (
	db *gorm.DB
)

func Connect() *gorm.DB {

	d, err := gorm.Open("postgres", testconfig.GetConfig().POSTGRES_URL)

	if err != nil {
		fmt.Println("error connecting to db")
		panic(err)
	}
	db = d
	Migration()
	return db
}

func Migration() {
	db.AutoMigrate(&models.User{})

}

// func GetDB() *gorm.DB {
// 	return db
// }
