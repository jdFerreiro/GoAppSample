package bd

import (
	"context"
	"log"
	"time"

	"github.com/jdferreiro/GoAppSample/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* GetUserRelations Get user relations from DB */
func GetUserRelations(ID string, pageNumber int64, contentPerPage int64) ([]*models.Relation, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	var result []*models.Relation

	filter := bson.M{"userid": ID}

	skip := (pageNumber - 1) * contentPerPage

	queryOptions := options.Find()
	queryOptions.SetLimit(contentPerPage)
	queryOptions.SetSkip(skip)

	queryResult, err := col.Find(ctx, filter, queryOptions)
	if err != nil {
		log.Fatal("GetUserRelations error. " + err.Error())
		return result, false, err
	}

	if queryResult != nil {
		for queryResult.Next(context.TODO()) {
			var record models.Relation
			err := queryResult.Decode(&record)
			if err != nil {
			}

			result = append(result, &record)
		}
	}

	err = queryResult.Err()
	if err != nil {
		return result, false, err
	}
	queryResult.Close(ctx)

	return result, true, nil
}
