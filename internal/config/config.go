package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env_required:"true"`
	Listen  struct {
		Type string `yaml:"type" env_default:"port"`
		Host string `yaml:"host" env_default:"127.0.0.1"`
		Port string `yaml:"port" env_default:"8080"`
	} `yaml:"listen"`
}

var instance Config
var once sync.Once

func InitConfig() *Config {
	once.Do(func() {
		log.Println("reading config")
		err := cleanenv.ReadConfig("config.yml", &instance)
		if err != nil {
			log.Fatal(err)
		}
	})
	return &instance
}
