package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/helpers"
)

// Authenticate - check am user is logged in or not
func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authToken := c.Request().Header.Get("Authorization")
		splitToken := strings.Split(authToken, "Bearer ")

		if len(splitToken) != 2 {
			return c.String(http.StatusUnauthorized, "you need to login for access this page")
		}

		reqToken := splitToken[1]
		if reqToken == "" {
			return c.String(http.StatusInternalServerError, "failed to getting your token")
		}

		ok, err := helpers.VerifyToken(reqToken)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		if !ok {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return next(c)
	}
}
