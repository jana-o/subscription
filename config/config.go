package config

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {

	return Config{
		DBDriver: "postgres",
		DBSource: "postgres://postgres:password@localhost:5433/postgres?sslmode=disable",
		ServerAddress: "0.0.0.0:8080",
	}, nil
}
