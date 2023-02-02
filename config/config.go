package config

import (
	"strings"

	"github.com/spf13/viper"
)

type appenv string

var development = appenv("development")
var production = appenv("production")

type Config struct {
	AppEnv appenv `mapstructure:"app_env"`

	Database struct {
		URL string
	}

	Server struct {
		Port string
	} `mapstructure:",squash"`
}

// Addr returns the address the server should listen on.
func (c *Config) Addr() string {
	return ":" + c.Server.Port
}

// IsDev returns true if the app is running in development mode.
func (c *Config) IsDev() bool {
	return c.AppEnv == development
}

// IsProd returns true if the app is running in production mode.
func (c *Config) IsProd() bool {
	return c.AppEnv == production
}

func New() (*Config, error) {
	v := viper.New()
	v.AutomaticEnv()
	v.AllowEmptyEnv(true)
	v.SetDefault("app_env", development)
	v.SetDefault("database.url", nil)
	v.SetDefault("port", "3000")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	v.SetConfigName("config")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); err != nil && !ok {
		return nil, err
	}

	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
