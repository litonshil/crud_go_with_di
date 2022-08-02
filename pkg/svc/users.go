package svc

import (
	"fmt"

	"github.com/litonshil/crud_go_echo/pkg/domain"
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/types"
	"github.com/litonshil/crud_go_echo/pkg/utils"
)

type users struct {
	urepo domain.IUsersRepo
}

func NewUsersService(urepo domain.IUsersRepo) domain.IUsersSvc {
	return &users{
		urepo: urepo,
	}
}

// func (u *users) CreateUser(user *types.UserRegisterReq) error {
// 	var model_user = new(models.User)
// 	respErr := utils.StructToStruct(user, &model_user)
// 	if respErr != nil {
// 		return respErr
// 	}

// 	saveErr := u.urepo.CreateUser(model_user)
// 	if saveErr != nil {
// 		return saveErr
// 	}
// 	return saveErr
// }

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

func (u *users) UpdateUser(id int, user *types.UserRegisterReq) (*models.User, error) {
	var model_user = new(models.User)
	respErr := utils.StructToStruct(user, &model_user)
	if respErr != nil {
		return model_user, respErr
	}

	existed_user, existUserErr := u.urepo.GetUserById(id)
	if existUserErr != nil {
		fmt.Println("update error")
		return model_user, existUserErr
	}
	fmt.Println(existed_user)

	res, err := u.urepo.UpdateUser(id, model_user, existed_user)
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
