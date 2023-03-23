package validation

import (
	pbUser "api-gateway/internal/infra/proto/user"
	validatorPer "api-gateway/internal/infra/validator"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateUserUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		user := new(pbUser.UserLogin)

		_ = c.Bind(&user)
		if err := v.Struct(user); err != nil {
			errs := err.(validator.ValidationErrors)

			mapErrors := map[string]interface{}{"errors": validatorPer.GenerateMessage(v, errs)}
			return c.JSON(http.StatusBadRequest, mapErrors)
		}
		c.Set("data-user", user)
		return next(c)
	}
}
