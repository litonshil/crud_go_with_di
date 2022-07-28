package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/litonshil/crud_go_echo/pkg/models"
)


type dbs struct {
	DB *gorm.DB
}

func NewDb(db *gorm.DB) *dbs {
	return &dbs{
		DB: db,
	}
}

type IUsers interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetAUsers(id int) (models.User, error)
	UpdateUser(id int, user *models.User, old_user *models.User) (*models.User, error)
	DeleteUser(id int) error
}