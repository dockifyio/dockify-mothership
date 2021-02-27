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

	if err := userLogin.LoginWithFirebase(w); err != nil {
		Utilities.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
