package impl

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/repository"
)

type dbs struct {
	DB *gorm.DB
}

// NewUsersRepository will create an object that represent the User.Repository implementations
func NewUsersRepository(dbc *gorm.DB) repository.IUsers {
	return &dbs{
		DB: dbc,
	}
}

func (db *dbs) CreateUser(user *models.User) error {
	err := db.DB.Create(&user).Error
	return err
}

func (db *dbs) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := db.DB.Where("email = ?", email).Find(&user).Error
	return &user, err
}

func (db *dbs) GetAllUsers() ([]models.User, error) {
	var all_users []models.User
	err := db.DB.Find(&all_users).Error
	fmt.Println(all_users)
	return all_users, err
}

func (db *dbs) GetAUsers(id int) ([]models.User, error) {
	var user []models.User
	err := db.DB.Where("id = ?", id).Find(&user).Error
	fmt.Println("user", user)
	return user, err
}