package entry

import (
	"api-gateway/cmd/config"
	"api-gateway/internal/infra/grpc"
	validatorPer "api-gateway/internal/infra/validator"
	"flag"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var router *echo.Echo

func init() {
	var configPath = ""
	configPath = *flag.String("config", "", "")

	if configPath == "" {
		configPath = "../data/config.yml"
	}

	setConfiguration(configPath)
}

func setConfiguration(configPath string) {
	config.Setup(configPath)

}

func Run() {
	grpc.NewCon()
	grpc.NewCon_Auth()
	validatorPer.NewValidator()
	conf := config.GetConfig()
	router = echo.New()
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())
	ioc := genIoc()
	NewHandler(router, ioc)
	router.Start(":" + conf.Server.Port)
}
