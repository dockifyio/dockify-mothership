package v1

import (
	"bytes"
	"encoding/json"
	"github.com/dockifyio/dockify-mothership/api/v1/SignUp"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SignUpMock struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BadPayloadUserSignUpMock struct {
	Email    string	`json:"email"`
}

type FireBaseSignUpResponsePayloadMock struct {
	IdToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalId      string `json:"localId"`
}

type FireBaseSignUpBadRequestResponsePayloadMock struct {
	Error string `json:"error"`
}

type RandomPayloadUserSignUpMock struct {
	Random    string	`json:"random"`
}

var signUpUrl = "localhost:8080/v1/signup"

// Test for signup user with bad payload: client error
func TestSignUpUserBadPayload(t *testing.T){
	var fireBaseSignUpResponsePayload FireBaseSignUpResponsePayloadMock
	var fireBaseSignUpBadRequestResponsePayloadMock FireBaseSignUpBadRequestResponsePayloadMock

	userSignUpMockBadPayload := BadPayloadUserSignUpMock{Email: "test@gmail.com"}
	requestBody, err := json.Marshal(userSignUpMockBadPayload)
	assert.Nil(t, err, "Couldn't marshall user Signup mock object:")

	req, err := http.NewRequest("POST", signUpUrl, bytes.NewBuffer(requestBody))
	assert.Nil(t, err, "Failure while making a new POST request:")

	rec := httptest.NewRecorder()
	SignUp.SignUpUser(rec, req)

	//handler.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Expected status bad request as the status code")
	// check the payload response here as well make sure
	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	assert.Nil(t, err, "Error with payload response:")

	err = json.Unmarshal(body, &fireBaseSignUpResponsePayload)
	// expect it to be nil
	assert.Nil(t, err, "Couldn't unmarshall firebase sign up response payload:")
	assert.Equal(t, "", fireBaseSignUpResponsePayload.Email, "Expected return type of email to be empty string")
	assert.Equal(t, "", fireBaseSignUpResponsePayload.IdToken, "Expected return type of IdToken to be empty string")
	assert.Equal(t, "", fireBaseSignUpResponsePayload.RefreshToken, "Expected return type of RefreshToken to be empty string")
	assert.Equal(t, "", fireBaseSignUpResponsePayload.ExpiresIn, "Expected return type of ExpiresIn to be empty string")
	assert.Equal(t, "", fireBaseSignUpResponsePayload.LocalId, "Expected return type of LocalId to be empty string")

	err = json.Unmarshal(body, &fireBaseSignUpBadRequestResponsePayloadMock)
	assert.Equal(t, "Invalid request payload", fireBaseSignUpBadRequestResponsePayloadMock.Error, "Expected invalid request payload with Bad payload on Sign up")
}

// Test for completely bad payload just a random payload
func TestSignUpUserRandomPayload(t *testing.T){
	var fireBaseSignUpResponsePayload FireBaseSignUpResponsePayloadMock
	var fireBaseSignUpBadRequestResponsePayloadMock FireBaseSignUpBadRequestResponsePayloadMock

	userSignUpMockBadPayload := RandomPayloadUserSignUpMock{Random: "test@gmail.com"}
	requestBody, err := json.Marshal(userSignUpMockBadPayload)
	assert.Nil(t, err, "Couldn't marshall user signup mock object:")

	req, err := http.NewRequest("POST", signUpUrl, bytes.NewBuffer(requestBody))
	assert.Nil(t, err, "Failure while making a new POST request:")

	rec := httptest.NewRecorder()
	SignUp.SignUpUser(rec, req)
	//rec := httptest.NewRecorder()

	//handler.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Expected status bad request as the status code")
	// check the payload response here as well make sure
	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	assert.Nil(t, err, "Error with payload response:")

	err = json.Unmarshal(body, &fireBaseSignUpResponsePayload)
	// expect it to be nil
	assert.Nil(t, err, "Couldn't unmarshall signup response payload:")
	assert.Equal(t, "", fireBaseSignUpResponsePayload.Email, "Expected return type of email to be empty string")
	assert.Equal(t, "", fireBaseSignUpResponsePayload.IdToken, "Expected return type of IdToken to be empty string")
	assert.Equal(t, "", fireBaseSignUpResponsePayload.RefreshToken, "Expected return type of RefreshToken to be empty string")
	assert.Equal(t, "", fireBaseSignUpResponsePayload.ExpiresIn, "Expected return type of ExpiresIn to be empty string")
	assert.Equal(t, "", fireBaseSignUpResponsePayload.LocalId, "Expected return type of LocalId to be empty string")

	err = json.Unmarshal(body, &fireBaseSignUpBadRequestResponsePayloadMock)
	assert.Equal(t, "Invalid request payload", fireBaseSignUpBadRequestResponsePayloadMock.Error, "Expected invalid request payload with Bad payload on signup")
}

// Test for signup user with the same credentials
