package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Field string `required:"true"`
}

func TestParseEnv(t *testing.T) {
	var c TestConfig

	t.Run("should return an error cause a nil configuration", func(t *testing.T) {
		expectedError := "Undefined structure to parse a new configuration tester"
		err := ParseEnv("tester", nil)
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err.Error())
		}
	})
	t.Run("should return an error cause required value not set", func(t *testing.T) {
		expectedError := "required key TESTER_FIELD missing value"
		err := ParseEnv("tester", &c)
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err.Error())
		}
	})
	t.Run("should be ok", func(t *testing.T) {
		os.Setenv("TESTER_FIELD", "TEST_OK")
		defer os.Unsetenv("TESTER_FIELD")
		expectedFieldValue := "TEST_OK"
		err := ParseEnv("tester", &c)
		if assert.NoError(t, err) {
			assert.Equal(t, expectedFieldValue, c.Field)
		}
	})
}
