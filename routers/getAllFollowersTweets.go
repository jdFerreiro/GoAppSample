package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jdferreiro/GoAppSample/bd"
)

/* GetAllFollowersTweets - Get all user data */
func GetAllFollowersTweets(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Page is required and must be greather than 0", http.StatusBadRequest)
		return
	}

	pageLimit, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Limit is required and must be greather than 0", http.StatusBadRequest)
		return
	}

	pageValue := int64(pageNum)
	limitValue := int64(pageLimit)

	result, status := bd.GetFollowersTweets(userId, pageValue, limitValue)
	if status == false {
		http.Error(w, "GetAllFollowersTweets failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
