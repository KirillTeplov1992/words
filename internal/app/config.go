package app

import(
	"words/internal/store"
)

type Config struct {
	BindAddr string
	LogLevel string
	Store *store.Config
}

func NewConfig() *Config{
	return &Config{
		BindAddr: ":5000",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}