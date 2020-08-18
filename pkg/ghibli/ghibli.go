package ghibli

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Client ghibli to communicate with the api ghibli
type Client interface {
	Ping() error
	Movies() ([]Movie, error)
	Peoples() ([]People, error)
}

type client struct {
	driver *http.Client
	url    string
}

// New instanciate a Client service and return it
func New(appName string) (Client, error) {
	var c client
	var err error
	var conf Config

	if conf, err = NewConfig(appName); err != nil {
		return nil, err
	}
	c.driver = &http.Client{}
	c.url = conf.URL
	if err = c.Ping(); err != nil {
		return nil, err
	}
	return &c, nil
}

func (c client) Ping() error {
	var request *http.Request
	var response *http.Response
	var err error

	if request, err = http.NewRequest(http.MethodGet, c.url, nil); err != nil {
		return err
	}
	if response, err = c.driver.Do(request); err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return errors.New(ErrorGhibliConnection)
	}
	return nil
}

func (c client) Movies() ([]Movie, error) {
	var movies []Movie
	var request *http.Request
	var response *http.Response
	var err error

	if request, err = http.NewRequest(http.MethodGet, c.url+"/films", nil); err != nil {
		return nil, err
	}
	if response, err = c.driver.Do(request); err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("Request to ghibli client return an error: ", response.Status)
		return nil, errors.New(ErrorGhibliRequest)
	}
	if err = json.NewDecoder(response.Body).Decode(&movies); err != nil {
		return nil, errors.New(ErrorInternal)
	}
	return movies, nil
}

func (c client) Peoples() ([]People, error) {
	var peoples []People
	var request *http.Request
	var response *http.Response
	var err error

	if request, err = http.NewRequest(http.MethodGet, c.url+"/people", nil); err != nil {
		return nil, err
	}
	if response, err = c.driver.Do(request); err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("Request to ghibli client return an error: ", response.Status)
		return nil, errors.New(ErrorGhibliRequest)
	}
	if err = json.NewDecoder(response.Body).Decode(&peoples); err != nil {
		return nil, errors.New(ErrorInternal)
	}
	return peoples, nil
}
