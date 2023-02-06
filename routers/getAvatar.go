package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/jdferreiro/GoAppSample/bd"
)

/* GetAvatar - Get avatar from database */
func GetAvatar(w http.ResponseWriter, r *http.Request) {
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

	avatarFile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		http.Error(w, "Avatar image not found. ", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, avatarFile)
	if err != nil {
		http.Error(w, "Error copying avatar image.", http.StatusBadRequest)
	}
}
