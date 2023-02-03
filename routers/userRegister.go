package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jdferreiro/GoAppSample/bd"
	"github.com/jdferreiro/GoAppSample/models"
)

/* userRegister create user in DB */
func UserRegister(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Data received failed "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email required", 400)
		return
	}

	if len(t.Passw) < 8 {
		http.Error(w, "Password must be grather than eight characters", 400)
		return
	}

	_, userExists, _ := bd.CheckUserExists(t.Email)
	if userExists == true {
		http.Error(w, "User exists", 400)
		return
	}

	_, status, err := bd.AddUser(t)
	if err != nil {
		http.Error(w, "Error while register user"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Could not user register", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
