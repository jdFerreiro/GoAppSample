package bd

import (
	"context"
	"time"

	"github.com/jdferreiro/GoAppSample/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* AddTweet Add new tweet */
func AddTweet(t models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	record := bson.M{
		"userid":      t.UserId,
		"message":     t.Message,
		"messagedate": t.MessageDate,
	}

	result, err := col.InsertOne(ctx, record)
	if err != nil {
		return string(""), false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
