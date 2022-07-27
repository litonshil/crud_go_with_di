package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	consts "github.com/litonshil/crud_go_echo/pkg/const"

	"github.com/litonshil/crud_go_echo/pkg/models"
	// "github.com/litonshil/crud_go_echo/pkg/domain"
	"github.com/litonshil/crud_go_echo/pkg/svc"
	"github.com/litonshil/crud_go_echo/pkg/token"
	"github.com/litonshil/crud_go_echo/pkg/types"
	"github.com/litonshil/crud_go_echo/pkg/utils"
)

var validate = validator.New()

type userRepo struct {
	uSvc svc.IUsers
}

func NewUserController(e *echo.Echo, uSvc svc.IUsers) {
	userc := &userRepo{
		uSvc: uSvc,
	}
	sub := e.Group("/user")
	sub.POST("/registration", userc.Registration)
	sub.POST("/login", userc.Login)
	sub.GET("/users", userc.GetAllUsers)
	sub.GET("/:id", userc.GetAUsers)
	// sub.PUT("/:id", userc.UpdateUser)
	// sub.DELETE("/:id", userc.DeleteUser)
}

// Registration create a user
func (ur *userRepo) Registration(c echo.Context) error {
	var user = new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	if validationerr := validate.Struct(user); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	// auth_token := c.Request().Header.Get("Authorization")
	// split_token := strings.Split(auth_token, "Bearer ")
	// fmt.Println(split_token)
	// claims, err := utils.DecodeToken(split_token[1])
	// fmt.Println("extracted token\n", claims)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	// }

	if err := ur.uSvc.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Send username and password via email
	if err := utils.SendEmail(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "user created successfullys")
}

// Login login user
func (ur *userRepo) Login(c echo.Context) error {
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

	model_user, err := ur.uSvc.GetUserByEmail(user.Email)
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
func (ur *userRepo) GetAllUsers(c echo.Context) error {

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	_, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	res, err := ur.uSvc.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// GetAUsers fetch an specific user based on id
func (ur *userRepo) GetAUsers(c echo.Context) error {

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	_, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	id := c.Param("id")
	user_id, _ := strconv.Atoi(id)
	res, err := ur.uSvc.GetAUsers(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// // // UpdateUser update an user
// func (ur *userRepo) UpdateUser(c echo.Context) error {

// 	auth_token := c.Request().Header.Get("Authorization")
// 	split_token := strings.Split(auth_token, "Bearer ")
// 	_, err := utils.DecodeToken(split_token[1])
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
// 	}

// 	var user = new(models.User)
// 	var old_user = new(models.User)

// 	if err := c.Bind(user); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	id := c.Param("id")

// 	user_id, _ := strconv.Atoi(id)

// 	res, err := ur.repo.UpdateUser(user_id, user, old_user)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, res)
// }

// // // DeleteUser delete an user
// func (ur *userRepo) DeleteUser(c echo.Context) error {

// 	auth_token := c.Request().Header.Get("Authorization")
// 	split_token := strings.Split(auth_token, "Bearer ")
// 	_, err := utils.DecodeToken(split_token[1])
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
// 	}

// 	id := c.Param("id")
// 	user_id, _ := strconv.Atoi(id)
// 	err_delete := ur.repo.DeleteUser(user_id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err_delete.Error())
// 	}
// 	return c.JSON(http.StatusOK, "user deleted successfully")
// }
