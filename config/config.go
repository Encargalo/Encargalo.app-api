package config

import (
	"log"
	"sync"

	"github.com/andresxlp/gosuite/config"
)

var (
	Once sync.Once
	cfg  *Config
)

func Get() *Config {
	if cfg == nil {
		log.Panic("Configuration has not yet been initialized")
	}
	return cfg
}

type Config struct {
	Server   Server   `env:"server"`
	Postgres Postgres `env:"postgres"`
	Mongo    Mongo    `env:"mongo"`
	Redis    Redis    `env:"redis"`
	Twilio   Twilio   `env:"twilio"`
}

type Server struct {
	Port int `env:"port"`
}

type Postgres struct {
	Host     string `env:"host"`
	Port     int    `env:"port"`
	User     string `env:"user"`
	Password string `env:"password"`
	Database string `env:"database"`
}

type Mongo struct {
	Url string `env:"url"`
}

type Redis struct {
	Host     string `env:"host"`
	Port     int    `env:"port"`
	User     string `env:"user"`
	Password string `env:"password"`
}

type Twilio struct {
	Username string `env:"username"`
	Password string `env:"password"`
	Phone    string `env:"phone"`
}

func Environments() {
	Once.Do(func() {
		cfg = new(Config)
		if err := config.GetConfigFromEnv(cfg); err != nil {
			log.Panicf("error parsing enviroment vars \n%v", err)
		}
	})
}
