package jwt

import (
	objectValues "api-gateway/internal/domain/object_values"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

func NewJwtClient() JwtClient {
	return JwtClient{}
}

type JwtClient struct{}

func (j *JwtClient) GenerateToken(entityJWT objectValues.JwtEntity) (string, error) {

	claims := objectValues.JwtCustomClaims{
		IdUser: entityJWT.IdUser,
		Nombre: entityJWT.Nombre,
		Email:  entityJWT.Email,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte("secret"))

}

func (j *JwtClient) GetConfig() middleware.JWTConfig {

	config := middleware.JWTConfig{
		Claims:     &objectValues.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	return config
}
