package game

import "fmt"

type Room int

type RoomStat struct {
	Lv1Name     string
	Lv2Name     string
	Lv3Name     string
	Tier        int
	Category    string
	Ability     string
	MinDwellers int
	CostBase    int
	CostAdd     int
	CostLv2     int
	CostLv3     int
	Storage     int
	Yield       int
	Bonus       int // production bonus
}

func (rs RoomStat) Name(level int) string {
	switch level {
	case 1:
		return rs.Lv1Name
	case 2:
		if rs.Lv2Name != "" {
			return rs.Lv2Name
		}
	case 3:
		if rs.Lv3Name != "" {
			return rs.Lv3Name
		}
	}
	return rs.Lv1Name
}

func (rs RoomStat) Capacity(level, size int) int {
	switch rs.Category {
	case "People":
		baseCap := rs.Storage + 2*(level-1)
		return baseCap*size + 2*(size-1)
	case "Storage":
		return rs.Storage*size + 5*(level-1)*size
	case "Food", "Power":
		baseCap := rs.Storage + (level-1)*rs.Storage/2
		return baseCap * size

	}
	fmt.Println("Missing room:", rs.Lv1Name)
	return 0
}

func (rs RoomStat) Production(level, size int) int {
	if rs.Yield == 0 {
		return 0
	}
	kicker := 1 + rs.Tier

	baseCap := rs.Yield + kicker*(level-1)
	if rs.Bonus > 0 {
		baseCap += (level - 1) ^ rs.Bonus
	}
	return baseCap*size + kicker*(size-1)
}

// UpgradeCost computes the cost to upgrade a room of size provided
// from provided level to the next level.
func (rs RoomStat) UpgradeCost(level, size int) int {
	var cost int
	switch level {
	case 1:
		cost = rs.CostLv2
	case 2:
		cost = rs.CostLv3
	default:
		return 0
	}

	// vault door
	if rs.Category == "" {
		return cost
	}

	return cost + cost*(size-1)/2
}

var Rooms = map[Room]RoomStat{
	-1: {
		Lv1Name: "Vault Door",
		CostLv2: 500,
		CostLv3: 2000,
	},
	0: {
		Lv1Name: "Dirt",
	},
	1: {
		Lv1Name:  "Elevator",
		CostBase: 100,
		CostAdd:  25,
	},
	// storage
	2: {
		Lv1Name:  "PeopleU1",
		Lv2Name:  "PeopleU2",
		Lv3Name:  "PeopleU3",
		Category: "People",
		CostBase: 100,
		CostAdd:  25,
		CostLv2:  250,
		CostLv3:  750,
		Storage:  8,
	},
	3: {
		Lv1Name:     "StorageU1",
		Lv2Name:     "StorageU2",
		Lv3Name:     "StorageU3",
		Category:    "Storage",
		MinDwellers: 12,
		CostBase:    100,
		CostAdd:     25,
		CostLv2:     750,
		CostLv3:     2250,
		Storage:     10,
	},
	// food production
	4: {
		Lv1Name:  "Food1U1",
		Lv2Name:  "Food1U2",
		Lv3Name:  "Food1U3",
		Category: "Food",
		CostBase: 100,
		CostAdd:  25,
		CostLv2:  250,
		CostLv3:  750,
		Yield:    8,
		Storage:  50,
		Tier:     1,
	},
	5: {
		Lv1Name:  "Food2U1",
		Lv2Name:  "Food2U2",
		Lv3Name:  "Food2U3",
		Category: "Food",
		CostBase: 1200,
		CostAdd:  300,
		CostLv2:  3000,
		CostLv3:  9000,
		Yield:    10,
		Storage:  60,
		Tier:     2,
	},
	// power production
	6: {
		Lv1Name:  "Power1U1",
		Lv2Name:  "Power1U2",
		Lv3Name:  "Power1U3",
		Category: "Power",
		CostBase: 100,
		CostAdd:  25,
		CostLv2:  250,
		CostLv3:  750,
		Yield:    10,
		Storage:  75,
		Tier:     1,
		Bonus:    4,
	},
	7: {
		Lv1Name:  "Power2U1",
		Lv2Name:  "Power2U2",
		Lv3Name:  "Power2U3",
		Category: "Power",
		CostBase: 1200,
		CostAdd:  300,
		CostLv2:  3000,
		CostLv3:  9000,
		Yield:    16,
		Storage:  200,
		Tier:     2,
		Bonus:    4,
	},
}
