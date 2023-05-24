package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBHost         string `mapstructure:"DB_HOST"`
	DBUserName     string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	JwtSecret           string `mapstructure:"JWT_SECRET"`
	JwtRefreshSecret    string `mapstructure:"JWT_REFRESH_SECRET"`
	JwtExpiresIn        int    `mapstructure:"JWT_EXPIRED_IN"`
	JwtRefreshExpiresIn int    `mapstructure:"JWT_REFRESH_EXPIRED_IN"`
}

func NewEnv() (*Env, error) {
	env := Env{}
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file config.yaml : ", err)
		return &env, err
	}
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
		return &env, err
	}
	return &env, nil
}
