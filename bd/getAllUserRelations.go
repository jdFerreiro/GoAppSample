package bd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jdferreiro/GoAppSample/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* GetAllUserRelations Get all user relations from DB */
func GeAllUserRelations(ID string) ([]*models.Relation, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	var result []*models.Relation

	filter := bson.M{"userid": ID}

	queryResult, err := col.Find(ctx, filter)
	if err != nil {
		log.Fatal("GetAllUserRelations error. " + err.Error())
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
		fmt.Println(err.Error())
		return result, false, err
	}

	queryResult.Close(ctx)

	return result, true, nil
}
