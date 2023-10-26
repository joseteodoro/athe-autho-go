package configs

import (
	"os"
)

func stringValueOrDefault(value, defaultValue string) string {
	if value != "" {
		return value
	}
	return defaultValue
}

type ApplicationConfig struct {
	Port    string
	Version string
}

func NewApplicationConfig() *ApplicationConfig {
	port := stringValueOrDefault(os.Getenv("PORT"), "8080")
	version := stringValueOrDefault(os.Getenv("APPLICATION_VERSION"), "")
	return &ApplicationConfig{
		Port:    port,
		Version: version,
	}
}

type DatabaseApplicationConfig struct {
	Dialect          string
	ConnectionString string
}

func NewDatabaseConfig() *DatabaseApplicationConfig {
	dialect := stringValueOrDefault(os.Getenv("DB_DIALECT"), "sqlite3")
	connectionString := stringValueOrDefault(os.Getenv("DB_CONNECTION_STRING"), "local.db")
	return &DatabaseApplicationConfig{
		Dialect:          dialect,
		ConnectionString: connectionString,
	}
}
