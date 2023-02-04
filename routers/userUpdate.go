package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jdferreiro/GoAppSample/bd"
	"github.com/jdferreiro/GoAppSample/models"
)

/* UserUpdate - update user data */
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid data received "+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.UpdateUser(t, UserID)
	if !status {
		http.Error(w, "Invalid user data "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusOK)
}
