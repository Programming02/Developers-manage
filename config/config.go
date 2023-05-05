package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	Host                      string `json:"host" required:"true" envconfig:"HOST"`
	Port                      string `json:"port" required:"true" envconfig:"PORT"`
	PostgresHost              string `json:"postgresHost" required:"true" envconfig:"POSTGRES_HOST"`
	PostgresPort              string `json:"postgresPort" required:"true" envconfig:"POSTGRES_PORT"`
	PostgresUser              string `json:"postgresUser" required:"true" envconfig:"POSTGRES_USER"`
	PostgresPassword          string `json:"postgresPassword" required:"true" envconfig:"POSTGRES_PASSWORD"`
	PostgresDB                string `json:"postgresDB" required:"true" envconfig:"POSTGRES_DB"`
	CasbinConfigPath          string `envconfig:"CASBIN_CONFIG_PATH" required:"true"`
	MiddlewareRolesPath       string `envconfig:"MIDDLEWERE_ROLES_PATH" required:"true"`
	SigningKey                string `envconfig:"SIGNIN_KEY" required:"true"`
	JWTSecretKey              string `envconfig:"JWT_SECRET_KEY" required:"true"`
	PostgresMigrationsPath    string `envconfig:"POSTGRES_MIGRATIONS_PATH" required:"true"`
	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int
}

func Load() (Config, error) {
	return Config{
		Host:                   os.Getenv("HOST"),
		Port:                   os.Getenv("PORT"),
		PostgresHost:           os.Getenv("POSTGRES_HOST"),
		PostgresPort:           os.Getenv("POSTGRES_PORT"),
		PostgresUser:           os.Getenv("POSTGRES_USER"),
		PostgresPassword:       os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:             os.Getenv("POSTGRES_DB"),
		PostgresMigrationsPath: os.Getenv("POSTGRES_MIGRATIONS_PATH"),
	}, nil
}

//type Config struct {
//	Host                string `envconfig:"HOST" required:"true"`
//	Port                string `envconfig:"PORT" required:"true"`
//	CasbinConfigPath    string `envconfig:"CASBIN_CONFIG_PATH" required:"true"`
//	MiddlewareRolesPath string `envconfig:"MIDDLEWERE_ROLES_PATH" required:"true"`
//	SigningKey          string `envconfig:"SIGNIN_KEY" required:"true"`
//	JWTSecretKey        string `envconfig:"JWT_SECRET_KEY" required:"true"`
//
//	JWTSecretKeyExpireMinutes int
//	JWTRefreshKey             string
//	JWTRefreshKeyExpireHours  int
//	PostgresConfig
//}
//
//type PostgresConfig struct {
//	PostgresHost           string `envconfig:"POSTGRES_HOST" required:"true"`
//	PostgresPort           string `envconfig:"POSTGRES_PORT" required:"true"`
//	PostgresUser           string `envconfig:"POSTGRES_USER" required:"true"`
//	PostgresPassword       string `envconfig:"POSTGRES_PASSWORD" required:"true"`
//	PostgresDB             string `envconfig:"POSTGRES_DB" required:"true"`
//	PostgresMigrationsPath string `envconfig:"POSTGRES_MIGRATIONS_PATH" required:"true"`
//}
//
//func Load() *Config {
//	return Config{
//		Host: os.Getenv("HOST"),
//		Port: os.Getenv("PORT"),
//	}
//}
