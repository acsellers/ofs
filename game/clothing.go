package game

type Clothing int

type ClothingStat struct {
	ID   ItemID
	Name string
	Tier int
	Sex

	Stats SPECIAL
}

type Sex int

const (
	Both Sex = iota
	MaleOnly
	FemaleOnly
)

func init() {
	for k, c := range Clothes {
		c.ID = ItemID(k)
		Clothes[k] = c
	}
}

var Clothes = map[Clothing]ClothingStat{
	// baseline
	0: {
		Name: "Vault Suit",
		Tier: 0,
	},

	// vault suit
	1: {
		Name: "Armored Vault Uniform",
		Tier: 1,
		Stats: SPECIAL{
			Perception: 3,
		},
	},
	2: {
		Name: "Sturdy Vault Suit",
		Tier: 2,
		Stats: SPECIAL{
			Perception: 5,
		},
	},
	3: {
		Name: "Heavy Vault Suit",
		Tier: 3,
		Stats: SPECIAL{
			Perception: 7,
		},
	},

	// battle armor
	4: {
		Name: "Battle Armor",
		Tier: 1,
		Stats: SPECIAL{
			Strength:  2,
			Endurance: 1,
		},
	},
	5: {
		Name: "Sturdy Battle Armor",
		Tier: 2,
		Stats: SPECIAL{
			Strength:  3,
			Endurance: 2,
		},
	},
	6: {
		Name: "Heavy Battle Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:  4,
			Endurance: 3,
		},
	},

	// combat armor
	7: {
		Name: "Combat Armor",
		Tier: 1,
		Stats: SPECIAL{
			Strength: 2,
			Agility:  1,
		},
	},
	8: {
		Name: "Sturdy Combat Armor",
		Tier: 2,
		Stats: SPECIAL{
			Strength: 3,
			Agility:  2,
		},
	},
	9: {
		Name: "Heavy Combat Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength: 4,
			Agility:  3,
		},
	},

	// formal wear
	10: {
		Name: "Formal Wear",
		Tier: 1,
		Stats: SPECIAL{
			Luck: 3,
		},
	},
	11: {
		Name: "Fancy Formal Wear",
		Tier: 2,
		Stats: SPECIAL{
			Luck: 5,
		},
	},
	12: {
		Name: "Lucky Formal Wear",
		Tier: 3,
		Stats: SPECIAL{
			Luck: 7,
		},
	},

	// jumpsuit
	13: {
		Name: "Handyman Jumpsuit",
		Tier: 1,
		Stats: SPECIAL{
			Agility: 3,
		},
	},
	14: {
		Name: "Advanced Jumpsuit",
		Tier: 2,
		Stats: SPECIAL{
			Agility: 5,
		},
	},
	15: {
		Name: "Expert Jumpsuit",
		Tier: 3,
		Stats: SPECIAL{
			Agility: 7,
		},
	},

	// lab coat
	16: {
		Name: "Lab Coat",
		Tier: 1,
		Stats: SPECIAL{
			Intelligence: 3,
		},
	},
	17: {
		Name: "Advanced Lab Coat",
		Tier: 2,
		Stats: SPECIAL{
			Intelligence: 5,
		},
	},
	18: {
		Name: "Expert Lab Coat",
		Tier: 3,
		Stats: SPECIAL{
			Intelligence: 7,
		},
	},

	// leather armor
	19: {
		Name: "Leather Armor",
		Tier: 1,
		Stats: SPECIAL{
			Strength:  1,
			Endurance: 2,
		},
	},
	20: {
		Name: "Sturdy Leather Armor",
		Tier: 2,
		Stats: SPECIAL{
			Strength:  2,
			Endurance: 3,
		},
	},
	21: {
		Name: "Heavy Leather Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:  3,
			Endurance: 4,
		},
	},

	// merc gear
	22: {
		Name: "Merc Gear",
		Tier: 1,
		Stats: SPECIAL{
			Perception: 1,
			Agility:    1,
			Luck:       1,
		},
	},
	23: {
		Name: "Sturdy Merc Gear",
		Tier: 2,
		Stats: SPECIAL{
			Perception: 1,
			Agility:    2,
			Luck:       2,
		},
	},
	24: {
		Name: "Heavy Merc Gear",
		Tier: 3,
		Stats: SPECIAL{
			Perception: 2,
			Agility:    3,
			Luck:       2,
		},
	},

	// military fatigues
	25: {
		Name: "Military Fatigues",
		Tier: 1,
		Stats: SPECIAL{
			Strength: 3,
		},
	},
	26: {
		Name: "Officer Fatigues",
		Tier: 2,
		Stats: SPECIAL{
			Strength: 5,
		},
	},
	27: {
		Name: "Commander Fatigues",
		Tier: 3,
		Stats: SPECIAL{
			Strength: 7,
		},
	},

	// nightwear
	28: {
		Name: "Nightwear",
		Tier: 1,
		Stats: SPECIAL{
			Luck: 3,
		},
	},
	29: {
		Name: "Naughty Nightwear",
		Tier: 2,
		Stats: SPECIAL{
			Luck: 5,
		},
	},
	30: {
		Name: "Lucky Nightwear",
		Tier: 3,
		Stats: SPECIAL{
			Luck: 7,
		},
	},

	// officer uniform
	31: {
		Name: "Junior Officer Uniform",
		Tier: 1,
		Stats: SPECIAL{
			Charisma:     1,
			Intelligence: 2,
		},
	},
	32: {
		Name: "Officer Uniform",
		Tier: 2,
		Stats: SPECIAL{
			Charisma:     2,
			Intelligence: 3,
		},
	},
	33: {
		Name: "Commander Uniform",
		Tier: 3,
		Stats: SPECIAL{
			Charisma:     3,
			Intelligence: 4,
		},
	},

	// radiation suit
	34: {
		Name: "Radiation Suit",
		Tier: 1,
		Stats: SPECIAL{
			Perception: 1,
			Endurance:  2,
		},
	},
	35: {
		Name: "Advanced Radiation Suit",
		Tier: 2,
		Stats: SPECIAL{
			Perception: 2,
			Endurance:  3,
		},
	},
	36: {
		Name: "Expert Radiation Suit",
		Tier: 3,
		Stats: SPECIAL{
			Perception: 3,
			Endurance:  4,
		},
	},

	// raider armor
	37: {
		Name: "Raider Armor",
		Tier: 1,
		Stats: SPECIAL{
			Perception: 1,
			Agility:    2,
		},
	},
	38: {
		Name: "Sturdy Raider Armor",
		Tier: 2,
		Stats: SPECIAL{
			Perception: 2,
			Agility:    3,
		},
	},
	39: {
		Name: "Heavy Raider Armor",
		Tier: 3,
		Stats: SPECIAL{
			Perception: 3,
			Agility:    4,
		},
	},

	// robes
	40: {
		Name: "Initiate Robe",
		Tier: 1,
		Stats: SPECIAL{
			Charisma: 2,
			Agility:  1,
		},
	},
	41: {
		Name: "Scribe Robe",
		Tier: 2,
		Stats: SPECIAL{
			Charisma: 3,
			Agility:  2,
		},
	},
	42: {
		Name: "Scribe Rothchild's Robe",
		Tier: 3,
		Stats: SPECIAL{
			Perception:   2,
			Endurance:    1,
			Charisma:     2,
			Intelligence: 2,
		},
	},

	// wasteland gear
	43: {
		Name: "Wasteland Gear",
		Tier: 1,
		Stats: SPECIAL{
			Endurance: 3,
		},
	},
	44: {
		Name: "Sturdy Wasteland Gear",
		Tier: 2,
		Stats: SPECIAL{
			Endurance: 5,
		},
	},
	45: {
		Name: "Heavy Wasteland Gear",
		Tier: 3,
		Stats: SPECIAL{
			Endurance: 7,
		},
	},

	// wasteland medic
	46: {
		Name: "Wasteland Medic",
		Tier: 1,
		Stats: SPECIAL{
			Perception: 2,
			Luck:       1,
		},
	},
	47: {
		Name: "Wasteland Doctor",
		Tier: 2,
		Stats: SPECIAL{
			Perception: 3,
			Luck:       2,
		},
	},
	48: {
		Name: "Wasteland Surgeon",
		Tier: 3,
		Stats: SPECIAL{
			Perception: 4,
			Luck:       3,
		},
	},

	// rare outfits
	49: {
		Name: "Clergy Outfit",
		Tier: 2,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Charisma: 4,
			Luck:     1,
		},
	},
	50: {
		Name: "Comedian Outfit",
		Tier: 2,
		Stats: SPECIAL{
			Perception: 2,
			Charisma:   2,
			Luck:       1,
		},
	},
	51: {
		Name: "Engineer Outfit",
		Tier: 2,
		Stats: SPECIAL{
			Endurance:    2,
			Intelligence: 2,
			Luck:         1,
		},
	},
	52: {
		Name: "Greaser Outfit",
		Tier: 2,
		Stats: SPECIAL{
			Charisma: 2,
			Agility:  2,
			Luck:     1,
		},
	},
	53: {
		Name: "Horror Fan Outfit",
		Tier: 2,
		Stats: SPECIAL{
			Endurance: 4,
			Luck:      1,
		},
	},
	54: {
		Name: "Knight Armor",
		Tier: 2,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Strength:   2,
			Perception: 2,
			Luck:       1,
		},
	},
	55: {
		Name: "Librarian Outfit",
		Tier: 2,
		Sex:  FemaleOnly,
		Stats: SPECIAL{
			Intelligence: 4,
			Luck:         1,
		},
	},
	56: {
		Name: "Mayor Outfit",
		Tier: 2,
		Stats: SPECIAL{
			Charisma:     2,
			Intelligence: 2,
			Luck:         1,
		},
	},
	57: {
		Name: "Medieval Ruler Outfit",
		Tier: 2,
		Stats: SPECIAL{
			Perception: 2,
			Charisma:   2,
			Luck:       1,
		},
	},
	58: {
		Name: "Movie Fan Outfit",
		Tier: 2,
		Sex:  FemaleOnly,
		Stats: SPECIAL{
			Perception: 4,
			Luck:       1,
		},
	},
	59: {
		Name: "Ninja Outfit",
		Tier: 2,
		Stats: SPECIAL{
			Agility: 4,
			Luck:    1,
		},
	},
	60: {
		Name: "Nobility Outfit",
		Tier: 2,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Endurance:    2,
			Intelligence: 2,
			Luck:         1,
		},
	},
	61: {
		Name: "Professor Outfit",
		Tier: 2,
		Stats: SPECIAL{
			Intelligence: 4,
			Luck:         1,
		},
	},
	62: {
		Name: "Republic Robes",
		Tier: 2,
		Sex:  FemaleOnly,
		Stats: SPECIAL{
			Charisma: 4,
			Luck:     1,
		},
	},
	63: {
		Name: "Sci-fi Fan Outfit",
		Tier: 2,
		Stats: SPECIAL{
			Intelligence: 2,
			Agility:      2,
			Luck:         1,
		},
	},
	64: {
		Name: "Soldier Uniform",
		Tier: 2,
		Stats: SPECIAL{
			Strength:  2,
			Endurance: 2,
			Luck:      1,
		},
	},
	65: {
		Name: "Sports Fan Outfit",
		Tier: 2,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Strength: 4,
			Luck:     1,
		},
	},
	66: {
		Name: "Surgeon Outfit",
		Tier: 2,
		Stats: SPECIAL{
			Perception: 2,
			Agility:    2,
			Luck:       1,
		},
	},
	67: {
		Name: "Survivor Armor",
		Tier: 2,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Strength: 2,
			Agility:  2,
			Luck:     1,
		},
	},
	68: {
		Name: "Wrestler Outfit",
		Tier: 2,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Strength:  2,
			Endurance: 2,
			Luck:      1,
		},
	},

	// legendary outfits
	69: {
		Name: "Autumn's Uniform",
		Tier: 3,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Strength:   2,
			Perception: 2,
			Endurance:  2,
			Charisma:   1,
		},
	},
	70: {
		Name: "Bittercup's Outfit",
		Tier: 3,
		Sex:  FemaleOnly,
		Stats: SPECIAL{
			Strength:   2,
			Perception: 2,
			Endurance:  2,
			Charisma:   1,
		},
	},
	71: {
		Name: "Confessor Cromwell's Rags",
		Tier: 3,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Perception:   2,
			Endurance:    2,
			Intelligence: 1,
			Luck:         2,
		},
	},
	72: {
		Name: "Elder Robe",
		Tier: 3,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Charisma: 4,
			Agility:  3,
		},
	},
	73: {
		Name: "Eulogy Jones' Suit",
		Tier: 3,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Perception:   2,
			Charisma:     2,
			Intelligence: 1,
			Luck:         2,
		},
	},
	74: {
		Name: "Minuteman Uniform",
		Tier: 3,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Strength:     2,
			Perception:   2,
			Intelligence: 2,
			Agility:      2,
		},
	},
	75: {
		Name: "Scribe Rothchild's Robe",
		Tier: 3,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Perception:   2,
			Endurance:    1,
			Charisma:     2,
			Intelligence: 2,
		},
	},
	76: {
		Name: "Sheriff's Duster",
		Tier: 3,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Perception: 2,
			Endurance:  5,
		},
	},
	77: {
		Name: "Tenpenny's Suit",
		Tier: 3,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Perception:   2,
			Charisma:     2,
			Intelligence: 2,
			Luck:         1,
		},
	},
	78: {
		Name: "Three Dog's Outfit",
		Tier: 3,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Perception: 2,
			Charisma:   5,
		},
	},
	79: {
		Name: "Tunnel Snakes' Outfit",
		Tier: 3,
		Sex:  MaleOnly,
		Stats: SPECIAL{
			Perception: 2,
			Endurance:  1,
			Charisma:   2,
			Agility:    2,
		},
	},

	// power armor
	80: {
		Name: "T-45a Power Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:   2,
			Perception: 3,
		},
	},
	81: {
		Name: "T-45d Power Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:   2,
			Perception: 4,
		},
	},
	82: {
		Name: "T-45f Power Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:   2,
			Perception: 5,
		},
	},
	83: {
		Name: "T-51a Power Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:   3,
			Perception: 1,
		},
	},
	84: {
		Name: "T-51d Power Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:   3,
			Perception: 2,
		},
	},
	85: {
		Name: "T-51f Power Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:   4,
			Perception: 3,
		},
	},
	86: {
		Name: "T-60a Power Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:  2,
			Endurance: 3,
		},
	},
	87: {
		Name: "T-60d Power Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:  2,
			Endurance: 4,
		},
	},
	88: {
		Name: "T-60f Power Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:   1,
			Perception: 1,
			Endurance:  5,
		},
	},
	89: {
		Name: "X-01 Mk I Power Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:   3,
			Perception: 1,
			Endurance:  1,
		},
	},
	90: {
		Name: "X-01 Mk IV Power Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:   4,
			Perception: 1,
			Endurance:  1,
		},
	},
	91: {
		Name: "X-01 Mk VI Power Armor",
		Tier: 3,
		Stats: SPECIAL{
			Strength:   5,
			Perception: 1,
			Endurance:  1,
		},
	},
}
