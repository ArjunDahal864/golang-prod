package conf

// Basic Defination of Config
type Config struct {
	Database Database `toml:"database"`
	Log      Logger   `toml:"loger"`
	Server   Server   `toml:"server"`
}

type Database struct {
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