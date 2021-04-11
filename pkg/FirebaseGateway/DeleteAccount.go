package FirebaseGateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type FireBaseDeleteAccountResponsePayload struct {
	Success string `json:"success"`
}

type DeleteFirebaseAccount struct {
	IDToken string
}

func (deleteAccountInfo *DeleteFirebaseAccount) ValidateDeleteAccountInput() error {
	if deleteAccountInfo.IDToken == "" {
		return errors.New("invalid deleteAccountInfo input")
	}
	return nil
}

func (deleteAccountInfo *DeleteFirebaseAccount) DeleteFirebaseAccount(fireBaseApiKey string) (int,error) {
	// call Firebase API to Delete account here
	requestBody, err := json.Marshal(map[string]string{
		"idToken": deleteAccountInfo.IDToken,
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}
	fireBaseDeleteAccountEndpoint := "https://identitytoolkit.googleapis.com/v1/accounts:delete?key=" + fireBaseApiKey

	req, err := http.NewRequest("POST", fireBaseDeleteAccountEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return resp.StatusCode, nil
}

