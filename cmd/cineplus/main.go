package main

import (
	"flag"
	"fmt"

	"github.com/ymohl-cl/cineplus/cmd/cineplus/puller"
	"github.com/ymohl-cl/cineplus/pkg/ghibli"
)

var appName = flag.String("appName", "cineplus", "application name")

func main() {
	var err error
	var conf Config
	var ghibliDriver ghibli.Client

	if conf, err = NewConfig(*appName); err != nil {
		fmt.Println("Can't load application parameters env: ", err.Error())
	}
	if ghibliDriver, err = ghibli.New(*appName); err != nil {
		fmt.Println("error to initialize the ghibli client: ", err.Error())
		return
	}
	pullerDriver := puller.Puller{
		GhibliDriver: ghibliDriver,
		TickTimer:    conf.RefreshTime,
	}
	defer pullerDriver.Close()
	go pullerDriver.Start()
	fmt.Println("All is done !")
}
