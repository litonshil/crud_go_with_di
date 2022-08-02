package svc

import (
	"fmt"

	"github.com/litonshil/crud_go_echo/pkg/domain"
	"github.com/litonshil/crud_go_echo/pkg/models"
)

type users struct {
	urepo domain.IUsersRepo
}

func NewUsersService(urepo domain.IUsersRepo) domain.IUsersSvc {
	return &users{
		urepo: urepo,
	}
}

func (u *users) CreateUser(user *models.User) error {
	saveErr := u.urepo.CreateUser(user)
	if saveErr != nil {
		return saveErr
	}
	return saveErr
}

func (u *users) GetUserByEmail(email string) (*models.User, error) {
	res, err := u.urepo.GetUserByEmail(email)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (u *users) GetUsers() ([]models.User, error) {
	res, err := u.urepo.GetUsers()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (u *users) GetUser(id int) (models.User, error) {
	res, err := u.urepo.GetUserById(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (u *users) UpdateUser(id int, user *models.User) (*models.User, error) {

	existed_user, existUserErr := u.urepo.GetUserById(id)
	if existUserErr != nil {
		fmt.Println("update error")
		return user, existUserErr
	}
	fmt.Println(existed_user)

	res, err := u.urepo.UpdateUser(id, user, existed_user)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (u *users) DeleteUser(id int) error {
	_, existUserErr := u.urepo.GetUserById(id)
	if existUserErr != nil {
		fmt.Println("passed")
		return existUserErr
	}

	deleteErr := u.urepo.DeleteUser(id)
	if deleteErr != nil {
		return deleteErr
	}
	return deleteErr
}
