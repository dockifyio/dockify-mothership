package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type FireBaseLoginPayload struct {
	idToken      string `json:"idToken"`
	email        string `json:"email"`
	refreshToken string `json:"refreshToken"`
	expiresIn    string `json:"expiresIn"`
	localId      string `json:"localId"`
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

func (userLoginInfo *UserLogin) loginWithFirebase(w http.ResponseWriter) error {
	// call Firebase API to login here
	//https: //identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=[API_KEY]
	requestBody, err := json.Marshal(map[string]string{
		"email":             userLoginInfo.Email,
		"password":          userLoginInfo.Password,
		"returnSecureToken": "true",
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Error")
	}
	fireBaseSignInEndpoint := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=" + "API_TOKEN_HERE"
	// body := strings.NewReader(`fulladdress=22280+S+209th+Way%2C+Queen+Creek%2C+AZ+85142`)
	req, err := http.NewRequest("POST", fireBaseSignInEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		// handle err
		fmt.Println("ERRRRRRRRR")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		// handle err
		fmt.Println("yoooo")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle err
		fmt.Println("dodookei")
	}
	fmt.Println(string(body))
	// resp, err := http.Post(fireBaseSignInEndpoint, "application/json", bytes.NewBuffer(requestBody))
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, "Internal Error calling Firebase Platform")
	// }

	// if resp.StatusCode != 200 {
	// 	respondWithError(w, resp.StatusCode, "Couldn't fufill request")
	// }

	// defer resp.Body.Close()

	// var fireBaseLoginPayload FireBaseLoginPayload

	// body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, "Internal Error from parsing firebase platform payload")
	// }

	// err = json.Unmarshal(body, &fireBaseLoginPayload)

	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, "Internal Error couldn't unmarshall firebase login payload")
	// }
	// log.Println(string(body))

	// respondWithJSON(w, http.StatusOK, fireBaseLoginPayload)
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

	if err := userLogin.loginWithFirebase(w); err != nil {
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

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
