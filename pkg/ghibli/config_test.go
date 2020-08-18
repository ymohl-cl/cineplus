package ghibli

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("should return an error cause URL required field is not set", func(t *testing.T) {
		expectedError := "required key TESTER_URL missing value"
		c, err := NewConfig("tester")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err.Error())
			assert.Equal(t, Config{}, c)
		}
	})
	t.Run("should be ok", func(t *testing.T) {
		os.Setenv("TESTER_URL", "TEST_OK")
		defer os.Unsetenv("TESTER_URL")
		expectedFieldValue := "TEST_OK"
		c, err := NewConfig("tester")
		if assert.NoError(t, err) {
			assert.Equal(t, Config{URL: expectedFieldValue}, c)
		}
	})
}
