package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* User database model for user */
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	LastName string             `bson:"lastname" json:"lastname,omitempty"`
	BornDate time.Time          `bson:"borndate" json:"borndate,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Passw    string             `bson:"passw" json:"passw,omitempty"`
	Avatar   string             `bson:"avatar" json:"avatar,omitempty"`
	Banner   string             `bson:"banner" json:"banner,omitempty"`
	Location string             `bson:"location" json:"location,omitempty"`
	WebSite  string             `bson:"website" json:"website,omitempty"`
}
