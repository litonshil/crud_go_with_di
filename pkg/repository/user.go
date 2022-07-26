package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/litonshil/crud_go_echo/pkg/models"
)

type IUsers interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetAUsers(id int) ([]models.User, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
	GetUserByEmail(email string) (*models.User, error)
}

type dbs struct {
	DB *gorm.DB
}

func NewDb(db *gorm.DB) *dbs {
	return &dbs{
		DB: db,
	}
}

func (db *dbs) CreateUser(user *models.User) error {
	err := db.DB.Create(&user).Error
	return err
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

func (db *dbs) UpdateUser(id int, user *models.User) (*models.User, error) {

	err := db.DB.Model(&user).Where("id = ?", id).Update(&user).Error
	fmt.Println("user", user)
	return user, err
}

func (db *dbs) DeleteUser(id int) error {
	var user []models.User
	err := db.DB.Where("id = ?", id).Delete(&user).Error
	return err
}

func (db *dbs) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := db.DB.Where("email = ?", email).Find(&user).Error
	return &user, err
}
