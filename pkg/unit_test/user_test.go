package unit_test

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	// "github.com/litonshil/crud_go_echo/config"
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/repository"
	"github.com/stretchr/testify/assert"
)

var u = &models.User{
	Id: 6,
}

var (
	db *gorm.DB
)

var userJson = models.User{
	Id:       6,
	Email:    "lalala@gmail.com",
	Name:     "Kaium",
	Address:  "Chittagong",
	Type:     "normal-user",
	Password: "1234",
}

func GetDB() *gorm.DB {
	// dns := fmt.Sprintf("%s/%s?charset=utf8mb4&parseTime=True&loc=Local", config.GetConfig().SqlUri, config.GetConfig().SqlDb)
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/user_crud?charset=utf8mb4&parseTime=True&loc=Local")
	// d, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Println("error connecting to db")
		panic(err)
	}
	db = d
	return db
}

func TestGetUserByID(t *testing.T) {
	db := GetDB()
	repo := repository.NewUsersRepository(db)

	user, err := repo.GetUserById(u.Id)
	assert.NotNil(t, user)
	assert.NoError(t, err)
	assert.Equal(t, userJson, user)
}
