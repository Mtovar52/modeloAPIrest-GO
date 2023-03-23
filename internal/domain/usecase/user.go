package usecase

import (
	objectValues "api-gateway/internal/domain/object_values"
	"api-gateway/internal/infra/cache"
	"api-gateway/internal/infra/jwt"
	pbUser "api-gateway/internal/infra/proto/user"
	middle "api-gateway/internal/middlewares"
	"net/http"

	jwtC "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserUseCase struct {
	jwt   jwt.JwtClient
	cache *cache.CacheProvider
}

func NewUserUserCase(jwt jwt.JwtClient, cache *cache.CacheProvider) UserUseCase {

	return UserUseCase{
		jwt:   jwt,
		cache: cache,
	}
}

func (u UserUseCase) LoginUser(a echo.Context, userPB *pbUser.UserLogin, userLogin *pbUser.FindVerifRequest) (interface{}, int) {

	hash := userPB.GetPassword()
	contraseña := userLogin.Password

	checkPasswordHash := middle.CheckPasswordHash(contraseña, hash)
	if checkPasswordHash == true {
		token, err := u.jwt.GenerateToken(objectValues.JwtEntity{
			IdUser: userPB.GetId(),
			Email:  userLogin.Email,
		})

		if err != nil {
			return objectValues.NewResponseWithData(http.StatusUnauthorized, "no authorization", "el usuario no es valido", nil), http.StatusUnauthorized
		}
		u.cache.Set("token-"+userPB.GetEmail(), token)
		return objectValues.NewResponseWithData(http.StatusOK, "ok", "usuario autorizado", map[string]string{
			"token": token,
		}), http.StatusOK
	}
	return objectValues.NewResponseWithData(http.StatusUnauthorized, "no autorizado", "el usuario no es valido", nil), http.StatusUnauthorized
}

func (u UserUseCase) VerifyUser(c echo.Context) (map[string]interface{}, int) {
	user := c.Get("user").(*jwtC.Token)
	if user.Valid {

		claims := user.Claims.(*objectValues.JwtCustomClaims)

		byteToken := u.cache.Get("token-" + claims.Email)
		if byteToken != nil {
			return map[string]interface{}{
				"valid": true,
				"data":  claims,
			}, http.StatusOK
		}
	}
	return map[string]interface{}{
		"valid": false,
	}, http.StatusUnauthorized
}

func (u *UserUseCase) ChangePassword() {

}

func (u *UserUseCase) SendRestartEmail() {

}
