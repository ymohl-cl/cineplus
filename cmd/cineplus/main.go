package main

import (
	"flag"
	"fmt"

	"github.com/ymohl-cl/cineplus/cmd/cineplus/app"
)

var appName = flag.String("appName", "cineplus", "application name")

func main() {
	app, err := app.New(*appName)
	if err != nil {
		fmt.Printf("Failed to load application: %s", err.Error())
	}
	if err := app.Start(); err != nil {
		fmt.Printf("Error occurred: %s", err.Error())
	}
}
