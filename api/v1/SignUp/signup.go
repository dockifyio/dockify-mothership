package SignUp

import (
	"encoding/json"
	"github.com/dockifyio/dockify-mothership/pkg/FirebaseGateway"
	"github.com/dockifyio/dockify-mothership/pkg/Utilities"
	"net/http"
)

type SignUpHandler struct {
	FireBaseApiKey string
}

func (signUpHandler *SignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var userSignUp FirebaseGateway.SignUp
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userSignUp); err != nil {
		Utilities.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := userSignUp.ValidateSignUpUserInput()
	if err != nil {
		Utilities.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	fireBaseLoginResponsePayload, statusCode, err := userSignUp.SignUpWithFirebaseEmailAndPassword(signUpHandler.FireBaseApiKey)
	if err != nil {
		Utilities.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	Utilities.RespondWithJSON(w, statusCode, fireBaseLoginResponsePayload)

}
