package config

import "time"

type Config struct {
	DatabaseHost              string        `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort              string        `env:"DATABASE_PORT" envDefault:"27017"`
	DatabaseUserName          string        `env:"DATABASE_USER_NAME" envDefault:"root"`
	DatabasePassword          string        `env:"DATABASE_PASSWORD" envDefault:"pAsswordFORroot"`
	DatabaseConnectionTimeout time.Duration `env:"DATABASE_CONNECTION_TIMEOUT" envDefault:"10s"`
}
