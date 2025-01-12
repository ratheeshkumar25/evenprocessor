package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	URL string `mapstructure:"URL"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}

	// Validate URL
	if config.URL == "" {
		log.Fatalf("Webhook URL is not set in config")
	}

	log.Printf("Loaded webhook URL: %s", config.URL)

	return &config
}
