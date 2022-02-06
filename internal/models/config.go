package models

type Config struct {
	Token       string `env:"TOKEN"`
	DatabaseUri string `env:"DATABASE_URI"`
}