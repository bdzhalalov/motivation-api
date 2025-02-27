package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Addr     string `mapstructure:"addr"`
	LogLevel string `mapstructure:"log_level"`
	APIKey   string `mapstructure:"api_key"`
	DbName   string `mapstructure:"DB_NAME"`
	DbHost   string `mapstructure:"DB_HOST"`
	DbPort   string `mapstructure:"DB_PORT"`
	DbUser   string `mapstructure:"DB_USER"`
	DbPass   string `mapstructure:"DB_PASSWORD"`
}

var Cfg Config

func InitConfig() Config {

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}

	err := viper.Unmarshal(&Cfg)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
	}

	return Cfg
}
