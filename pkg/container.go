package continer

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/connection"
	"github.com/litonshil/crud_go_echo/pkg/controllers"
	"github.com/litonshil/crud_go_echo/pkg/routes"

	// "github.com/litonshil/crud_go_echo/pkg/repository"
	repoImpl "github.com/litonshil/crud_go_echo/pkg/repository"
	svcImpl "github.com/litonshil/crud_go_echo/pkg/svc"
)

func Init(e *echo.Echo) {
	db := connection.GetDB()
	client := connection.Redis()

	userRepo := repoImpl.NewUsersRepository(db, client)
	userSvc := svcImpl.NewUsersService(userRepo)
	authSvc := svcImpl.NewAuthService(userRepo)

	userCr := controllers.NewUserController(userSvc)
	authCr := controllers.NewAuthController(authSvc)
	routes.InitRoute(e, userCr, authCr)

}
