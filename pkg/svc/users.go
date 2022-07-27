package svc

import "github.com/litonshil/crud_go_echo/pkg/models"

type IUsers interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetAUsers(id int) ([]models.User, error)
	UpdateUser(id int, user *models.User, old_user *models.User) (*models.User, error)
}