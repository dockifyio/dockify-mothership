package v1
// so with integration tests we have to be able to create some seed data and than run some
// integration tests using that seed
// with proper unit test we can mock all of this and not actually call the 3rd party libraries, etc
// but still make sure the unit of code does whats needed and required

import (
	"bytes"
	"encoding/json"
	"github.com/dockifyio/dockify-mothership/api/v1/Login"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserLoginMock struct {
	Email    string
	Password string
}

type BadPayloadUserLoginMock struct {
	Email    string
}

type FireBaseLoginResponsePayloadMock struct {
	IdToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalId      string `json:"localId"`
}

var loginUrl = "localhost:8080/v1/login"

// Simple test for logging in user with bad credentials: potentially user error
func TestLoginUserBadCredentials(t *testing.T){
	var fireBaseLoginResponsePayload FireBaseLoginResponsePayloadMock
	userLoginMock := UserLoginMock{Email: "test@gmail.com", Password: "testing123"}
	requestBody, err := json.Marshal(userLoginMock)
	assert.Nil(t, err, "Couldn't marshall user login mock object")
	req, err := http.NewRequest("POST", loginUrl, bytes.NewBuffer(requestBody))
	assert.Nil(t, err, "Failure while making a new POST request:")
	rec := httptest.NewRecorder()
	Login.LoginUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Expected status bad request as the status code")
	// check the payload response here as well make sure
	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err, "Error with payload response")

	err = json.Unmarshal(body, &fireBaseLoginResponsePayload)
	assert.Nil(t, err, "Couldn't unmarshall firebaselogin response payload:" )

	assert.Equal(t, "", fireBaseLoginResponsePayload.Email, "Expected return type of email to be empty string")
	assert.Equal(t, "", fireBaseLoginResponsePayload.IdToken, "Expected return type of IdToken to be empty string")
	assert.Equal(t, "", fireBaseLoginResponsePayload.RefreshToken, "Expected return type of RefreshToken to be empty string")
	assert.Equal(t, "", fireBaseLoginResponsePayload.ExpiresIn, "Expected return type of ExpiresIn to be empty string")
	assert.Equal(t, "", fireBaseLoginResponsePayload.LocalId, "Expected return type of LocalId to be empty string")
}

// Test for login user with bad payload: client error
func TestLoginUserBadPayload(t *testing.T){
	var fireBaseLoginResponsePayload FireBaseLoginResponsePayloadMock
	userLoginMockPadPayload := BadPayloadUserLoginMock{Email: "test@gmail.com"}
	requestBody, err := json.Marshal(userLoginMockPadPayload)
	assert.Nil(t, err, "Couldn't marshall user login mock object:")

	req, err := http.NewRequest("POST", loginUrl, bytes.NewBuffer(requestBody))
	assert.Nil(t, err, "Failure while making a new POST request:")

	rec := httptest.NewRecorder()
	Login.LoginUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Expected status bad request as the status code")
	// check the payload response here as well make sure
	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err, "Error with payload response:")

	err = json.Unmarshal(body, &fireBaseLoginResponsePayload)
	assert.Nil(t, err, "Couldn't unmarshall firebaselogin response payload:")
	assert.Equal(t, "", fireBaseLoginResponsePayload.Email, "Expected return type of email to be empty string")
	assert.Equal(t, "", fireBaseLoginResponsePayload.IdToken, "Expected return type of IdToken to be empty string")
	assert.Equal(t, "", fireBaseLoginResponsePayload.RefreshToken, "Expected return type of RefreshToken to be empty string")
	assert.Equal(t, "", fireBaseLoginResponsePayload.ExpiresIn, "Expected return type of ExpiresIn to be empty string")
	assert.Equal(t, "", fireBaseLoginResponsePayload.LocalId, "Expected return type of LocalId to be empty string")
}

// TODO: Test for login user with correct credentials and payload: happy path