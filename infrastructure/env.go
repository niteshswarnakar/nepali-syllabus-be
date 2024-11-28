package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	ServerPort       string `mapstructure:"SERVER_PORT"`
}

func MyEnv(env_string string) *Env {
	if env_string == "" {
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigFile(env_string)
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Failed to read env file data ", err.Error())
	}

	env := &Env{}
	err = viper.Unmarshal(env)
	if err != nil {
		log.Fatalln("Failed to unmarshal env data ", err.Error())
	}

	return env

}
