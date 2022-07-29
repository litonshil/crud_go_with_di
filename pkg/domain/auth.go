package domain

import "github.com/labstack/echo/v4"

type IAuth interface {
	Login(c echo.Context) error
}