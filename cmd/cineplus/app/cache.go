package app

import (
	"sync"

	"github.com/ymohl-cl/cineplus/pkg/ghibli"
)

// Cache cineplus data
type Cache struct {
	movies  *[]ghibli.Movie
	peoples *[]ghibli.People
	sync.Mutex
}

// Update the data store in the cache
func (c *Cache) Update(movies *[]ghibli.Movie, peoples *[]ghibli.People) {
	c.Lock()
	defer c.Unlock()

	c.movies = movies
	c.peoples = peoples
}

// Data return a cache copy to prevent latency from an update
func (c *Cache) Data() (movies []ghibli.Movie, peoples []ghibli.People) {
	c.Lock()
	defer c.Unlock()

	if c.movies != nil {
		movies = *(c.movies)
	}
	if c.peoples != nil {
		peoples = *(c.peoples)
	}
	return
}
