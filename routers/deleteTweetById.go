package routers

import (
	"net/http"

	"github.com/jdferreiro/GoAppSample/bd"
)

/* DeleteTweetById - Delete Tweet by Tweet Id */
func DeleteTweetById(w http.ResponseWriter, r *http.Request) {
	tweetID := r.URL.Query().Get("tweetId")
	if len(tweetID) < 1 {
		http.Error(w, "Tweet ID not get", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(tweetID, UserID)
	if err != nil {
		http.Error(w, "Error deleting tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
