package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *Configuration // publica

type Configuration struct {
	Server ServerConfiguration
	Domain DomainConfiguration
}

type DomainConfiguration struct {
	Google           string
	GrpcCategory     string
	GrpcAuth         string
	GrpcNotification string
}

type ServerConfiguration struct {
	Port   string
	Secret string
	Mode   string
}

// SetupDB initialize configuration
func Setup(configPath string) {
	var configuration *Configuration // private
	delete_vlog()
	delete_manifest()
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error al leer el archivo de configuración, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("No se puede decodificar en estructura, %v", err)
	}
	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
