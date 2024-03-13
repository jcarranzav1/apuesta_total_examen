package config

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort        int           `required:"true" split_words:"true"`
	Postfix           string        `required:"true"`
	DbHost            string        `required:"true" split_words:"true"`
	DbPort            int           `required:"true" split_words:"true"`
	DbUser            string        `required:"true" split_words:"true"`
	DbPassword        string        `required:"true" split_words:"true"`
	DbName            string        `required:"true" split_words:"true"`
	JwtExpirationTime time.Duration `required:"true" split_words:"true"`
	JwtSecretKey      []byte        `required:"true" split_words:"true"`
}

var once sync.Once
var Cfg Config

func Environments() Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}

		if err := envconfig.Process("", &Cfg); err != nil {
			log.Panicf("Error parsing environment vars %#v", err)
		}
	})

	return Cfg
}

func IsDevEnvironment() bool {
	return strings.EqualFold(Environments().Postfix, "dev")
}
