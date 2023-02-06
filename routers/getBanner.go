package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/jdferreiro/GoAppSample/bd"
)

/* GetBanner - Get banner from database */
func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("userId")
	if len(ID) < 1 {
		http.Error(w, "User id is required", http.StatusBadRequest)
		return
	}

	profile, err := bd.GetProfile(ID)
	if err != nil {
		http.Error(w, "Error getting user profile. "+err.Error(), http.StatusBadRequest)
		return
	}

	bannerFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "banner image not found. ", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, bannerFile)
	if err != nil {
		http.Error(w, "Error copying banner image.", http.StatusBadRequest)
	}
}
