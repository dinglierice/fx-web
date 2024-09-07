package conf

type Config struct {
	HttpPort    string
	Environment string
	LogLevel    string
}

func ProvideDevConfig() *Config {
	return &Config{
		HttpPort:    ":8080",
		Environment: "development",
		LogLevel:    "debug",
	}
}
