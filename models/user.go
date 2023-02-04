package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* User database model for user */
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	LastName string             `bson:"lastName" json:"lastName,omitempty"`
	BornDate time.Time          `bson:"bornDate" json:"bornDate,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Passw    string             `bson:"passw" json:"passw"`
	Avatar   string             `bson:"avatar" json:"avatar,omitempty"`
	Banner   string             `bson:"banner" json:"banner,omitempty"`
	Location string             `bson:"location" json:"location,omitempty"`
	WebSite  string             `bson:"webSite" json:"webSite,omitempty"`
}
