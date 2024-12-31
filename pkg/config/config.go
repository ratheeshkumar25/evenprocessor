package config

type Config struct {
	Host     string `mapstructure:"HOST"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Database string `mapstructure:"DBNAME"`
	Port     string `mapstructure:"PORT"`
	Sslmode  string `mapstructure:"SSL"`
}
