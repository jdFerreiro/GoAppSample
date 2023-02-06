package routers

import (
	"net/http"

	"github.com/jdferreiro/GoAppSample/bd"
	"github.com/jdferreiro/GoAppSample/models"
)

/* DeleteUserRelation - Delete user relation */
func DeleteUserRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID not get", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserId = UserID
	t.UserRelationId = ID

	status, err := bd.DeleteRelation(t)
	if err != nil {
		http.Error(w, "Error while deleting user relation "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Could not delete user relation", 400)
		return
	}

	w.WriteHeader(http.StatusOK)
}
