package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jdferreiro/GoAppSample/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* GetUserRelation Get user relations from DB */
func GetUserRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	var result models.Relation

	filter := bson.M{
		"userid":         t.UserId,
		"userrelationid": t.UserRelationId,
	}

	fmt.Println(filter)

	err := col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println("GetUserRelations error. " + err.Error())
		return false, err
	}

	return true, nil
}
