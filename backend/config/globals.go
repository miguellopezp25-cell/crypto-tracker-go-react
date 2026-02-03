package config

import (
	"fmt"

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
	viper.AddConfigPath(".")   // Opción 1: Directorio raíz
	viper.AddConfigPath("..")  // PARA LOS TESTS (sube un nivel desde /service a /backend)
	viper.AddConfigPath("../") // Opción 3: Un nivel arriba (por si corres desde un subdirectorio)
	viper.SetConfigName("config")
	viper.AddConfigPath("./config") // Opción 2: Carpeta llamada config
	viper.SetConfigType("yaml")
	// En tu función de carga de configuración:
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Si no hay archivo, no pasa nada (útil para tests)
			fmt.Println("Aviso: Trabajando sin archivo de configuración físico")
		} else {
			return nil, err // Si es otro error, sí reportalo
		}
	}

	err = viper.Unmarshal(&conf)
	return conf, err

}
