package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jdferreiro/GoAppSample/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* GetAllUsers - Get all users */
func GetAllUsers(ID string, page int64, limit int64, search string, searchType string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.User

	skip := (page - 1) * limit

	queryOptions := options.Find()
	queryOptions.SetSkip(skip)
	queryOptions.SetLimit(limit)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	queryResult, err := col.Find(ctx, query, queryOptions)
	if err != nil {
		fmt.Println("getUsers error. " + err.Error())
		return results, false
	}

	for queryResult.Next(context.TODO()) {
		var record models.User
		err := queryResult.Decode(&record)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserId = ID
		r.UserRelationId = record.ID.Hex()

		addRecord := false
		var status bool

		status, err = GetUserRelation(r)
		if searchType == "new" && status == false {
			addRecord = true
		}
		if searchType == "follow" && status == true {
			addRecord = true
		}
		if r.UserRelationId == ID {
			addRecord = false
		}

		if addRecord {
			record.Passw = ""
			results = append(results, &record)
		}
	}

	err = queryResult.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	queryResult.Close(ctx)

	return results, true

}
