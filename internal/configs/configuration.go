package configs

import (
	"os"
	"strings"
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

type SecurityConfig struct {
	AdminUser       string
	AdminSecret     string
	AdminMfaEnabled bool
	AdminMfaSecret  string
	CookieSecret    string
}

func NewSecurityConfig() *SecurityConfig {
	adminUser := stringValueOrDefault(os.Getenv("ADMIN_USER"), "49eaa6bf50e138fd")
	adminSecret := stringValueOrDefault(os.Getenv("ADMIN_SECRET"), "a05a0e57a4ce0269e9d99518cfb9e")
	adminMfaEnabled := strings.ToLower(stringValueOrDefault(os.Getenv("ADMIN_MFA_ENABLED"), "true")) == "true"
	adminMfaSecret := stringValueOrDefault(os.Getenv("ADMIN_MFA_SECRET"), "a05a0e57a4ce0269e9d99518cfb9e3421b9d2497b70381213eb311483721c343")
	cookieSecret := stringValueOrDefault(os.Getenv("ADMIN_COOKIE_SECRET"), "f9f11106-ea2f-4278-9a3b-ed1e3c641519")
	return &SecurityConfig{
		AdminUser:       adminUser,
		AdminSecret:     adminSecret,
		AdminMfaEnabled: adminMfaEnabled,
		AdminMfaSecret:  adminMfaSecret,
		CookieSecret:    cookieSecret,
	}
}

type DatabaseApplicationConfig struct {
	Dialect          string
	ConnectionString string
	ShowSql          bool
}

func NewDatabaseConfig() *DatabaseApplicationConfig {
	dialect := stringValueOrDefault(os.Getenv("DB_DIALECT"), "sqlite3")
	connectionString := stringValueOrDefault(os.Getenv("DB_CONNECTION_STRING"), "local.db")
	showSql := strings.ToLower(stringValueOrDefault(os.Getenv("DB_SHOW_SQL"), "true")) == "true"
	return &DatabaseApplicationConfig{
		Dialect:          dialect,
		ConnectionString: connectionString,
		ShowSql:          showSql,
	}
}
