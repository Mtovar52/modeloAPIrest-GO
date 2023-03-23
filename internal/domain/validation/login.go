package validation

import (
	"net/http"

	objectValues "api-gateway/internal/domain/object_values"
	userLogin "api-gateway/internal/infra/proto/user"
	validatorPer "api-gateway/internal/infra/validator"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

func ValidateLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		user := new(userLogin.FindVerifRequest)

		_ = c.Bind(&user)
		if err := v.Struct(user); err != nil {
			errs := err.(validator.ValidationErrors)
			mapErrors := map[string]interface{}{"errors": validatorPer.GenerateMessage(v, errs)}
			return c.JSON(http.StatusBadRequest, mapErrors)
		}
		c.Set("user-data", user)
		return next(c)
	}
}

func ValidateVerify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		verify := new(objectValues.Verify)

		_ = c.Bind(&verify)
		if err := v.Struct(verify); err != nil {
			errs := err.(validator.ValidationErrors)
			mapErrors := map[string]interface{}{"errors": validatorPer.GenerateMessage(v, errs)}
			return c.JSON(http.StatusBadRequest, mapErrors)
		}
		c.Set("verify", verify)
		return next(c)
	}
}
