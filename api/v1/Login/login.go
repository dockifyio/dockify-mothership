package Login

import (
	"encoding/json"
	"github.com/dockifyio/dockify-mothership/pkg/FirebaseGateway"
	"github.com/dockifyio/dockify-mothership/pkg/Utilities"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLogin FirebaseGateway.UserLogin
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userLogin); err != nil {
		Utilities.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	err := userLogin.ValidateLoginUserInput()
	if err := decoder.Decode(&userLogin); err != nil {
		Utilities.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	fireBaseLoginResponsePayload, statusCode, err := userLogin.LoginWithFirebase(w)
	if err != nil {
		Utilities.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	Utilities.RespondWithJSON(w, statusCode, fireBaseLoginResponsePayload)
}
