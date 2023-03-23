package user_controller

import (
	su "api-gateway/internal/domain/service/user"

	cases "api-gateway/internal/domain/usecase"
	pbUser "api-gateway/internal/infra/proto/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userEntry struct {
	userCaseuse cases.UserUseCase
}

func NewUserEntry(userCaseuse cases.UserUseCase) *userEntry {
	return &userEntry{
		userCaseuse,
	}
}

func (u *userEntry) Verify(c echo.Context) error {
	data, status := u.userCaseuse.VerifyUser(c)
	return c.JSON(status, data)
}

func CreateUser(c echo.Context) error {
	user := c.Get("data-user").(*pbUser.User)
	serviceUser := su.NewServiceUser()
	response := serviceUser.CreateUser(user) //respuesta del service
	return c.JSON(int(response.GetStatus()), response)
}

func UpdateUser(c echo.Context) error {
	user := c.Get("data-user").(*pbUser.UserLogin)
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		fmt.Println("There is no valid identification number ID")
	}
	serviceUser := su.NewServiceUser()
	response := serviceUser.UpdateUser(&pbUser.UpdateRequestUser{
		Id:   int64(ID),
		User: user,
	})
	return c.JSON(int(response.GetStatus()), response)
}

func ListUser(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParams().Get("offset"))
	if err != nil {
		offset = 10
		fmt.Println(err)
	}
	serviceUser := su.NewServiceUser()
	response := serviceUser.ListUser(int32(offset))
	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		fmt.Println("There is no valid identification number ID")
	}
	serviceUser := su.NewServiceUser()
	response := serviceUser.DeleteUser(int64(ID))
	return c.JSON(http.StatusOK, response)
}

func GetByIdUser(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		fmt.Println("There is no valid identification number ID")
	}
	serviceUser := su.NewServiceUser()
	response := serviceUser.GetByIdUser(&pbUser.GetById{
		Id: int64(ID),
	})

	return c.JSON(http.StatusOK, response)
}

func (u *userEntry) Auth(c echo.Context) error {

	userLogin := c.Get("user-data").(*pbUser.FindVerifRequest)

	serviceUser := su.NewServiceUser()
	response := serviceUser.Authservice(c.Request().Context(), userLogin)

	data, status := u.userCaseuse.LoginUser(c, response, userLogin)
	return c.JSON(status, data)
}
