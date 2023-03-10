package bd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jdferreiro/GoAppSample/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* GetTweet Get tweet from DB */
func GetTweet(ID string, pageNumber int64, contentPerPage int64) ([]*models.Tweet, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var result []*models.Tweet

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"userid": objID}

	skip := (pageNumber - 1) * contentPerPage

	queryOptions := options.Find()
	queryOptions.SetSkip(skip)
	queryOptions.SetLimit(contentPerPage)
	queryOptions.SetSort(bson.D{{Key: "messagedate", Value: -1}})

	queryResult, err := col.Find(ctx, filter, queryOptions)
	if err != nil {
		log.Fatal("getTweet error. " + err.Error())
		return result, false, err
	}

	if queryResult != nil {
		for queryResult.Next(context.TODO()) {
			var record models.Tweet
			err := queryResult.Decode(&record)
			if err != nil {
			}

			result = append(result, &record)
		}
	}

	err = queryResult.Err()
	if err != nil {
		fmt.Println(err.Error())
		return result, false, err
	}

	queryResult.Close(ctx)

	return result, true, nil
}
