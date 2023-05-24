package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DebugMode bool `mapstructure:"DEBUG_MODE"`

	AppHost string `mapstructure:"APP_HOST"`
	AppPort int    `mapstructure:"APP_PORT"`

	MongoUrl string `mapstructure:"MONGO_URL"`
	DBName   string `mapstructure:"DB_NAME"`
}

var config *Config

func GetConfig(path string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error happened during config read. \n %+v1", err))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error happened during config unmarshal. \n %+v1", err))
	}

	return config
}
