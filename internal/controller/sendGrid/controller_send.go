package sendgrid_controller

import (
	sw "api-gateway/internal/domain/service/sendGrid"
	pbSend "api-gateway/internal/infra/proto/sendGrid"

	"github.com/labstack/echo/v4"
)

func SendEmail(c echo.Context) error {
	send := c.Get("send-email").(*pbSend.Send)
	serviceSend := sw.NewServiceSend()
	response := serviceSend.SendEmail(c.Request().Context(), send)
	return c.JSON(int(response.GetStatus()), response)
}
