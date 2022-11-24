package utils

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBUrl                string        `mapstructure:"DB_URL"`
	Host                 string        `mapstructure:"HOST"`
	JwtSecretKey         string        `mapstructure:"JWT_SECRET"`
	JwtRefreshSecretKey  string        `mapstructure:"JWT_REFRESH"`
	PasetoSymmetricKey   string        `mapstructure:"PASETO_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	FrontendUrl          string        `mapstructure:"FRONTEND_URL"`
	CldUrl               string        `mapstructure:"CLD_URL"`
}

func LoadConfig(path, configName string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = viper.Unmarshal(&config)

	if err != nil {
		return Config{}, err
	}

	log.Println("Loaded config file")
	return config, nil
}
