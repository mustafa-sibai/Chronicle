package game

const (
	MaxSkinColor    = 9
	MaxFaceID       = 9
	MaxHairStyleID  = 9
	MaxHairColor    = 9
	MaxFacialHairID = 9
)

type Appearance struct {
	SkinColor    int `bson:"skinColor" json:"skinColor"`
	FaceID       int `bson:"faceId" json:"faceId"`
	HairStyleID  int `bson:"hairStyleId" json:"hairStyleId"`
	HairColor    int `bson:"hairColor" json:"hairColor"`
	FacialHairID int `bson:"facialHairId" json:"facialHairId"`
}

func (a Appearance) Valid() bool {
	return inRange(a.SkinColor, MaxSkinColor) &&
		inRange(a.FaceID, MaxFaceID) &&
		inRange(a.HairStyleID, MaxHairStyleID) &&
		inRange(a.HairColor, MaxHairColor) &&
		inRange(a.FacialHairID, MaxFacialHairID)
}

func inRange(v, max int) bool {
	return v >= 0 && v <= max
}
