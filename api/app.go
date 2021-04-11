package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(fireBaseApiKey string) {
	a.Router = mux.NewRouter()

	a.initializeRoutes(fireBaseApiKey)
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
