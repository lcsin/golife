package config

type Database struct {
	Driver string `mapstructure:"driver"`
	Source string `mapstructure:"source"`
}
