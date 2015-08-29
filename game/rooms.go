package game

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
	Storage     [3][3]int
	Yield       [3][3]int
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

type Storage int

const (
	People Storage = iota
	Items
	Food
	Water
	Power
	Health
	Radiation
)

func (rs RoomStat) Stores() []Storage {
	switch rs.Category {
	case "People":
		return []Storage{People}
	case "Storage":
		return []Storage{
			Items,
			Health,
			Radiation,
		}
	case "Food":
		return []Storage{Food}
	case "Water":
		return []Storage{Water}
	case "Power":
		return []Storage{Power}
	case "Health":
		return []Storage{Health}
	case "Radiation":
		return []Storage{Radiation}
	}
	return []Storage{}
}

func (rs RoomStat) Produces() []Storage {
	switch rs.Category {
	case "Food":
		return []Storage{Food}
	case "Water":
		return []Storage{Water}
	case "Food+Water":
		return []Storage{Food, Water}
	case "Power":
		return []Storage{Power}
	case "Health":
		return []Storage{Health}
	case "Radiation":
		return []Storage{Radiation}
	}
	return []Storage{}
}

func (rs RoomStat) Capacity(level, size int) int {
	return rs.Storage[level-1][size-1]
}

func (rs RoomStat) Production(level, size int) int {
	return rs.Yield[level-1][size-1]
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
		Storage: [3][3]int{
			{8, 18, 32},
			{10, 24, 38},
			{12, 28, 46},
		},
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
		Storage: [3][3]int{
			{10, 20, 30},
			{15, 30, 45},
			{20, 40, 60},
		},
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
		Yield: [3][3]int{
			{8, 18, 30},
			{10, 22, 36},
			{12, 26, 42},
		},
		Storage: [3][3]int{
			{50, 100, 150},
			{75, 150, 225},
			{100, 200, 300},
		},
		Tier: 1,
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
		Yield: [3][3]int{
			{10, 23, 36},
			{13, 29, 45},
			{16, 35, 54},
		},
		Storage: [3][3]int{
			{60, 120, 180},
			{90, 180, 270},
			{120, 240, 360},
		},
		Tier:        2,
		MinDwellers: 70,
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
		Yield: [3][3]int{
			{14, 30, 44},
			{17, 36, 55},
			{20, 42, 66},
		},
		Storage: [3][3]int{
			{75, 150, 225},
			{112, 224, 336},
			{150, 300, 450},
		},
		Tier: 1,
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
		Yield: [3][3]int{
			{20, 44, 66},
			{24, 55, 77},
			{28, 66, 88},
		},
		Storage: [3][3]int{
			{200, 400, 600},
			{300, 600, 900},
			{400, 800, 1200},
		},
		Tier:        2,
		MinDwellers: 60,
	},
	// water production
	8: {
		Lv1Name:  "Water1U1",
		Lv2Name:  "Water1U2",
		Lv3Name:  "Water1U3",
		Category: "Water",
		CostBase: 100,
		CostAdd:  25,
		CostLv2:  250,
		CostLv3:  750,
		Yield: [3][3]int{
			{8, 18, 30},
			{10, 22, 36},
			{12, 26, 42},
		},
		Storage: [3][3]int{
			{8, 18, 32},
			{10, 24, 38},
			{12, 28, 46},
		},
		Tier: 1,
	},
	9: {
		Lv1Name:  "Water2U1",
		Lv2Name:  "Water2U2",
		Lv3Name:  "Water2U3",
		Category: "Water",
		CostBase: 1200,
		CostAdd:  300,
		CostLv2:  3000,
		CostLv3:  9000,
		Yield: [3][3]int{
			{10, 23, 36},
			{13, 29, 45},
			{16, 35, 54},
		},
		Storage: [3][3]int{
			{50, 100, 150},
			{75, 150, 225},
			{100, 200, 300},
		},
		Tier:        2,
		MinDwellers: 80,
	},
	// medical
	10: {
		Lv1Name:  "Health1U1",
		Lv2Name:  "Health1U2",
		Lv3Name:  "Health1U3",
		Category: "Health",
		CostBase: 400,
		CostAdd:  100,
		CostLv2:  1000,
		CostLv3:  3000,
		Yield: [3][3]int{
			{3, 6, 9},
			{4, 8, 12},
			{6, 12, 18},
		},
		Storage: [3][3]int{
			{10, 20, 30},
			{12, 24, 36},
			{15, 30, 45},
		},
		Tier:        1,
		MinDwellers: 14,
	},
	11: {
		Lv1Name:  "Radiation1U1",
		Lv2Name:  "Radiation1U2",
		Lv3Name:  "Radiation1U3",
		Category: "Radiation",
		CostBase: 400,
		CostAdd:  100,
		CostLv2:  1000,
		CostLv3:  3000,
		Yield: [3][3]int{
			{3, 6, 9},
			{4, 8, 12},
			{6, 12, 18},
		},
		Storage: [3][3]int{
			{10, 20, 30},
			{12, 24, 36},
			{15, 30, 45},
		},
		Tier:        1,
		MinDwellers: 16,
	},
	// food+water
	12: {
		Lv1Name:  "Meal1U1",
		Lv2Name:  "Meal1U2",
		Lv3Name:  "Meal1U3",
		Category: "Food+Water",
		CostBase: 2500,
		CostAdd:  500,
		CostLv2:  4000,
		CostLv3:  10000,
		Yield: [3][3]int{
			{8, 18, 30},
			{10, 22, 36},
			{12, 26, 42},
		},
		Storage: [3][3]int{
			{50, 100, 150},
			{75, 150, 225},
			{100, 200, 300},
		},
		Tier: 1,
	},
}
