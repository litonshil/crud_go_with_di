package impl

import (
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/repository"
	"github.com/litonshil/crud_go_echo/pkg/svc"
)

type users struct {
	urepo repository.IUsers
}

func NewUsersService(urepo repository.IUsers) svc.IUsers {
	return &users{
		urepo: urepo,
	}
}

func (u *users) CreateUser(user *models.User) (error){
	saveErr := u.urepo.CreateUser(user)
	if saveErr != nil {
		return saveErr
	}
	return saveErr
}

func (u *users) GetUserByEmail(email string) (*models.User, error) {
	res,err := u.urepo.GetUserByEmail(email)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (u *users) GetAllUsers() ([]models.User, error) {
	res,err := u.urepo.GetAllUsers()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (u *users) GetAUsers(id int) ([]models.User, error) {
	res,err := u.urepo.GetAUsers(id)
	if err != nil {
		return res, err
	}
	return res, nil
}
