package conf

// Basic Defination of Config
type Config struct {
	Database SQLDatabase `toml:"database"`
	Logger   Logger      `toml:"logger"`
	Server   Server      `toml:"server"`
}

type SQLDatabase struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Name     string `toml:"name"`
	Driver   string `toml:"driver"`
}

type Server struct {
	Listen string `toml:"listen"`
}

type Logger struct {
	Level string `toml:"level"`
}
