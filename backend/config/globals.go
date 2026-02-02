package config

import (
	"github.com/spf13/viper"
)

type ApiConfig struct {
	APIURL string `mapstructure:"API_URL_CRYPTO"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}
type Config struct {
	Api    ApiConfig    `mapstructure:"api"`
	Server ServerConfig `mapstructure:"server"`
}

func LoadConfig() (conf *Config, err error) {
	viper.AddConfigPath(".")        // Opción 1: Directorio raíz
	viper.AddConfigPath("./config") // Opción 2: Carpeta llamada config
	viper.AddConfigPath("../")      // Opción 3: Un nivel arriba (por si corres desde un subdirectorio)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&conf)
	return

}
