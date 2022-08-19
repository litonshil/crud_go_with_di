package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string `mapstructure:"PORT"`
	SqlUri   string `mapstructure:"SQL_URI"`
	SqlDb    string `mapstructure:"SQL_DB_NAME"`
	Username string `mapstructure:"USER_NAME"`

	//SMTP
	Email        string `mapstructure:"EMAIL"`
	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPort     string `mapstructure:"SMTP_PORT"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`
	SecretKey    string `mapstructure:"SECRETKEY"`

	// TEST
	Mode         string `mapstructure:"MODE"`
	POSTGRES_URL string `mapstructure:"POSTGRES_URL"`

	// REDIS
	RedisHost   string `mapstructure:"REDIS_HOST"`
	RedisPort   string `mapstructure:"REDIS_PORT"`
	RedisPass   string `mapstructure:"REDIS_PASS"`
	UserPrefix  string `mapstructure:"USER_PREFIX"`
	TokenPrefix string `mapstructure:"TOKEN_PREFIX"`
}

func init_config() (config Config, err error) {
	viper.AddConfigPath(".") //for test path: ../../ and for main its: .
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return

}

func GetConfig() Config {
	config, err := init_config()
	if err != nil {
		fmt.Println(err.Error())
	}

	return config
}
