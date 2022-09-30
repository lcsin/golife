package config

import (
	"log"

	"github.com/spf13/viper"
)

var cfg AppConfig

type AppConfig struct {
	Database    Database    `mapstructure:"database"`
	Application Application `mapstructure:"application"`
	Zap         Zap         `mapstructure:"zap"`
}

type Application struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

func Init(path string) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Read config file failed: %s", err)
	}
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Resolve config file failed: %s", err)
	}
}

func App() *AppConfig {
	return &cfg
}
