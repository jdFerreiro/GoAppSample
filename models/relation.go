package models

/* Relation Structure */
type Relation struct {
	UserId         string `bson:"userid" json:"userId"`
	UserRelationId string `bson:"userrelationid" json:"userRelationId"`
}
