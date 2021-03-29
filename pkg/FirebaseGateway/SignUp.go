package FirebaseGateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type SignUp struct {
	Email    string
	Password string
}

type FireBaseSignUpResponsePayload struct {
	IdToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalId      string `json:"localId"`
}

func (userSignUpInfo *SignUp) ValidateSignUpUserInput() error {
	if userSignUpInfo.Email == "" || userSignUpInfo.Password == "" {
		return errors.New("invalid user input")
	}
	return nil
}

func (userSignUpInfo *SignUp) SignUpWithFirebaseEmailAndPassword() (FireBaseSignUpResponsePayload, int,error) {
	// call Firebase API to Sign up here
	var fireBaseSignUpResponsePayload FireBaseSignUpResponsePayload
	requestBody, err := json.Marshal(map[string]string{
		"email":             userSignUpInfo.Email,
		"password":          userSignUpInfo.Password,
		"returnSecureToken": "true",
	})
	if err != nil {
		return fireBaseSignUpResponsePayload, http.StatusInternalServerError, err
	}
	fireBaseSignUpEndpoint := "https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=" + "API_TOKEN_HERE"
	// body := strings.NewReader(`fulladdress=22280+S+209th+Way%2C+Queen+Creek%2C+AZ+85142`)
	req, err := http.NewRequest("POST", fireBaseSignUpEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return fireBaseSignUpResponsePayload, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return fireBaseSignUpResponsePayload, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fireBaseSignUpResponsePayload, http.StatusInternalServerError, err
	}

	err = json.Unmarshal(body, &fireBaseSignUpResponsePayload)
	if err != nil {
		return fireBaseSignUpResponsePayload, http.StatusInternalServerError, err
	}
	return fireBaseSignUpResponsePayload, resp.StatusCode, nil
}
