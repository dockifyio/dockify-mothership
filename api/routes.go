package api

import (
	"github.com/dockifyio/dockify-mothership/api/v1/Account"
	"github.com/dockifyio/dockify-mothership/api/v1/Login"
	"github.com/dockifyio/dockify-mothership/api/v1/SignUp"
)

func (a *App) initializeRoutes(fireBaseApiKey string) {
	a.Router.Handle("/v1/login", &Login.LoginHandler{FireBaseApiKey: fireBaseApiKey}).Methods("POST")
	a.Router.Handle("/v1/signup", &SignUp.SignUpHandler{FireBaseApiKey: fireBaseApiKey}).Methods("POST")
	a.Router.Handle("/v1/deleteaccount", &Account.DeleteAccountHandler{FireBaseApiKey: fireBaseApiKey}).Methods("POST")
}
