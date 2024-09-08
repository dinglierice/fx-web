package conf

type Config struct {
	HttpPort    string
	Environment string
	LogLevel    string
	DB          DBConfig
}

type DBConfig struct {
	Driver          string
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int // 以秒为单位
	AutoMigrate     bool
}

func ProvideDevConfig() *Config {
	return &Config{
		HttpPort:    ":8080",
		Environment: "development",
		LogLevel:    "debug",
	}
}
