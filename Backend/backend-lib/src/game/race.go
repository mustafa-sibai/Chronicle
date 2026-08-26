package game

type Race string

const (
	Race_Human    Race = "human"
	Race_Dwarf    Race = "dwarf"
	Race_NightElf Race = "night_elf"
	Race_Gnome    Race = "gnome"
	Race_Orc      Race = "orc"
	Race_Undead   Race = "undead"
	Race_Tauren   Race = "tauren"
	Race_Troll    Race = "troll"
)

func (r Race) Valid() bool {
	switch r {
	case Race_Human, Race_Dwarf, Race_NightElf, Race_Gnome, Race_Orc, Race_Undead, Race_Tauren, Race_Troll:
		return true
	}
	return false
}
