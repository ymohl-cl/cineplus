package ghibli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeople_IsInMovie(t *testing.T) {
	movie1 := "90b72513-afd4-4570-84de-a56c312fdf81"
	movie2 := "0440483e-ca0e-4120-8c50-4c8cd9b965d6"
	people := People{
		FilmsURL: []string{
			"https://ghibliapi.herokuapp.com/films/0440483e-ca0e-4120-8c50-4c8cd9b965d6",
			"https://ghibliapi.herokuapp.com/films/34277bec-7401-43fa-a00a-5aee64b45b08",
		},
	}

	t.Run("Should return false because people is not in the movie "+movie1, func(t *testing.T) {
		ok := people.IsInMovie(movie1)
		assert.Equal(t, false, ok)
	})
	t.Run("Should be ok int the movie "+movie2, func(t *testing.T) {
		ok := people.IsInMovie(movie2)
		assert.Equal(t, true, ok)
	})
}
