package game

type Gender string

const (
	Gender_Male   Gender = "male"
	Gender_Female Gender = "female"
)

func (g Gender) Valid() bool {
	return g == Gender_Male || g == Gender_Female
}
