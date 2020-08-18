package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymohl-cl/cineplus/pkg/ghibli"
)

func TestCache_Update(t *testing.T) {
	movies := &[]ghibli.Movie{ghibli.Movie{ID: "movie1"}, ghibli.Movie{ID: "movie2"}}
	peoples := &[]ghibli.People{ghibli.People{ID: "people1"}, ghibli.People{ID: "people2"}}
	t.Run("Should be ok", func(t *testing.T) {
		c := &Cache{}
		c.Update(movies, peoples)
		assert.Equal(t, movies, c.movies)
		assert.Equal(t, peoples, c.peoples)
	})
}

func TestCache_Data(t *testing.T) {
	c := &Cache{}
	c.movies = &[]ghibli.Movie{ghibli.Movie{ID: "movie1"}, ghibli.Movie{ID: "movie2"}}
	c.peoples = &[]ghibli.People{ghibli.People{ID: "people1"}, ghibli.People{ID: "people2"}}
	t.Run("Should be ok", func(t *testing.T) {
		movies, peoples := c.Data()
		assert.EqualValues(t, movies, *c.movies)
		assert.EqualValues(t, peoples, *c.peoples)
	})
}
