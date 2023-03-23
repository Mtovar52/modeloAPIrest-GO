package validation

import (
	pbSend "api-gateway/internal/infra/proto/sendGrid"
	validator_Send "api-gateway/internal/infra/validator"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateSend(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator_Send.NewValidator()
		send := new(pbSend.Send)

		_ = c.Bind(&send)
		if err := v.Struct(send); err != nil {
			errs := err.(validator.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validator_Send.GenerateMessage(v, errs))
		}
		c.Set("send-email", send)
		return next(c)
	}
}
