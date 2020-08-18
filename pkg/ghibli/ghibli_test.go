package ghibli

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Ping(t *testing.T) {
	t.Run("Should return an error because bad status code is returned", func(t *testing.T) {
		expectedErr := "connection lost with ghibli api"
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()

		c := client{
			driver: &http.Client{},
			url:    ts.URL,
		}
		err := c.Ping()
		if assert.Error(t, err) {
			assert.Equal(t, expectedErr, err.Error())
		}
	})

	t.Run("Should be ok", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()

		c := client{
			driver: &http.Client{},
			url:    ts.URL,
		}
		err := c.Ping()
		assert.NoError(t, err)
	})
}

func TestClient_Movies(t *testing.T) {
	t.Run("Should return an error because bad status code is returned", func(t *testing.T) {
		expectedErr := "request to ghibli client error with status 500 Internal Server Error"
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()

		c := client{
			driver: &http.Client{},
			url:    ts.URL,
		}
		movies, err := c.Movies()
		if assert.Error(t, err) {
			assert.Nil(t, movies)
			assert.Equal(t, expectedErr, err.Error())
		}
	})
	t.Run("Should return an error because the payload can't be parsed", func(t *testing.T) {
		expectedErr := "internal error with the ghibli client"
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`plop`))
		}))
		defer ts.Close()

		c := client{
			driver: &http.Client{},
			url:    ts.URL,
		}
		movies, err := c.Movies()
		if assert.Error(t, err) {
			assert.Nil(t, movies)
			assert.Equal(t, expectedErr, err.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{"id":"id1", "title":"title1"},{"id":"id2", "title":"title2"}]`))
		}))
		defer ts.Close()

		c := client{
			driver: &http.Client{},
			url:    ts.URL,
		}
		movies, err := c.Movies()
		if assert.NoError(t, err) {
			assert.EqualValues(t, []Movie{Movie{ID: "id1", Title: "title1"},
				Movie{ID: "id2", Title: "title2"}}, movies)
		}
	})
}

func TestClient_Peoples(t *testing.T) {
	t.Run("Should return an error because bad status code is returned", func(t *testing.T) {
		expectedErr := "request to ghibli client error with status 500 Internal Server Error"
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()

		c := client{
			driver: &http.Client{},
			url:    ts.URL,
		}
		peoples, err := c.Peoples()
		if assert.Error(t, err) {
			assert.Nil(t, peoples)
			assert.Equal(t, expectedErr, err.Error())
		}
	})
	t.Run("Should return an error because the payload can't be parsed", func(t *testing.T) {
		expectedErr := "internal error with the ghibli client"
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`plop`))
		}))
		defer ts.Close()

		c := client{
			driver: &http.Client{},
			url:    ts.URL,
		}
		peoples, err := c.Peoples()
		if assert.Error(t, err) {
			assert.Nil(t, peoples)
			assert.Equal(t, expectedErr, err.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{"id":"id1", "name":"name1"},{"id":"id2", "name":"name2"}]`))
		}))
		defer ts.Close()

		c := client{
			driver: &http.Client{},
			url:    ts.URL,
		}
		peoples, err := c.Peoples()
		if assert.NoError(t, err) {
			assert.EqualValues(t, []People{People{ID: "id1", Name: "name1"},
				People{ID: "id2", Name: "name2"}}, peoples)
		}
	})
}
