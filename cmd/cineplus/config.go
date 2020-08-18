package main

import "github.com/kelseyhightower/envconfig"

// Config api cineplus
type Config struct {
	RefreshTime int64 `requires:"true"`
}

// NewConfig parse the environment values to return a initialized configuration
func NewConfig(appName string) (Config, error) {
	var err error
	var c Config

	if err = envconfig.Process(appName, &c); err != nil {
		return Config{}, err
	}
	return c, nil
}
