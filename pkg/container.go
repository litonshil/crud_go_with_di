package continer

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/controllers"
	"github.com/litonshil/crud_go_echo/pkg/database"
	// "github.com/litonshil/crud_go_echo/pkg/repository"
	repoImpl "github.com/litonshil/crud_go_echo/pkg/repository/impl"
	svcImpl "github.com/litonshil/crud_go_echo/pkg/svc/impl"
)

func Init(e *echo.Echo) {
	db := database.GetDB()

	userRepo := repoImpl.NewUsersRepository(db)
	userSvc := svcImpl.NewUsersService(userRepo)
	authSvc := svcImpl.NewAuthService(userRepo)

	controllers.NewUserController(e,userSvc)
	controllers.NewAuthController(e, authSvc, userSvc)


}