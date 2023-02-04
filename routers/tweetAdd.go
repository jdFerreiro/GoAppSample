package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jdferreiro/GoAppSample/bd"
	"github.com/jdferreiro/GoAppSample/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* TweetAdd - add new tweet in DB */
func TweetAdd(w http.ResponseWriter, r *http.Request) {
	var t models.TweetMessage
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Data received failed "+err.Error(), 400)
		return
	}

	if len(t.Message) == 0 {
		http.Error(w, "Tweet message is required", 400)
		return
	}

	objUserID, _ := primitive.ObjectIDFromHex(UserID)

	dataTweet := models.Tweet{
		UserId:      objUserID,
		Message:     t.Message,
		MessageDate: time.Now().UTC(),
	}

	_, status, err := bd.AddTweet(dataTweet)
	if err != nil {
		http.Error(w, "Error while add tweet"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Could not add tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
