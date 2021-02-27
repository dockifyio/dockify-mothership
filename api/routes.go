package api

import "github.com/dockifyio/dockify-mothership/api/v1/Login"

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/v1/login", Login.LoginUser).Methods("POST")
}
