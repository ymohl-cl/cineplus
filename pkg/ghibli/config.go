package ghibli

import "github.com/kelseyhightower/envconfig"

// Config ghibli client
type Config struct {
	URL   string `required:"true"`
	Token string `required:"false"`
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
