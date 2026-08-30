package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Quest struct {
	ID          bson.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string         `bson:"title" json:"title"`
	Description string         `bson:"description" json:"description"`
	Objective   QuestObjective `bson:"objective" json:"objective"`
	Reward      QuestReward    `bson:"reward" json:"reward"`
	CreatedAt   time.Time      `bson:"createdAt" json:"createdAt"`
}

type QuestObjective struct {
	Description     string `bson:"description" json:"description"`
	CompletionCount int    `bson:"completionCount" json:"completionCount"`
}

type QuestReward struct {
	ItemTemplateIDs       []bson.ObjectID `bson:"itemTemplateIds" json:"itemTemplateIds"`
	ChoiceItemTemplateIDs []bson.ObjectID `bson:"choiceItemTemplateIds" json:"choiceItemTemplateIds"`
	Experience            int             `bson:"experience" json:"experience"`
}
