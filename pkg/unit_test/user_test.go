package unit_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/repository"
	"github.com/litonshil/crud_go_echo/pkg/unit_test/connection"
	"github.com/stretchr/testify/assert"
)

var u = models.User{
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
	Id:       6,
	Email:    "abcdef@gmail.com",
	Name:     "abcde",
	Address:  "abcde",
	Type:     "admin",
	Password: "abcde",
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
	// assert.Equal(t, userJson, user)
}

func TestCreateUser(t *testing.T) {
	d := connection.Connect()
	db = d
	repo := repository.NewUsersRepository(db)
	err := repo.CreateUser(userReg)
	assert.NoError(t, err)
	// db.Delete(&models.User{})
}

var updatedUser = &models.User{
	Id:       3,
	Email:    "abcd@gmail.com",
	Name:     "bbbbb",
	Address:  "bbbbb",
	Type:     "admin",
	Password: "bbbbb",
}

var existed_user = models.User{
	Id:       3,
	Email:    "bbbbbcd@gmail.com",
	Name:     "bbbbb",
	Address:  "bbbbb",
	Type:     "admin",
	Password: "bbbbb",
}

func TestUpdateUser(t *testing.T) {
	d := connection.Connect()
	db = d
	repo := repository.NewUsersRepository(db)
	users, err := repo.UpdateUser(3, updatedUser, existed_user)
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestDeleteUser(t *testing.T) {
	d := connection.Connect()
	db = d
	repo := repository.NewUsersRepository(db)
	err := repo.DeleteUser(4)
	assert.NoError(t, err)
}

func TestGetUsers(t *testing.T) {
	d := connection.Connect()
	db = d
	repo := repository.NewUsersRepository(db)
	users, err := repo.GetUsers()
	assert.NoError(t, err)
	assert.NotNil(t, users)
}
