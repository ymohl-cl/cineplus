package app

import (
	"fmt"
	"time"

	"github.com/ymohl-cl/cineplus/pkg/ghibli"
)

// Puller configuration
type Puller struct {
	ghibli    ghibli.Client
	tickTimer int64
	done      chan bool
	cache     *Cache
}

// Start the puller data
func (p *Puller) Start() {
	p.pull()
	t := time.NewTicker(time.Duration(p.tickTimer) * time.Millisecond)
	p.done = make(chan bool)

	go func() {
		for {
			select {
			case <-p.done:
				return
			case <-t.C:
				p.pull()
			}
		}
	}()
}

func (p *Puller) pull() {
	movies, err := p.ghibli.Movies()
	if err != nil {
		fmt.Println("error to get the movies list: ", err.Error())
		return
	}
	peoples, err := p.ghibli.Peoples()
	if err != nil {
		fmt.Println("error to get the peoples list: ", err.Error())
		return
	}
	p.cache.Update(&movies, &peoples)
}

// Close stop the puller
func (p *Puller) Close() {
	p.done <- true
}
