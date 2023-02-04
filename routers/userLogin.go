package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jdferreiro/GoAppSample/bd"
	"github.com/jdferreiro/GoAppSample/jwt"
	"github.com/jdferreiro/GoAppSample/models"
)

/* Login login user */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid Email or Password. "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required. ", 400)
		return
	}

	document, exists := bd.Login(t.Email, t.Passw)
	if exists == false {
		http.Error(w, "Invalid User or Password.", 400)
		return
	}

	jwtKey, err := jwt.CreateJWT(document)
	if err != nil {
		http.Error(w, "Can not create token. "+err.Error(), 400)
		return
	}

	resp := models.LoginResponse{
		JwtToken: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/*
		cokieExp := time.Now().Add(24 * time.Hour)
		http.SetCookie(w, &http.Cookie{
			Name: "GoSampleAppToken"
			Value: jwtKey,
			Expires: cokieExp,
		})
	*/
}
