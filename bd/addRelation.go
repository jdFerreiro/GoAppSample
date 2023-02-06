package bd

import (
	"context"
	"time"

	"github.com/jdferreiro/GoAppSample/models"
)

/* AddRelation - Add new user relation to database */
func AddRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
