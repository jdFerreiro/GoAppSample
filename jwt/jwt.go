package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jdferreiro/GoAppSample/models"
)

/* CreateJWT generate a new json web token */
func CreateJWT(t models.User) (string, error) {

	privateKey := []byte("904pp54mpl3_CursoGo")
	payload := jwt.MapClaims{
		"email":    t.Email,
		"name":     t.Name,
		"lastName": t.LastName,
		"location": t.Location,
		"webSite":  t.WebSite,
		"_id":      t.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(privateKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
