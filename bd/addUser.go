package bd

import (
	"context"
	"time"

	"github.com/jdferreiro/GoAppSample/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* AddUser add record to DB */
func AddUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	u.Passw, _ = EncryptPassw(u.Passw)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
