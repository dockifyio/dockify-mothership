package v1
// so with integration tests we have to be able to create some seed data and than run some
// integration tests using that seed
// with proper unit test we can mock all of this and not actually call the 3rd party libraries, etc
// but still make sure the unit of code does whats needed and required

//import (
//	"bytes"
//	"encoding/json"
//	"net/http"
//	"testing"
//)
//
//type UserLoginMock struct {
//	Email    string
//	Password string
//}
//
//var loingUrl string = "locahost:8080/v1/login"
//// Simple test for logging in user
//func LoginUserTest(t *testing.T){
//	userLoginMock := UserLoginMock{Email: "test@gmail.com", Password: "testing123"}
//	requestBody, err := json.Marshal(userLoginMock)
//	if err != nil {
//
//	}
//	req, err := http.NewRequest("POST", loingUrl, bytes.NewBuffer(requestBody))
//}
