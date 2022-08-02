package domain

import (
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/types"
)

type IUsersSvc interface {
	CreateUser(user *types.UserRegisterReq) error
	GetUserByEmail(email string) (*models.User, error)
	GetUsers() ([]models.User, error)
	GetUser(id int) (models.User, error)
	UpdateUser(id int, user *types.UserRegisterReq) (*models.User, error)
	DeleteUser(id int) error
}

type IUsersRepo interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUsers() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	UpdateUser(id int, user *models.User, old_user models.User) (*models.User, error)
	DeleteUser(id int) error
}
