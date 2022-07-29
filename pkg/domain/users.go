package domain

import "github.com/labstack/echo/v4"

type IUsers interface {
	Registration(c echo.Context) error
	GetAllUsers(c echo.Context) error
	GetAUsers(c echo.Context) error
	UpdateUser(c echo.Context) error 
	DeleteUser(c echo.Context) error
}