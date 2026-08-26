package game

type Faction string

const (
	Faction_Sovereign Faction = "sovereign"
	Faction_Exiled    Faction = "exiled"
)

var raceFactions = map[Race]Faction{
	Race_Human:    Faction_Sovereign,
	Race_Dwarf:    Faction_Sovereign,
	Race_NightElf: Faction_Sovereign,
	Race_Gnome:    Faction_Sovereign,
	Race_Orc:      Faction_Exiled,
	Race_Undead:   Faction_Exiled,
	Race_Tauren:   Faction_Exiled,
	Race_Troll:    Faction_Exiled,
}

// Faction returns the faction a race belongs to. ok is false for an unknown race.
func (r Race) Faction() (faction Faction, ok bool) {
	faction, ok = raceFactions[r]
	return faction, ok
}
