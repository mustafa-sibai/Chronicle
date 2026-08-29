package game

type Stats struct {
	Strength  int `bson:"strength" json:"strength"`
	Agility   int `bson:"agility" json:"agility"`
	Stamina   int `bson:"stamina" json:"stamina"`
	Intellect int `bson:"intellect" json:"intellect"`
	Spirit    int `bson:"spirit" json:"spirit"`
	Health    int `bson:"health" json:"health"`
	Mana      int `bson:"mana" json:"mana"`
}
