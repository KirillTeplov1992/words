package store

type Config struct{
	DatabaseURL string
}

func NewConfig() *Config{
	return &Config{
		DatabaseURL: "web:3758@/words?parseTime=true",
	}
}