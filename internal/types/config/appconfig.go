package config

type AppConfig struct {
	PostgresDBName     string `mapstructure:"PostgresDBName"`
	PostgresDBUser     string `mapstructure:"PostgresDBUser"`
	PostgresDBPassword string `mapstructure:"PostgresDBPassword"`
	PostgresDBPort     int    `mapstructure:"PostgresDBPort"`
	PostgresDBSSLMode  string `mapstructure:"PostgresDBSSLMode"`
	CookieDomain       string `mapstructure:"CookieDomain"`
	CookieHttpOnly     bool   `mapstructure:"CookieHttpOnly"`
	CorsOrigin         string `mapstructure:"CorsOrigin"`
	SQLDBName          string `mapstructure:"SQLDBName"`
	SQLDBUser          string `mapstructure:"SQLDBUser"`
	SQLDBPassword      string `mapstructure:"SQLDBPassword"`
	SQLDBPort          int    `mapstructure:"SQLDBPort"`
	SQLSSLMode         string `mapstructure:"SQLSSLMode"`
}
