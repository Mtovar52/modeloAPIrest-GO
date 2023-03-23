package validation

import (
	pbUser "api-gateway/internal/infra/proto/user"
	validatorPer "api-gateway/internal/infra/validator"
	hash "api-gateway/internal/middlewares"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		user := new(pbUser.User)

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

func ValidateNamePasswCREATE(user *pbUser.User) *pbUser.UserResponse {

	passwordUpdate, errs := hash.HashPassword(user.GetPassword())
	user.Password = passwordUpdate

	if errs != nil {
		return &pbUser.UserResponse{}
	}
	return nil
}

func ValidatePasswUserUPDATE(u *pbUser.UserLogin) *pbUser.UserResponse {

	password, err := hash.HashPassword(u.GetPassword())
	u.Password = password

	if err != nil {
		return &pbUser.UserResponse{}
	}
	return nil
}

func ValidatePaswords(u *pbUser.User) *pbUser.UserResponse {
	if u.Password != u.Password2 {
		return &pbUser.UserResponse{
			Title:   "password error",
			IsOk:    false,
			Message: "Passwords do not match",
			Status:  http.StatusNotAcceptable,
		}
	}
	return nil
}

func ValidateConditions(u *pbUser.User) *pbUser.UserResponse {
	contitions := u.GetConditions()
	if contitions == false {
		return &pbUser.UserResponse{
			Message: "Conditions should be True",
			Status:  http.StatusNotFound,
		}
	}
	return nil
}

func CompareHashAndPassword(hashedPassword, password []byte) error {
	return nil
}
