package middlewares

import (
	objectValues "api-gateway/internal/domain/object_values"

	jwtC "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CacheTokenUser(c echo.Context) int64 {

	user := c.Get("user").(*jwtC.Token)
	var userID int64

	if user.Valid {
		claims := user.Claims.(*objectValues.JwtCustomClaims)
		userID = claims.IdUser

	}
	return userID
}
