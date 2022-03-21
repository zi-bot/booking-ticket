package configs

import (
	"github.com/joho/godotenv"
	"goers/execption"
	"os"
)

type Configuration interface {
	Get(key string) string
}

type config struct {
}

func New(filenames ...string) Configuration {
	err := godotenv.Load(filenames...)
	execption.PanicError(err)
	return &config{}
}

func (conf *config) Get(key string) string {
	return os.Getenv(key)
}
