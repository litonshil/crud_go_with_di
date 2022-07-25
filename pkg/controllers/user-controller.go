package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	consts "github.com/litonshil/crud_go_echo/pkg/const"
	"github.com/litonshil/crud_go_echo/pkg/database"
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/repository"
	"github.com/litonshil/crud_go_echo/pkg/token"
	"github.com/litonshil/crud_go_echo/pkg/types"
	"github.com/litonshil/crud_go_echo/pkg/utils"
)

var db = database.GetDB()
var validate = validator.New()

// Registration create a user
func Registration(c echo.Context) error {
	var user = new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	if validationerr := validate.Struct(user); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	fmt.Println(split_token)
	claims, err := utils.DecodeToken(split_token[1])
	fmt.Println("extracted token\n", claims)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	if err := repository.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Send username and password via email
	if err := utils.SendEmail(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "user created successfullys")
}

// Login login user
func Login(c echo.Context) error {
	var user = new(types.User)
	var model_user = new(models.User)
	var tokens = new(types.Token)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	if validationerr := validate.Struct(user); validationerr != nil {
		fmt.Println("error")
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	model_user, err := repository.GetUserByEmail(user.Email)
	if model_user.Email == "" || err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	token, refresh_token, err := token.GenerateUserTokens(model_user.Email, model_user.Id, model_user.Type)
	tokens.User_Token = token
	tokens.User_Refreshtoken = refresh_token

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tokens)
}

// GetAllUsers fetch all user
func GetAllUsers(c echo.Context) error {

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	_, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	res, err := repository.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// GetAUsers fetch an specific user based on id
func GetAUsers(c echo.Context) error {

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	_, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	id := c.Param("id")
	user_id, _ := strconv.Atoi(id)
	res, err := repository.GetAUsers(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// checkEmptyUserField set all empty field with old data when an user is update
func checkEmptyUserField(user *models.User, old_user *models.User) *models.User {
	if user.Name == "" {
		user.Name = old_user.Name
	}
	if user.Address == "" {
		user.Address = old_user.Address
	}
	if user.Email == "" {
		user.Email = old_user.Email
	}
	if user.Type == "" {
		user.Type = old_user.Type
	}
	if user.Password == "" {
		user.Password = old_user.Password
	}
	return user
}

// UpdateUser update an user
func UpdateUser(c echo.Context) error {

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	_, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	var user = new(models.User)
	var old_user = new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")

	user_id, _ := strconv.Atoi(id)
	old_err := db.Model(old_user).Where("id = ?", id).Find(&old_user).Error

	if old_err != nil {
		return c.JSON(http.StatusInternalServerError, old_err.Error())
	}

	user.Id = user_id

	checkedUser := checkEmptyUserField(user, old_user)

	res, err := repository.UpdateUser(user_id, checkedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// DeleteUser delete an user
func DeleteUser(c echo.Context) error {

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	_, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	id := c.Param("id")
	user_id, _ := strconv.Atoi(id)
	err_delete := repository.DeleteUser(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err_delete.Error())
	}
	return c.JSON(http.StatusOK, "user deleted successfully")
}
