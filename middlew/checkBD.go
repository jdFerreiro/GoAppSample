package middlew

import (
	"net/http"

	"github.com/jdferreiro/GoAppSample/bd"
)

/* CheckDB check db Connection */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "DB Connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
