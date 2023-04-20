package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	HTTP struct {
		IP   string `yaml:"ip"`
		Port int    `yaml:"port"`
	} `yaml:"http"`
	VotingManagerGRPC struct {
		IP   string `yaml:"ip"`
		Port int    `yaml:"port"`
	} `yaml:"voting-manager-grpc"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
