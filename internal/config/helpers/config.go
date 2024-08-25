package helpers

import (
	"github.com/codepnw/server-template/internal/types/config"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (config.AppConfig, error) {
	if GetEnv("AppMode", "debug") == "release" {
		c := config.AppConfig{
			PostgresDBName:     GetEnv("PostgresDBName", ""),
			PostgresDBUser:     GetEnv("PostgresDBUser", ""),
			PostgresDBPassword: GetEnv("PostgresDBPassword", ""),
			PostgresDBPort:     5432,
			PostgresDBSSLMode:  GetEnv("PostgresDBSSLMode", ""),
			CookieDomain:       GetEnv("CookieDomain", ""),
			CookieHttpOnly:     GetEnvAsBool("CookieHttpOnly"),
			CorsOrigin:         GetEnv("CorsOrigin", ""),
			SQLDBName:          GetEnv("SQLDBName", ""),
			SQLDBUser:          GetEnv("SQLDBUser", ""),
			SQLDBPassword:      GetEnv("SQLDBPassword", ""),
			SQLDBPort:          5432,
			SQLSSLMode:         GetEnv("SQLSSLMode", ""),
		}
		return c, nil
	}
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config.AppConfig{}, err
	}

	var cfg config.AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return config.AppConfig{}, err
	}
	return cfg, nil
}
