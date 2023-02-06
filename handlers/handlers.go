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

	router.HandleFunc("/userRegister", middlew.CheckDB(routers.UserRegister)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.CheckDB(middlew.JwtValidation(routers.Profile))).Methods("GET")
	router.HandleFunc("/userUpdate", middlew.CheckDB(middlew.JwtValidation(routers.UserUpdate))).Methods("PUT")
	router.HandleFunc("/addTweet", middlew.CheckDB(middlew.JwtValidation(routers.TweetAdd))).Methods("POST")
	router.HandleFunc("/getTweetByUser", middlew.CheckDB(middlew.JwtValidation(routers.GetTweetsByUser))).Methods("GET")
	router.HandleFunc("/deleteTweetById", middlew.CheckDB(middlew.JwtValidation(routers.DeleteTweetById))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlew.CheckDB(middlew.JwtValidation(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/uploadBanner", middlew.CheckDB(middlew.JwtValidation(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckDB(middlew.JwtValidation(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/getBanner", middlew.CheckDB(middlew.JwtValidation(routers.GetBanner))).Methods("GET")

	router.HandleFunc("/addUserRelation", middlew.CheckDB(routers.AddUserRelation)).Methods("POST")
	router.HandleFunc("/deleteUserRelation", middlew.CheckDB(routers.DeleteUserRelation)).Methods("DELETE")
	router.HandleFunc("/getRelationsByUser", middlew.CheckDB(middlew.JwtValidation(routers.GetRelationsByUser))).Methods("GET")

	router.HandleFunc("/getAllUsers", middlew.CheckDB(middlew.JwtValidation(routers.GetAllUsers))).Methods("GET")
	router.HandleFunc("/getFollowersTweets", middlew.CheckDB(middlew.JwtValidation(routers.GetAllFollowersTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
