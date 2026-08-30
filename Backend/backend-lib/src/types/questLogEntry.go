package types

import "go.mongodb.org/mongo-driver/v2/bson"

type QuestLogEntry struct {
	QuestID  bson.ObjectID `bson:"questId" json:"questId"`
	Progress int           `bson:"progress" json:"progress"`
}
