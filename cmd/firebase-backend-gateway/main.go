package main

import (
	"github.com/dockifyio/firebase-backend-gateway/api/api"
)

func main() {
	// y := mux.NewRouter()
	//x := api.App()
	var app api.App
	app.Initialize()
	app.Run()
}
