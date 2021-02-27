package FirebaseGateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserLogin struct {
	Email    string
	Password string
}

type FireBaseLoginResponsePayload struct {
	IdToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalId      string `json:"localId"`
}

func (userLoginInfo *UserLogin) LoginWithFirebase(w http.ResponseWriter) error {
	// call Firebase API to login here
	//https: //identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=[API_KEY]
	requestBody, err := json.Marshal(map[string]string{
		"email":             userLoginInfo.Email,
		"password":          userLoginInfo.Password,
		"returnSecureToken": "true",
	})
	if err != nil {
		return err
	}
	fireBaseSignInEndpoint := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=" + "API_TOKEN_HERE"
	// body := strings.NewReader(`fulladdress=22280+S+209th+Way%2C+Queen+Creek%2C+AZ+85142`)
	req, err := http.NewRequest("POST", fireBaseSignInEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// Try to unmarshall the Request endpoint here
	// and if getting some trouble umarshalling the endpoint return err
	// to the client if successful unmarshall it and return some of the payload
	// required as json back to the client
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


