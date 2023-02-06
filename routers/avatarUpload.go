package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/jdferreiro/GoAppSample/bd"
	"github.com/jdferreiro/GoAppSample/models"
)

/* UploadAvatar - Upload Avatar and update database */
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	avatarFile, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileName string = "uploads/avatars/" + UserID + "." + extension

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Upload avatar failed. "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, avatarFile)
	if err != nil {
		http.Error(w, "Image copy failed. "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = UserID + "." + extension

	status, err = bd.UpdateUser(user, UserID)
	if err != nil || status == false {
		http.Error(w, "Error writing avatar to DB", http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)

}
