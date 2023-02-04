package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/jdferreiro/GoAppSample/middlew"
	"github.com/jdferreiro/GoAppSample/routers"
)

/* Manejadores set port, handler and listen serve */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/UserRegister", middlew.CheckDB(routers.UserRegister)).Methods("POST")
	router.HandleFunc("/Login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/Profile", middlew.CheckDB(middlew.JwtValidation(routers.Profile))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
