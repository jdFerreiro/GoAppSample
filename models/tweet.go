package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweet struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserId      primitive.ObjectID `bson:"userid" json:"userid,omitempty"`
	Message     string             `bson:"message" json:"message"`
	MessageDate time.Time          `bson:"messagedate" json:"messagedate,omitempty"`
}
