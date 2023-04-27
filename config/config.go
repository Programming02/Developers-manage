package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host                string `envconfig:"HOST" required:"true"`
	Port                string `envconfig:"PORT" required:"true"`
	CasbinConfigPath    string `envconfig:"CASBIN_CONFIG_PATH" required:"true"`
	MiddlewareRolesPath string `envconfig:"MIDDLEWERE_ROLES_PATH" required:"true"`
	SigningKey          string `envconfig:"SIGNIN_KEY" required:"true"`
	JWTSecretKey        string `envconfig:"JWT_SECRET_KEY" required:"true"`

	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int
	PostgresConfig
}

type PostgresConfig struct {
	PostgresHost           string `envconfig:"POSTGRES_HOST" required:"true"`
	PostgresPort           string `envconfig:"POSTGRES_PORT" required:"true"`
	PostgresUser           string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPassword       string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresDB             string `envconfig:"POSTGRES_DB" required:"true"`
	PostgresMigrationsPath string `envconfig:"POSTGRES_MIGRATIONS_PATH" required:"true"`
}

func Load() Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}
	}
	return cfg
}
