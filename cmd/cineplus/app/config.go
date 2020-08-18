package app

import "github.com/ymohl-cl/cineplus/pkg/config"

// Config api cineplus
type Config struct {
	RefreshTime int64  `required:"true"`
	Port        string `required:"true"`
}

// NewConfig parse the environment values to return a initialized configuration
func NewConfig(appName string) (Config, error) {
	var err error
	var c Config

	if err = config.ParseEnv(appName, &c); err != nil {
		return Config{}, err
	}
	return c, nil
}
