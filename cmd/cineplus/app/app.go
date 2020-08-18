package app

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ymohl-cl/cineplus/pkg/config"
	"github.com/ymohl-cl/cineplus/pkg/ghibli"
)

// App behavior
type App struct {
	driver *echo.Echo
	port   string
	puller *Puller
}

// New application
func New(appName string) (App, error) {
	var c Config
	var err error

	if err = config.ParseEnv(appName, &c); err != nil {
		return App{}, err
	}

	// create cache data
	cache := &Cache{}

	// create handler service
	h := Handler{
		cache: cache,
	}

	// create ghibli client
	var ghibliClient ghibli.Client
	if ghibliClient, err = ghibli.New(appName); err != nil {
		return App{}, errors.New("initialization ghibli client error: " + err.Error())
	}

	// configure app
	a := App{
		driver: echo.New(),
		port:   c.Port,
		puller: &Puller{
			ghibli:    ghibliClient,
			tickTimer: c.RefreshTime,
			cache:     cache,
		},
	}
	a.driver.GET("/ping", h.Ping)
	a.driver.GET("/movies", h.Movies)
	a.driver.Use(middleware.Logger())
	a.driver.Use(middleware.CORS())

	return a, nil
}

// Start the application
func (a App) Start() error {
	defer a.driver.Close()
	defer a.puller.Close()

	a.puller.Start()
	err := a.driver.Start(":" + a.port)
	if err != nil {
		return err
	}
	return nil
}
