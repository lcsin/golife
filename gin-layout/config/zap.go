package config

type Zap struct {
	Cut   bool   `mapstructure:"cut"`
	Path  string `mapstructure:"path"`
	Style string `mapstructure:"style"`
	Error bool   `mapstructure:"error"`
}
