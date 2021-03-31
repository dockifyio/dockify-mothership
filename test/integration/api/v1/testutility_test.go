package v1// internal/league/testutility_test.go
import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dockifyio/dockify-mothership/api/v1/Account"
	"github.com/dockifyio/dockify-mothership/api/v1/SignUp"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)


var SignUpUserResponsePayload FireBaseSignUpResponsePayloadMock
var deleteAccountUrl = "localhost:8080/v1/deleteaccount"

type DeleteAccountMockRequest struct {
	IDToken string
}

func signUpUser()  {
	var fireBaseSignUpBadRequestResponsePayloadMock FireBaseSignUpBadRequestResponsePayloadMock

	userSignUpPayload := SignUpMock{Email: "test123@gmail.com", Password: "testing123"}
	requestBody, err := json.Marshal(userSignUpPayload)
	if err != nil {
		fmt.Println("error marshalling userSignUpPayload")
		fmt.Println(err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", signUpUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Failure while making a new POST request:")
		fmt.Println(err)
		os.Exit(1)
	}

	rec := httptest.NewRecorder()
	SignUp.SignUpUser(rec, req)

	//handler.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Expected status good request as the status code 200")
		fmt.Printf("Response recieved was %d\n", res.StatusCode)
		fmt.Println(err)
		os.Exit(1)
	}

	// check the payload response here as well make sure
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Expected no error from res body reader in signup user")
		fmt.Println(err)
		os.Exit(1)
	}


	err = json.Unmarshal(body, &SignUpUserResponsePayload)
	if err != nil {
		fmt.Println("Expected no error when unmarshalling payload from signup endpoint")
		fmt.Println(err)
		os.Exit(1)
	}
	// expect it to be nil
	if SignUpUserResponsePayload.Email != "test123@gmail.com" {
		fmt.Println("Expected Email for signup to be the same")
		fmt.Println(err)
		os.Exit(1)
	}

	if SignUpUserResponsePayload.IdToken == "" {
		fmt.Println("Expected IdToken for signup to be the not empty")
		fmt.Println(err)
		os.Exit(1)
	}

	if SignUpUserResponsePayload.RefreshToken == "" {
		fmt.Println("Expected RefreshToken for signup to be the not empty")
		fmt.Println(err)
		os.Exit(1)
	}

	if SignUpUserResponsePayload.ExpiresIn == "" {
		fmt.Println("Expected ExpiresIn for signup to be the not empty")
		fmt.Println(err)
		os.Exit(1)
	}

	if SignUpUserResponsePayload.LocalId == "" {
		fmt.Println("Expected LocalId for signup to be the not empty")
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(body, &fireBaseSignUpBadRequestResponsePayloadMock)
	if fireBaseSignUpBadRequestResponsePayloadMock.Error != "" {
		fmt.Println("Expected no Error for bad request pay load")
		fmt.Println(err)
		os.Exit(1)
	}
}

func deleteAccount() {
	deleteAccountPayload := DeleteAccountMockRequest{IDToken: SignUpUserResponsePayload.IdToken}
	requestBody, err := json.Marshal(deleteAccountPayload)
	if err != nil {
		fmt.Println("error marshalling deleteAccountPayload")
		fmt.Println(err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", deleteAccountUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Failure while making a new POST request:")
		fmt.Println(err)
		os.Exit(1)
	}

	rec := httptest.NewRecorder()
	Account.DeleteAccount(rec, req)

	//handler.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Expected status good request for delete account status code 200")
		fmt.Printf("Response recieved was %d\n", res.StatusCode)
		fmt.Println(err)
		os.Exit(1)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	// Do something here.
	fmt.Printf("\033[1;36m%s\033[0m", "> Setting up test conditions\n")
	fmt.Printf("\033[1;36m%s\033[0m", "> Signing up a new user\n")
	signUpUser()
	fmt.Printf("\033[1;36m%s\033[0m", "> Test user was created sucessfully for use	\n")
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	// Do something here.
	fmt.Printf("\033[1;36m%s\033[0m", "> Tearing down test conditions\n")
	fmt.Printf("\033[1;36m%s\033[0m", "> Tearing down test account	\n")
	deleteAccount()
	fmt.Printf("\033[1;36m%s\033[0m", "> Successfully tore down test account	\n")
	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed\n")
	fmt.Printf("\n")
}
