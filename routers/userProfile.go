package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jdferreiro/GoAppSample/bd"
)

/* Profile - Get User Profile */
func Profile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID not get", http.StatusBadRequest)
		return
	}

	profile, err := bd.GetProfile(ID)
	if err != nil {
		http.Error(w, "Error found while getting user profile "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
