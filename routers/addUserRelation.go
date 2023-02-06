package routers

import (
	"net/http"

	"github.com/jdferreiro/GoAppSample/bd"
	"github.com/jdferreiro/GoAppSample/models"
)

/* AddUserRelation - Add new user relation */
func AddUserRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID not get", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserId = UserID
	t.UserRelationId = ID

	status, err := bd.AddRelation(t)
	if err != nil {
		http.Error(w, "Error while register user relation "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Could not add user relation", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
