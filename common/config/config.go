package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type config struct {
	Token string `env:"TOKEN"`
}

var Config config

func Load() error {
	err := godotenv.Load()

	opts := env.Options{RequiredIfNoDef: true}

	err = env.Parse(&Config, opts)

	return err
}
