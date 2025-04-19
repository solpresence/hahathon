package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `env:"ENV"`
	Token      string `env:"API_TOKEN"`
	Htppserver Httpserver
}

type Httpserver struct {
	Addr         string `env:"ADDR"`
	Timeout      string `env:"TIMEOUT"`
	IddleTimeout string `env:"IDLE_TIMEOUT"`
	MaxConn      string `env:"MAX_CONN"`
}

func MustLoad() *Config {
	envPath := "../../.env"

	var cfg Config
	if err := cleanenv.ReadConfig(envPath, &cfg); err != nil {
		log.Fatalf("can not read config: %s", err)
		panic(err)
	}

	return &cfg
}
