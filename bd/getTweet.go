package bd

import (
	"context"
	"log"
	"time"

	"github.com/jdferreiro/GoAppSample/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* GetTweet Get tweet from DB */
func GetTweet(ID string, pageNumber int64, contentPerPage int64) ([]*models.Tweet, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var result []*models.Tweet
	condicion := bson.M{
		"userId": ID,
	}

	queryOptions := options.Find()
	queryOptions.SetLimit(contentPerPage)
	queryOptions.SetSort(bson.D{{Key: "messagedate", Value: -1}})
	queryOptions.SetSkip((pageNumber - 1) * contentPerPage)

	queryResult, err := col.Find(ctx, condicion, queryOptions)
	if err != nil {
		log.Fatal("getTweet error. " + err.Error())
		return result, false, err
	}

	for queryResult.Next(context.TODO()) {
		var record models.Tweet
		err := queryResult.Decode(&record)
		if err != nil {
			return result, false, err
		}

		result = append(result, &record)
	}

	return result, true, nil

}
