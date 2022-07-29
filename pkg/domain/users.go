package domain

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/models"
)

type IUsersController interface {
	Registration(c echo.Context) error
	GetAllUsers(c echo.Context) error
	GetAUsers(c echo.Context) error
	UpdateUser(c echo.Context) error 
	DeleteUser(c echo.Context) error
}

type IUsersSvc interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetAUsers(id int) (models.User, error)
	UpdateUser(id int, user *models.User, old_user *models.User) (*models.User, error)
	DeleteUser(id int) error
}

type IUsersRepo interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetAUsers(id int) (models.User, error)
	UpdateUser(id int, user *models.User, old_user *models.User) (*models.User, error)
	DeleteUser(id int) error
}
