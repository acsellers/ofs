package game

import (
	"fmt"
	"strings"
	"unicode"
)

type Weapon int

type WeaponStat struct {
	ID     ItemID
	Name   string
	Level  int
	Parent string

	DamageMin int
	DamageMax int
}

func (ws WeaponStat) Sprite() string {
	name := ws.Name
	if ws.Parent != "" {
		name = ws.Parent
	}

	return strings.ToLower(
		strings.TrimFunc(
			name,
			func(r rune) bool {
				return unicode.IsLetter(r) || unicode.IsDigit(r)
			},
		),
	)
}

func (ws WeaponStat) DPS() float32 {
	return float32(ws.DamageMin+ws.DamageMax) / 2.0
}

func (ws WeaponStat) DamageRange() string {
	if ws.DamageMin == ws.DamageMax {
		return fmt.Sprint(ws.DamageMax)
	}
	return fmt.Sprintf("%d-%d", ws.DamageMin, ws.DamageMax)
}

func init() {
	for k, w := range Weapons {
		w.ID = ItemID(int(k) + 1000)
		Weapons[k] = w
	}
}

var Weapons = map[Weapon]WeaponStat{
	0: {
		Name:      "Fist",
		Level:     0,
		DamageMin: 1,
		DamageMax: 1,
	},

	// .32 pistol
	1: {
		Name:      "Rusty .32 Pistol",
		Level:     1,
		Parent:    ".32 Pistol",
		DamageMin: 1,
		DamageMax: 1,
	},
	2: {
		Name:      ".32 Pistol",
		Level:     1,
		DamageMin: 1,
		DamageMax: 2,
	},
	3: {
		Name:      "Enhanced .32 Pistol",
		Level:     1,
		Parent:    ".32 Pistol",
		DamageMin: 1,
		DamageMax: 3,
	},
	4: {
		Name:      "Hardened .32 Pistol",
		Level:     2,
		Parent:    ".32 Pistol",
		DamageMin: 1,
		DamageMax: 4,
	},
	5: {
		Name:      "Armor Piercing .32 Pistol",
		Level:     2,
		Parent:    ".32 Pistol",
		DamageMin: 1,
		DamageMax: 5,
	},
	6: {
		Name:      "Wild Bill's Sidearm",
		Level:     3,
		Parent:    ".32 Pistol",
		DamageMin: 1,
		DamageMax: 6,
	},

	// 10mm pistol
	7: {
		Name:      "Rusty 10mm Pistol",
		Level:     1,
		Parent:    "10mm Pistol",
		DamageMin: 2,
		DamageMax: 2,
	},
	8: {
		Name:      "10mm Pistol",
		Level:     1,
		DamageMin: 2,
		DamageMax: 3,
	},
	9: {
		Name:      "Enhanced 10mm Pistol",
		Level:     1,
		Parent:    "10mm Pistol",
		DamageMin: 2,
		DamageMax: 4,
	},
	10: {
		Name:      "Hardened 10mm Pistol",
		Level:     2,
		Parent:    "10mm Pistol",
		DamageMin: 2,
		DamageMax: 5,
	},
	11: {
		Name:      "Armor Piercing 10mm Pistol",
		Level:     2,
		Parent:    "10mm Pistol",
		DamageMin: 2,
		DamageMax: 6,
	},
	12: {
		Name:      "Lone Wanderer",
		Level:     3,
		Parent:    "10mm Pistol",
		DamageMin: 2,
		DamageMax: 7,
	},

	// Scoped .44
	13: {
		Name:      "Rusty Scoped .44",
		Level:     1,
		Parent:    "Scoped .44",
		DamageMin: 3,
		DamageMax: 3,
	},
	14: {
		Name:      "Scoped .44",
		Level:     1,
		Parent:    "Scoped .44",
		DamageMin: 3,
		DamageMax: 4,
	},
	15: {
		Name:      "Enhanced Scoped .44",
		Level:     1,
		Parent:    "Scoped .44",
		DamageMin: 3,
		DamageMax: 5,
	},
	16: {
		Name:      "Hardened Scoped .44",
		Level:     2,
		Parent:    "Scoped .44",
		DamageMin: 3,
		DamageMax: 6,
	},
	17: {
		Name:      "Armor Piercing Scoped .44",
		Level:     2,
		Parent:    "Scoped .44",
		DamageMin: 3,
		DamageMax: 7,
	},
	18: {
		Name:      "Blackhawk",
		Level:     3,
		Parent:    "Scoped .44",
		DamageMin: 3,
		DamageMax: 8,
	},

	// assault rifle
	19: {
		Name:      "Rusty Assault Rifle",
		Level:     2,
		Parent:    "Assault Rifle",
		DamageMin: 8,
		DamageMax: 8,
	},
	20: {
		Name:      "Assault Rifle",
		Level:     2,
		DamageMin: 8,
		DamageMax: 9,
	},
	21: {
		Name:      "Enhanced Assault Rifle",
		Level:     2,
		Parent:    "Assault Rifle",
		DamageMin: 8,
		DamageMax: 10,
	},
	22: {
		Name:      "Hardened Assault Rifle",
		Level:     2,
		Parent:    "Assault Rifle",
		DamageMin: 8,
		DamageMax: 11,
	},
	23: {
		Name:      "Armor Piercing Assault Rifle",
		Level:     2,
		Parent:    "Assault Rifle",
		DamageMin: 8,
		DamageMax: 12,
	},
	24: {
		Name:      "Infiltrator",
		Level:     3,
		Parent:    "Assault Rifle",
		DamageMin: 8,
		DamageMax: 13,
	},

	// bb gun
	25: {
		Name:      "Rusty BB Gun",
		Level:     1,
		Parent:    "BB Gun",
		DamageMin: 0,
		DamageMax: 1,
	},
	26: {
		Name:      "BB Gun",
		Level:     1,
		DamageMin: 0,
		DamageMax: 2,
	},
	27: {
		Name:      "Enhanced BB Gun",
		Level:     1,
		Parent:    "BB Gun",
		DamageMin: 0,
		DamageMax: 3,
	},
	28: {
		Name:      "Hardened BB Gun",
		Level:     2,
		Parent:    "BB Gun",
		DamageMin: 0,
		DamageMax: 4,
	},
	29: {
		Name:      "Armor Piercing BB Gun",
		Level:     2,
		Parent:    "BB Gun",
		DamageMin: 0,
		DamageMax: 5,
	},
	30: {
		Name:      "Red Rocket",
		Level:     3,
		Parent:    "BB Gun",
		DamageMin: 0,
		DamageMax: 6,
	},

	// hunting rifle
	31: {
		Name:      "Rusty Hunting Rifle",
		Level:     1,
		Parent:    "Hunting Rifle",
		DamageMin: 5,
		DamageMax: 5,
	},
	32: {
		Name:      "Hunting Rifle",
		Level:     1,
		DamageMin: 5,
		DamageMax: 6,
	},
	33: {
		Name:      "Enhanced Hunting Rifle",
		Level:     1,
		Parent:    "Hunting Rifle",
		DamageMin: 5,
		DamageMax: 7,
	},
	34: {
		Name:      "Hardened Hunting Rifle",
		Level:     2,
		Parent:    "Hunting Rifle",
		DamageMin: 5,
		DamageMax: 8,
	},
	35: {
		Name:      "Armor Piercing Hunting Rifle",
		Level:     2,
		Parent:    "Hunting Rifle",
		DamageMin: 5,
		DamageMax: 9,
	},
	36: {
		Name:      "Ol' Painless",
		Level:     3,
		Parent:    "Hunting Rifle",
		DamageMin: 5,
		DamageMax: 10,
	},

	// Lever-action Rifle
	37: {
		Name:      "Rusty Lever-action Rifle",
		Level:     1,
		Parent:    "Lever-action Rifle",
		DamageMin: 4,
		DamageMax: 4,
	},
	38: {
		Name:      "Lever-action Rifle",
		Level:     1,
		DamageMin: 4,
		DamageMax: 5,
	},
	39: {
		Name:      "Enhanced Lever-action Rifle",
		Level:     1,
		Parent:    "Lever-action Rifle",
		DamageMin: 4,
		DamageMax: 6,
	},
	40: {
		Name:      "Hardened Lever-action Rifle",
		Level:     2,
		Parent:    "Lever-action Rifle",
		DamageMin: 4,
		DamageMax: 7,
	},
	41: {
		Name:      "Armor Piercing Lever-action Rifle",
		Level:     2,
		Parent:    "Lever-action Rifle",
		DamageMin: 4,
		DamageMax: 8,
	},
	42: {
		Name:      "Lincoln's Repeater",
		Level:     3,
		Parent:    "Lever-action Rifle",
		DamageMin: 4,
		DamageMax: 9,
	},

	// Railway Rifle
	43: {
		Name:      "Rusty Railway Rifle",
		Level:     2,
		Parent:    "Railway Rifle",
		DamageMin: 14,
		DamageMax: 14,
	},
	44: {
		Name:      "Railway Rifle",
		Level:     2,
		DamageMin: 14,
		DamageMax: 15,
	},
	45: {
		Name:      "Enhanced Railway Rifle",
		Level:     2,
		Parent:    "Railway Rifle",
		DamageMin: 14,
		DamageMax: 16,
	},
	46: {
		Name:      "Hardened Railway Rifle",
		Level:     3,
		Parent:    "Railway Rifle",
		DamageMin: 14,
		DamageMax: 17,
	},
	47: {
		Name:      "Armor Piercing Railway Rifle",
		Level:     3,
		Parent:    "Railway Rifle",
		DamageMin: 14,
		DamageMax: 18,
	},
	48: {
		Name:      "Railmaster",
		Level:     3,
		Parent:    "Railway Rifle",
		DamageMin: 14,
		DamageMax: 19,
	},

	// Sniper Rifle
	49: {
		Name:      "Rusty Sniper Rifle",
		Level:     2,
		Parent:    "Sniper Rifle",
		DamageMin: 10,
		DamageMax: 10,
	},
	50: {
		Name:      "Sniper Rifle",
		Level:     2,
		DamageMin: 10,
		DamageMax: 11,
	},
	51: {
		Name:      "Enhanced Sniper Rifle",
		Level:     2,
		Parent:    "Sniper Rifle",
		DamageMin: 10,
		DamageMax: 12,
	},
	62: {
		Name:      "Hardened Sniper Rifle",
		Level:     2,
		Parent:    "Sniper Rifle",
		DamageMin: 10,
		DamageMax: 13,
	},
	63: {
		Name:      "Armor Piercing Sniper Rifle",
		Level:     2,
		Parent:    "Sniper Rifle",
		DamageMin: 10,
		DamageMax: 14,
	},
	64: {
		Name:      "Victory Rifle",
		Level:     3,
		Parent:    "Sniper Rifle",
		DamageMin: 10,
		DamageMax: 15,
	},

	// Combat Shotgun
	65: {
		Name:      "Rusty Combat Shotgun",
		Level:     2,
		Parent:    "Combat Shotgun",
		DamageMin: 13,
		DamageMax: 13,
	},
	66: {
		Name:      "Combat Shotgun",
		Level:     2,
		DamageMin: 13,
		DamageMax: 14,
	},
	67: {
		Name:      "Enhanced Combat Shotgun",
		Level:     2,
		Parent:    "Combat Shotgun",
		DamageMin: 13,
		DamageMax: 15,
	},
	68: {
		Name:      "Hardened Combat Shotgun",
		Level:     2,
		Parent:    "Combat Shotgun",
		DamageMin: 13,
		DamageMax: 16,
	},
	69: {
		Name:      "Double-barrel Combat Shotgun",
		Level:     3,
		Parent:    "Combat Shotgun",
		DamageMin: 13,
		DamageMax: 17,
	},
	70: {
		Name:      "Charon's Shotgun",
		Level:     3,
		Parent:    "Combat Shotgun",
		DamageMin: 13,
		DamageMax: 18,
	},

	// Sawed-off Shotgun
	71: {
		Name:      "Rusty Sawed-off Shotgun",
		Level:     1,
		Parent:    "Sawed-off Shotgun",
		DamageMin: 6,
		DamageMax: 6,
	},
	72: {
		Name:      "Sawed-off Shotgun",
		Level:     1,
		DamageMin: 6,
		DamageMax: 7,
	},
	73: {
		Name:      "Enhanced Sawed-off Shotgun",
		Level:     1,
		Parent:    "Sawed-off Shotgun",
		DamageMin: 6,
		DamageMax: 8,
	},
	74: {
		Name:      "Hardened Sawed-off Shotgun",
		Level:     2,
		Parent:    "Sawed-off Shotgun",
		DamageMin: 6,
		DamageMax: 9,
	},
	75: {
		Name:      "Double-barrel Sawed-off Shotgun",
		Level:     2,
		Parent:    "Sawed-off Shotgun",
		DamageMin: 6,
		DamageMax: 10,
	},
	76: {
		Name:      "Kneecapper",
		Level:     3,
		Parent:    "Sawed-off Shotgun",
		DamageMin: 6,
		DamageMax: 11,
	},

	// Shotgun
	77: {
		Name:      "Rusty Shotgun",
		Level:     2,
		Parent:    "Shotgun",
		DamageMin: 9,
		DamageMax: 9,
	},
	78: {
		Name:      "Shotgun",
		Level:     2,
		DamageMin: 9,
		DamageMax: 10,
	},
	79: {
		Name:      "Enhanced Shotgun",
		Level:     2,
		Parent:    "Shotgun",
		DamageMin: 9,
		DamageMax: 11,
	},
	80: {
		Name:      "Hardened Shotgun",
		Level:     2,
		Parent:    "Shotgun",
		DamageMin: 9,
		DamageMax: 12,
	},
	81: {
		Name:      "Double-barrel Shotgun",
		Level:     2,
		Parent:    "Shotgun",
		DamageMin: 9,
		DamageMax: 13,
	},
	82: {
		Name:      "Farmer's Daughter",
		Level:     3,
		Parent:    "Shotgun",
		DamageMin: 9,
		DamageMax: 14,
	},

	// Alien Blaster
	83: {
		Name:      "Rusty Alien Blaster",
		Level:     3,
		Parent:    "Alien Blaster",
		DamageMin: 18,
		DamageMax: 18,
	},
	84: {
		Name:      "Alien Blaster",
		Level:     3,
		DamageMin: 18,
		DamageMax: 19,
	},
	85: {
		Name:      "Tuned Alien Blaster",
		Level:     3,
		Parent:    "Alien Blaster",
		DamageMin: 18,
		DamageMax: 20,
	},
	86: {
		Name:      "Focused Alien Blaster",
		Level:     3,
		Parent:    "Alien Blaster",
		DamageMin: 18,
		DamageMax: 21,
	},
	87: {
		Name:      "Amplified Alien Blaster",
		Level:     3,
		Parent:    "Alien Blaster",
		DamageMin: 18,
		DamageMax: 22,
	},
	88: {
		Name:      "Destabilizer",
		Level:     3,
		Parent:    "Alien Blaster",
		DamageMin: 18,
		DamageMax: 23,
	},

	// Laser Pistol
	89: {
		Name:      "Rusty Laser Pistol",
		Level:     1,
		Parent:    "Laser Pistol",
		DamageMin: 7,
		DamageMax: 7,
	},
	90: {
		Name:      "Laser Pistol",
		Level:     2,
		DamageMin: 7,
		DamageMax: 8,
	},
	91: {
		Name:      "Tuned Laser Pistol",
		Level:     2,
		Parent:    "Laser Pistol",
		DamageMin: 7,
		DamageMax: 9,
	},
	92: {
		Name:      "Focused Laser Pistol",
		Level:     2,
		Parent:    "Laser Pistol",
		DamageMin: 7,
		DamageMax: 10,
	},
	93: {
		Name:      "Amplified Laser Pistol",
		Level:     2,
		Parent:    "Laser Pistol",
		DamageMin: 7,
		DamageMax: 11,
	},
	94: {
		Name:      "Smuggler's End",
		Level:     3,
		Parent:    "Laser Pistol",
		DamageMin: 7,
		DamageMax: 12,
	},

	// Plasma Pistol
	95: {
		Name:      "Rusty Plasma Pistol",
		Level:     2,
		Parent:    "Plasma Pistol",
		DamageMin: 11,
		DamageMax: 11,
	},
	96: {
		Name:      "Plasma Pistol",
		Level:     2,
		DamageMin: 11,
		DamageMax: 12,
	},
	97: {
		Name:      "Tuned Plasma Pistol",
		Level:     2,
		Parent:    "Plasma Pistol",
		DamageMin: 11,
		DamageMax: 13,
	},
	98: {
		Name:      "Focused Plasma Pistol",
		Level:     2,
		Parent:    "Plasma Pistol",
		DamageMin: 11,
		DamageMax: 14,
	},
	99: {
		Name:      "Amplified Plasma Pistol",
		Level:     3,
		Parent:    "Plasma Pistol",
		DamageMin: 11,
		DamageMax: 15,
	},
	100: {
		Name:      "MPXL Novasurge",
		Level:     3,
		Parent:    "Plasma Pistol",
		DamageMin: 11,
		DamageMax: 16,
	},

	// Gauss Rifle
	101: {
		Name:      "Rusty Gauss Rifle",
		Level:     2,
		Parent:    "Gauss Rifle",
		DamageMin: 16,
		DamageMax: 16,
	},
	102: {
		Name:      "Gauss Rifle",
		Level:     2,
		DamageMin: 16,
		DamageMax: 17,
	},
	103: {
		Name:      "Enhanced Gauss Rifle",
		Level:     3,
		Parent:    "Gauss Rifle",
		DamageMin: 16,
		DamageMax: 18,
	},
	104: {
		Name:      "Hardened Gauss Rifle",
		Level:     3,
		Parent:    "Gauss Rifle",
		DamageMin: 16,
		DamageMax: 19,
	},
	105: {
		Name:      "Acclerated Gauss Rifle",
		Level:     3,
		Parent:    "Gauss Rifle",
		DamageMin: 16,
		DamageMax: 20,
	},
	106: {
		Name:      "Magnetron 4000",
		Level:     3,
		Parent:    "Gauss Rifle",
		DamageMin: 16,
		DamageMax: 21,
	},

	// Laser Rifle
	107: {
		Name:      "Rusty Laser Rifle",
		Level:     2,
		Parent:    "Laser Rifle",
		DamageMin: 12,
		DamageMax: 12,
	},
	108: {
		Name:      "Laser Rifle",
		Level:     2,
		DamageMin: 12,
		DamageMax: 13,
	},
	109: {
		Name:      "Tuned Laser Rifle",
		Level:     2,
		Parent:    "Laser Rifle",
		DamageMin: 12,
		DamageMax: 14,
	},
	110: {
		Name:      "Focused Laser Rifle",
		Level:     2,
		Parent:    "Laser Rifle",
		DamageMin: 12,
		DamageMax: 15,
	},
	111: {
		Name:      "Amplified Laser Rifle",
		Level:     3,
		Parent:    "Laser Rifle",
		DamageMin: 12,
		DamageMax: 16,
	},
	112: {
		Name:      "Wazer Wifle",
		Level:     3,
		Parent:    "Laser Rifle",
		DamageMin: 12,
		DamageMax: 17,
	},

	// Plasma Rifle
	113: {
		Name:      "Rusty Plasma Rifle",
		Level:     2,
		Parent:    "Plasma Rifle",
		DamageMin: 17,
		DamageMax: 17,
	},
	114: {
		Name:      "Plasma Rifle",
		Level:     2,
		DamageMin: 17,
		DamageMax: 18,
	},
	115: {
		Name:      "Tuned Plasma Rifle",
		Level:     3,
		Parent:    "Plasma Rifle",
		DamageMin: 17,
		DamageMax: 19,
	},
	116: {
		Name:      "Focused Plasma Rifle",
		Level:     3,
		Parent:    "Plasma Rifle",
		DamageMin: 17,
		DamageMax: 20,
	},
	117: {
		Name:      "Amplified Plasma Rifle",
		Level:     3,
		Parent:    "Plasma Rifle",
		DamageMin: 17,
		DamageMax: 21,
	},
	118: {
		Name:      "Mean Green Monster",
		Level:     3,
		Parent:    "Plasma Rifle",
		DamageMin: 17,
		DamageMax: 22,
	},

	// Fat Man
	119: {
		Name:      "Rusty Fat Man",
		Level:     3,
		Parent:    "Fat Man",
		DamageMin: 22,
		DamageMax: 22,
	},
	120: {
		Name:      "Fat Man",
		Level:     3,
		DamageMin: 22,
		DamageMax: 23,
	},
	121: {
		Name:      "Enhanced Fat Man",
		Level:     3,
		Parent:    "Fat Man",
		DamageMin: 22,
		DamageMax: 24,
	},
	122: {
		Name:      "Hardened Fat Man",
		Level:     3,
		Parent:    "Fat Man",
		DamageMin: 22,
		DamageMax: 25,
	},
	123: {
		Name:      "Guided Fat Man",
		Level:     3,
		Parent:    "Fat Man",
		DamageMin: 22,
		DamageMax: 26,
	},
	124: {
		Name:      "MIRV",
		Level:     3,
		Parent:    "Fat Man",
		DamageMin: 22,
		DamageMax: 27,
	},

	// Flamer
	125: {
		Name:      "Rusty Flamer",
		Level:     2,
		Parent:    "Flamer",
		DamageMin: 15,
		DamageMax: 15,
	},
	126: {
		Name:      "Flamer",
		Level:     2,
		DamageMin: 15,
		DamageMax: 16,
	},
	127: {
		Name:      "Enhanced Flamer",
		Level:     2,
		Parent:    "Flamer",
		DamageMin: 15,
		DamageMax: 17,
	},
	128: {
		Name:      "Hardened Flamer",
		Level:     2,
		Parent:    "Flamer",
		DamageMin: 15,
		DamageMax: 18,
	},
	129: {
		Name:      "Pressurized Flamer",
		Level:     2,
		Parent:    "Flamer",
		DamageMin: 15,
		DamageMax: 19,
	},
	130: {
		Name:      "Burnmaster",
		Level:     3,
		Parent:    "Flamer",
		DamageMin: 15,
		DamageMax: 20,
	},

	// Gatling Laser
	131: {
		Name:      "Rusty Gatling Laser",
		Level:     3,
		Parent:    "Gatling Laser",
		DamageMin: 21,
		DamageMax: 21,
	},
	132: {
		Name:      "Gatling Laser",
		Level:     3,
		DamageMin: 21,
		DamageMax: 22,
	},
	133: {
		Name:      "Tuned Gatling Laser",
		Level:     3,
		Parent:    "Gatling Laser",
		DamageMin: 21,
		DamageMax: 23,
	},
	134: {
		Name:      "Focused Gatling Laser",
		Level:     3,
		Parent:    "Gatling Laser",
		DamageMin: 21,
		DamageMax: 24,
	},
	135: {
		Name:      "Amplified Gatling Laser",
		Level:     3,
		Parent:    "Gatling Laser",
		DamageMin: 21,
		DamageMax: 25,
	},
	136: {
		Name:      "Vengeance",
		Level:     3,
		Parent:    "Gatling Laser",
		DamageMin: 21,
		DamageMax: 26,
	},

	// Minigun
	137: {
		Name:      "Rusty Minigun",
		Level:     2,
		Parent:    "Minigun",
		DamageMin: 19,
		DamageMax: 19,
	},
	138: {
		Name:      "Minigun",
		Level:     2,
		DamageMin: 19,
		DamageMax: 20,
	},
	139: {
		Name:      "Enhanced Minigun",
		Level:     3,
		Parent:    "Minigun",
		DamageMin: 19,
		DamageMax: 21,
	},
	140: {
		Name:      "Hardened Minigun",
		Level:     3,
		Parent:    "Minigun",
		DamageMin: 19,
		DamageMax: 22,
	},
	141: {
		Name:      "Armor Piercing Minigun",
		Level:     3,
		Parent:    "Minigun",
		DamageMin: 19,
		DamageMax: 23,
	},
	142: {
		Name:      "Lead Belcher",
		Level:     3,
		Parent:    "Minigun",
		DamageMin: 19,
		DamageMax: 24,
	},

	// Missile Launcher
	143: {
		Name:      "Rusty Missile Launcher",
		Level:     3,
		Parent:    "Missile Launcher",
		DamageMin: 20,
		DamageMax: 20,
	},
	144: {
		Name:      "Missile Launcher",
		Level:     3,
		DamageMin: 20,
		DamageMax: 21,
	},
	145: {
		Name:      "Enhanced Missile Launcher",
		Level:     3,
		Parent:    "Missile Launcher",
		DamageMin: 20,
		DamageMax: 22,
	},
	146: {
		Name:      "Hardened Missile Launcher",
		Level:     3,
		Parent:    "Missile Launcher",
		DamageMin: 20,
		DamageMax: 23,
	},
	147: {
		Name:      "Guided Missile Launcher",
		Level:     3,
		Parent:    "Missile Launcher",
		DamageMin: 20,
		DamageMax: 24,
	},
	148: {
		Name:      "Miss Launcher",
		Level:     3,
		Parent:    "Missile Launcher",
		DamageMin: 20,
		DamageMax: 25,
	},

	// Laser Musket
	149: {
		Name:      "Laser Musket",
		Level:     3,
		DamageMin: 10,
		DamageMax: 13,
	},
}
