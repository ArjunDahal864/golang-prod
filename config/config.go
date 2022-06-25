package config

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBUser        string `mapstructure:"POSTGRES_USER"`
	DBPort        string `mapstructure:"POSTGRES_PORT"`
	DBPassword    string `mapstructure:"POSTGRES_PASSWORD"`
	DBHost        string `mapstructure:"POSTGRES_HOST"`
	DBName        string `mapstructure:"POSTGRES_DATABASE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}
