package conf

// Basic Defination of Config
type Config struct {
	Database SQLDatabase  
	Logger   Logger      
	Server   Server
}

type SQLDatabase struct {
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DATABASE_PORT"`
	Name     string `env:"POSTGRES_DB"`
}

type Server struct {
	Listen string `env:"SERVER_EXPOSE_PORT"`
}

type Logger struct {
	Level string `env:"LOG_LEVEL"`
}
