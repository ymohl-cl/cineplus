package app

import (
	"net/http"

	"github.com/labstack/echo"
)

// Handler services struct
type Handler struct {
	cache *Cache
}

// Ping method http GET
func (h Handler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, &Pong{Pong: true})
}

// Movies return movies data from the cache
func (h Handler) Movies(c echo.Context) error {
	data := jsonDataModel(h.cache)
	return c.JSON(http.StatusOK, &data)
}

// jsonDataModel builder with the current cache data
func jsonDataModel(cache *Cache) Data {
	d := Data{}
	movies, peoples := cache.Data()

	for _, m := range movies {
		movie := Movie{
			ID:          m.ID,
			Title:       m.Title,
			Description: m.Description,
			Director:    m.Director,
			Producer:    m.Producer,
			ReleaseDate: m.ReleaseDate,
			Score:       m.Score,
		}
		for _, p := range peoples {
			if p.IsInMovie(m.ID) {
				movie.PeopleIDS = append(movie.PeopleIDS, p.ID)
			}
		}
		d.Movies = append(d.Movies, movie)
	}
	for _, p := range peoples {
		people := People{
			ID:     p.ID,
			Name:   p.Name,
			Gender: p.Gender,
			Age:    p.Age,
		}
		d.Peoples = append(d.Peoples, people)
	}
	return d
}
