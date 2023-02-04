package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jdferreiro/GoAppSample/bd"
)

/* GetTweetsByUser - Get tweets by user Id */
func GetTweetsByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if len(userID) < 1 {
		http.Error(w, "User ID not get", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pageNumber")) < 1 {
		http.Error(w, "Did not get page number", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("recordLimit")) < 1 {
		http.Error(w, "Did not get record limit", http.StatusBadRequest)
		return
	}

	pageNumber, err := strconv.Atoi(r.URL.Query().Get("pageNumber"))
	if err != nil {
		http.Error(w, "page number must be grather than 0", http.StatusBadRequest)
		return
	}

	recordLimit, err := strconv.Atoi(r.URL.Query().Get("recordLimit"))
	if err != nil {
		http.Error(w, "Record limit must be grather than 0", http.StatusBadRequest)
		return
	}

	tweets, fail, err := bd.GetTweet(userID, int64(pageNumber), int64(recordLimit))
	if fail {
		http.Error(w, "Falló."+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}
