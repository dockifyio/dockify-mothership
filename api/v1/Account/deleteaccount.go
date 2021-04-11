package Account

import (
	"encoding/json"
	"github.com/dockifyio/dockify-mothership/pkg/FirebaseGateway"
	"github.com/dockifyio/dockify-mothership/pkg/Utilities"
	"net/http"
)

type DeleteAccountHandler struct {
	FireBaseApiKey string
}

func (deleteAccountHandler *DeleteAccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var deleteFirebaseAccount FirebaseGateway.DeleteFirebaseAccount
	var fireBaseDeleteAccountResponsePayload FirebaseGateway.FireBaseDeleteAccountResponsePayload
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&deleteFirebaseAccount); err != nil {
		Utilities.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()
	err := deleteFirebaseAccount.ValidateDeleteAccountInput()
	if err != nil {
		Utilities.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	statusCode, err := deleteFirebaseAccount.DeleteFirebaseAccount(deleteAccountHandler.FireBaseApiKey)
	if err != nil {
		Utilities.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fireBaseDeleteAccountResponsePayload.Success = "successfully deleted account"
	Utilities.RespondWithJSON(w, statusCode, fireBaseDeleteAccountResponsePayload)
}
