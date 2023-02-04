package models

/* TweetMessage - Get message from body */
type TweetMessage struct {
	Message string `bson:"message" json:"message"`
}
