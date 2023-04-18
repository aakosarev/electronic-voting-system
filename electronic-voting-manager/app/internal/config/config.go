package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	Blockchain struct {
		URL        string `yaml:"url"`
		PrivateKey string `yaml:"private_key"`
		GasLimit   int64  `yaml:"gas_limit"`
		GasPrice   int64  `yaml:"gas_price"`
		WeiFounds  int64  `yaml:"default_wei_founds"`
	} `yaml:"blockchain"`
	GRPC struct {
		IP   string `yaml:"ip"`
		Port int    `yaml:"port"`
	} `yaml:"grpc"`
	PostgreSQL struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"postgresql"`
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
