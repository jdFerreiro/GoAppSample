package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jdferreiro/GoAppSample/bd"
)

/* GetAllUsers - Get all user data */
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	searchType := r.URL.Query().Get("searchType")
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	search := r.URL.Query().Get("search")

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

	result, status := bd.GetAllUsers(userId, pageValue, limitValue, search, searchType)
	if status == false {
		http.Error(w, "GetAllUsers failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
