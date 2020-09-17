package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

type UserLogin struct {
	Email    string
	Password string
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func (userLoginInfo *UserLogin) loginWithFirebase() error {
	// call firebase API to login here
	return nil
}

func (a *App) loginUser(w http.ResponseWriter, r *http.Request) {
	var userLogin UserLogin
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userLogin); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := userLogin.loginWithFirebase(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/login", a.loginUser).Methods("POST")
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
