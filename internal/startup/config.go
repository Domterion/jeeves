package startup

import (
	"github.com/caarlos0/env/v6"
	"github.com/domterion/jeeves/internal/models"
	"github.com/joho/godotenv"
)

func InitConfig() (*models.Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	opts := env.Options{RequiredIfNoDef: true}

	config := models.Config{}
	err = env.Parse(&config, opts)

	return &config, err
}
