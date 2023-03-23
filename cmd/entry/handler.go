package entry

import (
	controller_user "api-gateway/internal/controller/user"

	usecase "api-gateway/internal/domain/usecase"
	"api-gateway/internal/domain/validation"
	"api-gateway/internal/infra/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewHandler(router *echo.Echo, userUseCase usecase.UserUseCase) {
	userEntry := controller_user.NewUserEntry(userUseCase)
	confAuth := jwt.NewJwtClient()

	auth := middleware.JWTWithConfig(confAuth.GetConfig())

	router.POST("/user", controller_user.CreateUser, validation.ValidateUser)
	router.GET("/user", controller_user.ListUser, auth)
	router.PUT("/user/:ID", controller_user.UpdateUser, auth, validation.ValidateUserUpdate)
	router.DELETE("/user/:ID", controller_user.DeleteUser, auth)
	router.GET("/user/:ID", controller_user.GetByIdUser, auth)
	router.POST("/user/auth", userEntry.Auth, validation.ValidateLogin)

}
