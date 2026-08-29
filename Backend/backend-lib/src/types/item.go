package types

import "go.mongodb.org/mongo-driver/v2/bson"

type Item struct {
	ID            bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TemplateID    bson.ObjectID `bson:"templateId" json:"templateId"`
	CurrentStacks int           `bson:"currentStacks" json:"currentStacks"`
}

type Bag struct {
	Item     `bson:",inline"`
	Contents []Item `bson:"contents" json:"contents"`
}
