package bd

import (
	"context"
	"time"

	"github.com/jdferreiro/GoAppSample/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* UpdateUser Update user record in DB */
func UpdateUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	userRecord := make(map[string]interface{})
	if len(u.Name) > 0 {
		userRecord["name"] = u.Name
	}
	if len(u.Avatar) > 0 {
		userRecord["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		userRecord["banner"] = u.Banner
	}
	if len(u.LastName) > 0 {
		userRecord["lastname"] = u.LastName
	}
	if len(u.Location) > 0 {
		userRecord["location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		userRecord["website"] = u.WebSite
	}
	userRecord["borndate"] = u.BornDate

	updString := bson.M{
		"$set": userRecord,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updString)
	if err != nil {
		return false, err
	}

	return true, nil
}
