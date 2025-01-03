package config

import "github.com/spf13/viper"

type Config struct {
	Host     string `mapstructure:"HOST"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Database string `mapstructure:"DBNAME"`
	Port     string `mapstructure:"PORT"`
	Sslmode  string `mapstructure:"SSL"`
	URL      string `mapstructure:"URL"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return &config
}
