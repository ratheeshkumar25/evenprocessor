package config

import "github.com/spf13/viper"

type Config struct {
	URL string `mapstructure:"URL"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return &config
}
