package api

import (
	"github.com/dockifyio/dockify-mothership/api/v1/Login"
	"github.com/dockifyio/dockify-mothership/api/v1/SignUp"
	"github.com/dockifyio/dockify-mothership/api/v1/Account"
)

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/v1/login", Login.LoginUser).Methods("POST")
	a.Router.HandleFunc("/v1/signup", SignUp.SignUpUser).Methods("POST")
	a.Router.HandleFunc("/v1/deleteaccount", Account.DeleteAccount).Methods("POST")
}
