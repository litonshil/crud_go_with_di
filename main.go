package main

import (
	"github.com/labstack/echo/v4"
	container "github.com/litonshil/crud_go_echo/pkg"
	"github.com/litonshil/crud_go_echo/pkg/connection"
)

func main() {
	connection.Connect()
	connection.ConnectRedis()
	e := echo.New()
	container.Init(e)
	e.Start(":8080")

}
