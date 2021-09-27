package chat

//Конфигурация проекта
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LoggerLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":5555",
		LoggerLevel: "debug",
	}
}

