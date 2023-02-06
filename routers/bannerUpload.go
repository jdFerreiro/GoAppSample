package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/jdferreiro/GoAppSample/bd"
	"github.com/jdferreiro/GoAppSample/models"
)

/* UploadBanner - Upload banner and update database */
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	bannerFile, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileName string = "uploads/banners/" + UserID + "." + extension

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Upload banner failed. "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, bannerFile)
	if err != nil {
		http.Error(w, "Image copy failed. "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = UserID + "." + extension

	status, err = bd.UpdateUser(user, UserID)
	if err != nil || status == false {
		http.Error(w, "Error writing banner to DB", http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)

}
