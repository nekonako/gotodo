package config

import (
	"os"
	"todo/exception"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImp struct {
}

func (c *configImp) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicIfErr(err)
	return &configImp{}
}
