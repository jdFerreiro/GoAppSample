package bd

import (
	"github.com/jdferreiro/GoAppSample/models"
	"golang.org/x/crypto/bcrypt"
)

/* Login function */
func Login(email string, password string) (models.User, bool) {
	user, finded, _ := CheckUserExists(email)
	if finded == false {
		return user, false
	}

	pPasw := []byte(password)
	dbPassw := []byte(user.Passw)
	err := bcrypt.CompareHashAndPassword(dbPassw, pPasw)
	if err != nil {
		return user, false
	}

	return user, true
}
