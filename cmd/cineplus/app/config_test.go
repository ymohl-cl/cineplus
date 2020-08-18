package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("should return an error cause REFRESHTIME required field is not set", func(t *testing.T) {
		expectedError := "required key TESTER_REFRESHTIME missing value"
		c, err := NewConfig("tester")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err.Error())
			assert.Equal(t, Config{}, c)
		}
	})
	t.Run("should return an error cause PORT required field is not set", func(t *testing.T) {
		os.Setenv("TESTER_REFRESHTIME", "60")
		defer os.Unsetenv("TESTER_REFRESHTIME")
		expectedError := "required key TESTER_PORT missing value"
		c, err := NewConfig("tester")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err.Error())
			assert.Equal(t, Config{}, c)
		}
	})
	t.Run("should be ok", func(t *testing.T) {
		os.Setenv("TESTER_REFRESHTIME", "60")
		defer os.Unsetenv("TESTER_REFRESHTIME")
		os.Setenv("TESTER_PORT", "8000")
		defer os.Unsetenv("TESTER_PORT")
		c, err := NewConfig("tester")
		if assert.NoError(t, err) {
			assert.Equal(t, Config{RefreshTime: 60, Port: "8000"}, c)
		}
	})
}
