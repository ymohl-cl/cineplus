package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// ParseEnv get the prefix and the configuration object to search and fill the model structure
func ParseEnv(prefix string, conf interface{}) error {
	var err error

	if conf == nil {
		return errors.Errorf("Undefined structure to parse a new configuration %s", prefix)
	}
	if err = envconfig.Process(prefix, conf); err != nil {
		return err
	}
	return nil
}
