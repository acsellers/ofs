package game

type Clothing int

type ClothingStat struct {
	ID     ItemID
	Name   string
	Tier   int // 0 to 5
	Resist int
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
		Name:   "Vault Suit",
		Tier:   0,
		Resist: 0,
	},

	// vault suit
	1: {
		Name:   "Armored Vault Uniform",
		Tier:   1,
		Resist: 1,
		Stats: SPECIAL{
			Power:     1,
			Awareness: 2,
		},
	},
	2: {
		Name:   "Sturdy Vault Suit",
		Tier:   2,
		Resist: 2,
		Stats: SPECIAL{
			Power:     2,
			Awareness: 4,
		},
	},
	3: {
		Name:   "Heavy Vault Suit",
		Tier:   3,
		Resist: 3,
		Stats: SPECIAL{
			Power:     3,
			Awareness: 6,
		},
	},

	// battle armor
	4: {
		Name:   "Battle Armor",
		Tier:   1,
		Resist: 3,
		Stats: SPECIAL{
			Power:    2,
			Tenacity: 1,
		},
	},
	5: {
		Name:   "Sturdy Battle Armor",
		Tier:   2,
		Resist: 3,
		Stats: SPECIAL{
			Power:    4,
			Tenacity: 2,
		},
	},
	6: {
		Name:   "Heavy Battle Armor",
		Tier:   3,
		Resist: 4,
		Stats: SPECIAL{
			Power:    6,
			Tenacity: 3,
		},
	},

	// combat armor
	7: {
		Name:   "Combat Armor",
		Tier:   1,
		Resist: 2,
		Stats: SPECIAL{
			Power:  2,
			Nimble: 1,
		},
	},
	8: {
		Name:   "Sturdy Combat Armor",
		Tier:   2,
		Resist: 3,
		Stats: SPECIAL{
			Power:  4,
			Nimble: 2,
		},
	},
	9: {
		Name:   "Heavy Combat Armor",
		Tier:   3,
		Resist: 3,
		Stats: SPECIAL{
			Power:  5,
			Nimble: 4,
		},
	},

	// formal wear
	10: {
		Name:   "Formal Wear",
		Tier:   1,
		Resist: 0,
		Stats: SPECIAL{
			Outgoing: 1,
			Success:  2,
		},
	},
	11: {
		Name:   "Fancy Formal Wear",
		Tier:   2,
		Resist: 0,
		Stats: SPECIAL{
			Outgoing: 2,
			Success:  4,
		},
	},
	12: {
		Name:   "Successy Formal Wear",
		Tier:   3,
		Resist: 1,
		Stats: SPECIAL{
			Outgoing: 2,
			Success:  7,
		},
	},

	// jumpsuit
	13: {
		Name:   "Handyman Jumpsuit",
		Tier:   1,
		Resist: 1,
		Stats: SPECIAL{
			Nimble: 3,
		},
	},
	14: {
		Name:   "Advanced Jumpsuit",
		Tier:   2,
		Resist: 1,
		Stats: SPECIAL{
			Nimble:   4,
			Erudite:  1,
			Tenacity: 1,
		},
	},
	15: {
		Name:   "Expert Jumpsuit",
		Tier:   3,
		Resist: 1,
		Stats: SPECIAL{
			Nimble:   7,
			Erudite:  1,
			Tenacity: 1,
		},
	},

	// lab coat
	16: {
		Name:   "Lab Coat",
		Tier:   1,
		Resist: 0,
		Stats: SPECIAL{
			Erudite: 3,
		},
	},
	17: {
		Name:   "Advanced Lab Coat",
		Tier:   2,
		Resist: 0,
		Stats: SPECIAL{
			Erudite: 5,
			Success: 1,
		},
	},
	18: {
		Name:   "Expert Lab Coat",
		Tier:   3,
		Resist: 0,
		Stats: SPECIAL{
			Erudite:  7,
			Success:  1,
			Outgoing: 1,
		},
	},

	// leather armor
	19: {
		Name:   "Leather Armor",
		Tier:   1,
		Resist: 2,
		Stats: SPECIAL{
			Power:    1,
			Tenacity: 2,
		},
	},
	20: {
		Name:   "Sturdy Leather Armor",
		Tier:   2,
		Resist: 2,
		Stats: SPECIAL{
			Power:    2,
			Tenacity: 4,
		},
	},
	21: {
		Name:   "Heavy Leather Armor",
		Tier:   3,
		Resist: 3,
		Stats: SPECIAL{
			Power:    3,
			Tenacity: 6,
		},
	},

	// merc gear
	22: {
		Name:   "Merc Gear",
		Tier:   1,
		Resist: 1,
		Stats: SPECIAL{
			Awareness: 1,
			Nimble:    2,
		},
	},
	23: {
		Name:   "Sturdy Merc Gear",
		Tier:   2,
		Resist: 2,
		Stats: SPECIAL{
			Awareness: 1,
			Nimble:    3,
			Success:   2,
		},
	},
	24: {
		Name:   "Heavy Merc Gear",
		Tier:   3,
		Resist: 2,
		Stats: SPECIAL{
			Awareness: 2,
			Nimble:    4,
			Success:   3,
		},
	},

	// military fatigues
	25: {
		Name:   "Military Fatigues",
		Tier:   1,
		Resist: 1,
		Stats: SPECIAL{
			Power: 3,
		},
	},
	26: {
		Name:   "Officer Fatigues",
		Tier:   2,
		Resist: 2,
		Stats: SPECIAL{
			Power:    5,
			Tenacity: 1,
		},
	},
	27: {
		Name:   "Commander Fatigues",
		Tier:   3,
		Resist: 2,
		Stats: SPECIAL{
			Power:    7,
			Tenacity: 2,
		},
	},

	// nightwear
	28: {
		Name:   "Nightwear",
		Tier:   1,
		Resist: 0,
		Stats: SPECIAL{
			Success: 3,
		},
	},
	29: {
		Name:   "Naughty Nightwear",
		Tier:   2,
		Resist: 0,
		Stats: SPECIAL{
			Success: 4,
			Charm:   2,
		},
	},
	30: {
		Name:   "Magical Nightwear",
		Tier:   3,
		Resist: 0,
		Stats: SPECIAL{
			Success: 7,
			Charm:   2,
		},
	},

	// officer uniform
	31: {
		Name:   "Junior Officer Uniform",
		Tier:   1,
		Resist: 1,
		Stats: SPECIAL{
			Charm:   1,
			Erudite: 2,
		},
	},
	32: {
		Name:   "Officer Uniform",
		Tier:   2,
		Resist: 1,
		Stats: SPECIAL{
			Charm:   2,
			Erudite: 4,
		},
	},
	33: {
		Name:   "Commander Uniform",
		Tier:   3,
		Resist: 1,
		Stats: SPECIAL{
			Charm:   3,
			Erudite: 4,
			Success: 2,
		},
	},

	// radiation suit
	34: {
		Name:   "Radiation Suit",
		Tier:   1,
		Resist: 1,
		Stats: SPECIAL{
			Awareness: 1,
			Tenacity:  2,
		},
	},
	35: {
		Name:   "Advanced Radiation Suit",
		Tier:   2,
		Resist: 2,
		Stats: SPECIAL{
			Awareness: 2,
			Tenacity:  4,
		},
	},
	36: {
		Name:   "Expert Radiation Suit",
		Tier:   3,
		Resist: 2,
		Stats: SPECIAL{
			Awareness: 4,
			Tenacity:  5,
		},
	},

	// raider armor
	37: {
		Name:   "Raider Armor",
		Tier:   1,
		Resist: 1,
		Stats: SPECIAL{
			Awareness: 1,
			Nimble:    2,
		},
	},
	38: {
		Name:   "Sturdy Raider Armor",
		Tier:   2,
		Resist: 2,
		Stats: SPECIAL{
			Awareness: 2,
			Nimble:    3,
			Tenacity:  1,
		},
	},
	39: {
		Name:   "Heavy Raider Armor",
		Tier:   3,
		Resist: 2,
		Stats: SPECIAL{
			Awareness: 3,
			Nimble:    4,
			Tenacity:  2,
		},
	},

	// robes
	40: {
		Name:   "Initiate Robe",
		Tier:   1,
		Resist: 0,
		Stats: SPECIAL{
			Charm:  2,
			Nimble: 1,
		},
	},
	41: {
		Name:   "Scribe Robe",
		Tier:   2,
		Resist: 1,
		Stats: SPECIAL{
			Charm:   3,
			Nimble:  2,
			Erudite: 1,
		},
	},
	42: {
		Name:   "Scribe Rothchild's Robe",
		Tier:   3,
		Resist: 1,
		Stats: SPECIAL{
			Awareness: 2,
			Outgoing:  3,
			Charm:     2,
			Erudite:   2,
		},
	},

	// wasteland gear
	43: {
		Name:   "Wasteland Gear",
		Tier:   1,
		Resist: 1,
		Stats: SPECIAL{
			Tenacity: 3,
		},
	},
	44: {
		Name:   "Sturdy Wasteland Gear",
		Tier:   2,
		Resist: 2,
		Stats: SPECIAL{
			Tenacity:  5,
			Awareness: 1,
		},
	},
	45: {
		Name:   "Heavy Wasteland Gear",
		Tier:   3,
		Resist: 3,
		Stats: SPECIAL{
			Tenacity:  7,
			Awareness: 2,
		},
	},

	// wasteland medic
	46: {
		Name:   "Wasteland Medic",
		Tier:   1,
		Resist: 1,
		Stats: SPECIAL{
			Awareness: 2,
			Success:   1,
		},
	},
	47: {
		Name:   "Wasteland Doctor",
		Tier:   2,
		Resist: 1,
		Stats: SPECIAL{
			Tenacity:  1,
			Awareness: 3,
			Success:   2,
		},
	},
	48: {
		Name:   "Wasteland Surgeon",
		Tier:   3,
		Resist: 1,
		Stats: SPECIAL{
			Tenacity:  2,
			Awareness: 4,
			Success:   3,
		},
	},

	// rare outfits
	49: {
		Name:   "Clergy Outfit",
		Tier:   2,
		Resist: 0,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Charm:    4,
			Success:  1,
			Outgoing: 1,
		},
	},
	50: {
		Name:   "Comedian Outfit",
		Tier:   2,
		Resist: 0,
		Stats: SPECIAL{
			Charm:    2,
			Success:  1,
			Outgoing: 3,
		},
	},
	51: {
		Name:   "Engineer Outfit",
		Tier:   2,
		Resist: 2,
		Stats: SPECIAL{
			Tenacity: 3,
			Erudite:  2,
			Success:  1,
		},
	},
	52: {
		Name:   "Greaser Outfit",
		Tier:   2,
		Resist: 1,
		Stats: SPECIAL{
			Charm:   3,
			Nimble:  2,
			Success: 1,
		},
	},
	53: {
		Name:   "Horror Fan Outfit",
		Tier:   2,
		Resist: 1,
		Stats: SPECIAL{
			Tenacity: 4,
			Success:  1,
			Outgoing: 1,
		},
	},
	54: {
		Name:   "Knight Armor",
		Tier:   2,
		Resist: 3,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Power:     2,
			Awareness: 2,
			Tenacity:  1,
			Success:   1,
		},
	},
	55: {
		Name:   "Librarian Outfit",
		Tier:   2,
		Resist: 0,
		Sex:    FemaleOnly,
		Stats: SPECIAL{
			Erudite: 4,
			Success: 2,
		},
	},
	56: {
		Name:   "Mayor Outfit",
		Tier:   2,
		Resist: 0,
		Stats: SPECIAL{
			Charm:   3,
			Erudite: 2,
			Success: 1,
		},
	},
	57: {
		Name:   "Medieval Ruler Outfit",
		Tier:   2,
		Resist: 2,
		Stats: SPECIAL{
			Awareness: 3,
			Charm:     2,
			Success:   1,
		},
	},
	58: {
		Name:   "Movie Fan Outfit",
		Tier:   2,
		Resist: 0,
		Sex:    FemaleOnly,
		Stats: SPECIAL{
			Awareness: 4,
			Success:   1,
			Outgoing:  1,
		},
	},
	59: {
		Name:   "Ninja Outfit",
		Tier:   2,
		Resist: 2,
		Stats: SPECIAL{
			Nimble:    4,
			Success:   1,
			Awareness: 1,
		},
	},
	60: {
		Name:   "Nobility Outfit",
		Tier:   2,
		Resist: 3,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Tenacity: 2,
			Erudite:  2,
			Success:  2,
		},
	},
	61: {
		Name:   "Professor Outfit",
		Tier:   2,
		Resist: 0,
		Stats: SPECIAL{
			Erudite: 4,
			Success: 2,
		},
	},
	62: {
		Name:   "Republic Robes",
		Tier:   2,
		Resist: 0,
		Sex:    FemaleOnly,
		Stats: SPECIAL{
			Charm:    4,
			Success:  1,
			Outgoing: 1,
		},
	},
	63: {
		Name:   "Sci-fi Fan Outfit",
		Tier:   2,
		Resist: 0,
		Stats: SPECIAL{
			Erudite: 2,
			Nimble:  3,
			Success: 1,
		},
	},
	64: {
		Name:   "Soldier Uniform",
		Tier:   2,
		Resist: 2,
		Stats: SPECIAL{
			Power:    3,
			Tenacity: 2,
			Success:  1,
		},
	},
	65: {
		Name:   "Sports Fan Outfit",
		Tier:   2,
		Resist: 2,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Power:    4,
			Tenacity: 1,
			Success:  1,
		},
	},
	66: {
		Name:   "Surgeon Outfit",
		Tier:   2,
		Resist: 1,
		Stats: SPECIAL{
			Awareness: 2,
			Nimble:    2,
			Success:   1,
			Erudite:   1,
		},
	},
	67: {
		Name:   "Survivor Armor",
		Tier:   2,
		Resist: 2,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Power:    2,
			Nimble:   2,
			Tenacity: 1,
			Success:  1,
		},
	},
	68: {
		Name:   "Wrestler Outfit",
		Tier:   2,
		Resist: 2,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Power:    3,
			Tenacity: 2,
			Success:  1,
		},
	},

	// legendary outfits
	69: {
		Name:   "Autumn's Uniform",
		Tier:   3,
		Resist: 2,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Power:     3,
			Awareness: 3,
			Tenacity:  2,
			Charm:     1,
		},
	},
	70: {
		Name:   "Bittercup's Outfit",
		Tier:   3,
		Resist: 1,
		Sex:    FemaleOnly,
		Stats: SPECIAL{
			Power:     2,
			Awareness: 3,
			Tenacity:  3,
			Charm:     1,
		},
	},
	71: {
		Name:   "Confessor Cromwell's Rags",
		Tier:   3,
		Resist: 1,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Awareness: 3,
			Tenacity:  3,
			Erudite:   1,
			Success:   2,
		},
	},
	72: {
		Name:   "Elder Robe",
		Tier:   3,
		Resist: 1,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Charm:    4,
			Nimble:   3,
			Tenacity: 1,
			Success:  1,
		},
	},
	73: {
		Name:   "Eulogy Jones' Suit",
		Tier:   3,
		Resist: 1,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Awareness: 2,
			Charm:     3,
			Erudite:   1,
			Success:   3,
		},
	},
	74: {
		Name:   "Minuteman Uniform",
		Tier:   3,
		Resist: 2,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Power:     3,
			Awareness: 2,
			Erudite:   2,
			Nimble:    2,
		},
	},
	75: {
		Name:   "TODO: New Clothing",
		Tier:   3,
		Resist: 1,
		Stats: SPECIAL{
			Awareness: 2,
			Tenacity:  3,
			Charm:     2,
			Erudite:   2,
		},
	},
	76: {
		Name:   "Sheriff's Duster",
		Tier:   3,
		Resist: 2,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Nimble:    2,
			Awareness: 2,
			Tenacity:  5,
		},
	},
	77: {
		Name:   "Tenpenny's Suit",
		Tier:   3,
		Resist: 1,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Awareness: 2,
			Charm:     3,
			Erudite:   3,
			Success:   1,
		},
	},
	78: {
		Name:   "Three Dog's Outfit",
		Tier:   3,
		Resist: 1,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Awareness: 1,
			Charm:     4,
			Outgoing:  4,
		},
	},
	79: {
		Name:   "Tunnel Snakes' Outfit",
		Tier:   3,
		Resist: 1,
		Sex:    MaleOnly,
		Stats: SPECIAL{
			Awareness: 3,
			Tenacity:  1,
			Charm:     3,
			Nimble:    2,
		},
	},

	// power armor
	80: {
		Name:   "T-45a Power Armor",
		Tier:   2,
		Resist: 3,
		Stats: SPECIAL{
			Power:     2,
			Awareness: 3,
			Tenacity:  1,
		},
	},
	81: {
		Name:   "T-45d Power Armor",
		Tier:   3,
		Resist: 1,
		Stats: SPECIAL{
			Power:     3,
			Awareness: 4,
			Tenacity:  2,
		},
	},
	82: {
		Name:   "T-45f Power Armor",
		Tier:   3,
		Resist: 4,
		Stats: SPECIAL{
			Power:     2,
			Awareness: 5,
			Tenacity:  2,
		},
	},
	83: {
		Name:   "T-51a Power Armor",
		Tier:   2,
		Resist: 3,
		Stats: SPECIAL{
			Power:     4,
			Awareness: 2,
		},
	},
	84: {
		Name:   "T-51d Power Armor",
		Tier:   3,
		Resist: 4,
		Stats: SPECIAL{
			Power:     5,
			Awareness: 4,
		},
	},
	85: {
		Name:   "T-51f Power Armor",
		Tier:   4,
		Resist: 5,
		Stats: SPECIAL{
			Power:     5,
			Awareness: 4,
			Tenacity:  3,
		},
	},
	86: {
		Name:   "T-60a Power Armor",
		Tier:   3,
		Resist: 4,
		Stats: SPECIAL{
			Awareness: 2,
			Power:     3,
			Tenacity:  4,
		},
	},
	87: {
		Name:   "T-60d Power Armor",
		Tier:   3,
		Resist: 5,
		Stats: SPECIAL{
			Awareness: 2,
			Power:     3,
			Tenacity:  4,
		},
	},
	88: {
		Name:   "T-60f Power Armor",
		Tier:   4,
		Resist: 5,
		Stats: SPECIAL{
			Awareness: 2,
			Power:     4,
			Tenacity:  6,
		},
	},
	89: {
		Name:   "X-01 Mk I Power Armor",
		Tier:   3,
		Resist: 4,
		Stats: SPECIAL{
			Power:     3,
			Awareness: 4,
			Tenacity:  2,
		},
	},
	90: {
		Name:   "X-01 Mk IV Power Armor",
		Tier:   4,
		Resist: 5,
		Stats: SPECIAL{
			Power:     7,
			Awareness: 3,
			Tenacity:  2,
		},
	},
	91: {
		Name:   "X-01 Mk VI Power Armor",
		Tier:   4,
		Resist: 5,
		Stats: SPECIAL{
			Power:     5,
			Awareness: 4,
			Tenacity:  3,
		},
	},
}
