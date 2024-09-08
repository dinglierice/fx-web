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

// ProvideDevConfig 开发环境配置
func ProvideDevConfig() *Config {
	dbConfig := DBConfig{
		Driver:          "mysql",
		Host:            "127.0.0.1",
		Port:            3306,
		User:            "root",
		Password:        "123456",
		Name:            "mall_db",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: 10,
		AutoMigrate:     true,
	}
	return &Config{
		HttpPort:    ":4000",
		Environment: "development",
		LogLevel:    "debug",
		DB:          dbConfig,
	}
}
