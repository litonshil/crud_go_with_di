package controllers

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/domain"
	"github.com/litonshil/crud_go_echo/pkg/types"
)

type UserRepo struct {
	uSvc domain.IUsersSvc
}

func NewUserController(uSvc domain.IUsersSvc) *UserRepo {
	userc := &UserRepo{
		uSvc: uSvc,
	}
	return userc
}

// GetUsers fetch all user
func (ur *UserRepo) GetUsers(c echo.Context) error {

	res, err := ur.uSvc.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// GetUser fetch an specific user based on id
func (ur *UserRepo) GetUser(c echo.Context) error {

	id := c.Param("id")
	user_id, _ := strconv.Atoi(id)
	res, err := ur.uSvc.GetUser(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// // UpdateUser update an user
func (ur *UserRepo) UpdateUser(c echo.Context) error {

	var user = new(types.UserRegisterReq)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")

	user_id, _ := strconv.Atoi(id)

	res, err := ur.uSvc.UpdateUser(user_id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// DeleteUser delete an user
func (ur *UserRepo) DeleteUser(c echo.Context) error {

	id := c.Param("id")
	user_id, _ := strconv.Atoi(id)
	err_delete := ur.uSvc.DeleteUser(user_id)
	if err_delete != nil {
		return c.JSON(http.StatusInternalServerError, err_delete.Error())
	}
	fmt.Println(err_delete)
	return c.JSON(http.StatusOK, "user deleted successfully")
}
