package objectValues

import (
	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	IdUser int64  `json:"userId"`
	Nombre string `json:"name"`
	Email  string `json:"email"`

	jwt.StandardClaims
}

type JwtEntity struct {
	IdUser int64  `json:"userId"`
	Nombre string `json:"name"`
	Email  string `json:"email"`
}
