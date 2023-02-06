package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jdferreiro/GoAppSample/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* GetFollowersTweets - Get tweets of user followers */
func GetFollowersTweets(ID string, page int64, limit int64) ([]*models.TweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	var results []*models.TweetsFollowers

	// skip := (page - 1) * limit

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	/*
		conditions = append(conditions, bson.M{"$sort": bson.M{"messagedate": -1}})
		conditions = append(conditions, bson.M{"$skip": skip})
		conditions = append(conditions, bson.M{"$limit": limit})
	*/
	dataResult, err := col.Aggregate(ctx, conditions)
	if err != nil {
		fmt.Println("GetTweetsFollowers error. " + err.Error())
		return results, false
	}
	err = dataResult.All(ctx, &results)
	if err != nil {
		fmt.Println("Procesing data error. " + err.Error())
		return results, false
	}

	dataResult.Close(ctx)

	return results, true
}
