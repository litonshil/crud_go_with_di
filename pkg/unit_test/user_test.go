package unit_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/repository"
	"github.com/litonshil/crud_go_echo/pkg/unit_test/connection"
	"github.com/stretchr/testify/assert"
)

var u = &models.User{
	Id: 1,
}

var (
	db *gorm.DB
)

var userJson = models.User{
	Id:       1,
	Email:    "aaaaa@gmail.com",
	Name:     "aaaaa",
	Address:  "aaaaa",
	Type:     "admin",
	Password: "aaaaa",
}

var userReg = &models.User{
	Id: 5,
	Email:    "bbbbb@gmail.com",
	Name:     "bbbbb",
	Address:  "bbbbb",
	Type:     "admin",
	Password: "bbbbb",
}

func GetDB() *gorm.DB {

	return db
}

func TestGetUserByID(t *testing.T) {
	d := connection.Connect()
	db = d
	repo := repository.NewUsersRepository(db)
	user, err := repo.GetUserById(u.Id)
	assert.NotNil(t, user)
	assert.NoError(t, err)
	assert.Equal(t, userJson, user)
}

func TestCreateUser(t *testing.T) {
	repo := repository.NewUsersRepository(db)
	err := repo.CreateUser(userReg)
	assert.NoError(t, err)
	db.Delete(&models.User{})

}
