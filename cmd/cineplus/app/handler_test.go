package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymohl-cl/cineplus/pkg/ghibli"
)

func TestJsonDataModel(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		c := &Cache{}
		c.movies = &[]ghibli.Movie{ghibli.Movie{ID: "m1"}, ghibli.Movie{ID: "m2"}}
		c.peoples = &[]ghibli.People{ghibli.People{ID: "p1", FilmsURL: []string{"http://url/m2"}},
			ghibli.People{ID: "p2", FilmsURL: []string{"http://url/m4"}}}

		expectedData := Data{
			Movies:  []Movie{Movie{ID: "m1"}, Movie{ID: "m2", PeopleIDS: []string{"p1"}}},
			Peoples: []People{People{ID: "p1"}, People{ID: "p2"}},
		}

		data := jsonDataModel(c)
		assert.EqualValues(t, expectedData, data)
	})
}
