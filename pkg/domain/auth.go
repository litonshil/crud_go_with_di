package domain

import "github.com/labstack/echo/v4"

type IAuthController interface {
	Login(c echo.Context) error
}