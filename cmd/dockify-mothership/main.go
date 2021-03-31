package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/dockifyio/dockify-mothership/api"
)

func main() {
	// y := mux.NewRouter()
	//x := api.App()
	myFigure := figure.NewFigure("Dockify Mothership", "", true)
	myFigure.Print()
	var app api.App
	app.Initialize()
	app.Run()
}
