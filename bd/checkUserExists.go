package bd

import (
	"context"
	"time"

	"github.com/jdferreiro/GoAppSample/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* CheckUserExists detect if user email was registerd */
func CheckUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")
	status := true

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		status = false
	}

	return result, status, ID
}
