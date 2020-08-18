package puller

import (
	"fmt"
	"time"

	"github.com/ymohl-cl/cineplus/pkg/ghibli"
)

// Puller configuration
type Puller struct {
	GhibliDriver ghibli.Client
	TickTimer    int64
	done         chan bool
}

// Start the puller data
func (p *Puller) Start() {
	var movies []ghibli.Movie
	var peoples []ghibli.People
	var err error

	ticker := time.NewTicker(time.Duration(p.TickTimer) * time.Second)
	p.done = make(chan bool)

	go func() {
		for {
			select {
			case <-p.done:
				return
			case <-ticker.C:
				if movies, err = p.GhibliDriver.Movies(); err != nil {
					fmt.Println("error to get the movies list: ", err.Error())
				}
				fmt.Println("movies: ", movies)

				if peoples, err = p.GhibliDriver.Peoples(); err != nil {
					fmt.Println("error to get the peoples list: ", err.Error())
				}
				fmt.Println("peoples: ", peoples)
			}
		}
	}()
}

// Close stop the puller
func (p *Puller) Close() {
	p.done <- true
}
