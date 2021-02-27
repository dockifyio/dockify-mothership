package main

import (
	"github.com/dockifyio/dockify-mothership/api"
)

func main() {
	// y := mux.NewRouter()
	//x := api.App()
	var app api.App
	app.Initialize()
	app.Run()
}
