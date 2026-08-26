package game

type Class string

const (
	Class_Warrior Class = "warrior"
	Class_Paladin Class = "paladin"
	Class_Hunter  Class = "hunter"
	Class_Rogue   Class = "rogue"
	Class_Priest  Class = "priest"
	Class_Shaman  Class = "shaman"
	Class_Mage    Class = "mage"
	Class_Warlock Class = "warlock"
	Class_Druid   Class = "druid"
)

var raceClasses = map[Race][]Class{
	Race_Human:    {Class_Warrior, Class_Paladin, Class_Rogue, Class_Priest, Class_Mage, Class_Warlock},
	Race_Dwarf:    {Class_Warrior, Class_Paladin, Class_Hunter, Class_Rogue, Class_Priest},
	Race_NightElf: {Class_Warrior, Class_Hunter, Class_Rogue, Class_Priest, Class_Druid},
	Race_Gnome:    {Class_Warrior, Class_Rogue, Class_Mage, Class_Warlock},
	Race_Orc:      {Class_Warrior, Class_Hunter, Class_Rogue, Class_Shaman, Class_Warlock},
	Race_Undead:   {Class_Warrior, Class_Rogue, Class_Priest, Class_Mage, Class_Warlock},
	Race_Tauren:   {Class_Warrior, Class_Hunter, Class_Shaman, Class_Druid},
	Race_Troll:    {Class_Warrior, Class_Hunter, Class_Rogue, Class_Priest, Class_Shaman, Class_Mage},
}

func (r Race) IsClassAllowedInRace(c Class) bool {
	for _, class := range raceClasses[r] {
		if class == c {
			return true
		}
	}
	return false
}
