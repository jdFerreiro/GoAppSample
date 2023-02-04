package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jdferreiro/GoAppSample/bd"
	"github.com/jdferreiro/GoAppSample/models"
)

/* User email */
var UserEmail string

/* User Id */
var UserID string

/* ProcessToken - Extract values from token */
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	privateKey := []byte("904pp54mpl3_CursoGo")
	claims := &models.Claim{}
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	if err == nil {
		_, founded, _ := bd.CheckUserExists(claims.Email)
		if founded == true {
			UserEmail = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, founded, UserID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New(("token inválido"))
	}

	return claims, false, string(""), err
}
