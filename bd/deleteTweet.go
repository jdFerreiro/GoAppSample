package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* DeleteTweet - Delete a tweet */
func DeleteTweet(TweetId string, UserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	objTweetID, _ := primitive.ObjectIDFromHex(TweetId)
	objUserId, _ := primitive.ObjectIDFromHex(UserId)

	filter := bson.M{
		"_id":    objTweetID,
		"userid": objUserId,
	}

	_, err := col.DeleteOne(ctx, filter)
	return err
}
