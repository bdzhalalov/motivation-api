package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Addr     string `mapstructure:"addr"`
	LogLevel string `mapstructure:"log_level"`
}

func InitConfig() Config {
	var cfg Config

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
	}

	return cfg
}
