package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* TweetsFollowers */
type TweetsFollowers struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId         string             `bson:"userid" json:"userId,omitempty"`
	UserRelationId string             `bson:"userrelationid" json:"userRelationId,omitempty"`
	Tweet          struct {
		Message     string    `bson:"message" json:"message,omitempty"`
		MessageDate time.Time `bson:"messagedate" json:"messageDate,omitempty"`
		ID          string    `bson:"_id" json:"_id,omitempty"`
	}
}
