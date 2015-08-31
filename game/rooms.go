package game

import "strings"

type Room int

type RoomStat struct {
	Code          string
	Lv1Name       string
	Lv2Name       string
	Lv3Name       string
	Tier          int
	Category      string
	Ability       string
	MinDwellers   int
	CostBase      int
	CostAdd       int
	Upgrade1Price [2]int
	Upgrade2Price [2]int
	CostLv2       int
	CostLv3       int
	Storage       [3][3]int
	Yield         [3][3]int
}

func (rs RoomStat) Name(level int) string {
	switch level {
	case 1:
		if rs.Lv1Name != "" {
			return rs.Lv1Name
		}
	case 2:
		if rs.Lv2Name != "" {
			return rs.Lv2Name
		}
	case 3:
		if rs.Lv3Name != "" {
			return rs.Lv3Name
		}
	}
	return rs.Code
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

func (rs RoomStat) Stores(level1, level2, size int) map[Storage]int {
	stores := make(map[Storage]int)
	if strings.Contains(rs.Code, "People") {
		stores[People] = rs.Storage[level1-1][size-1]
	}
	if strings.Contains(rs.Code, "Warehouse") {
		stores[Items] = rs.Storage[level1-1][size-1]
		stores[Health] = rs.Storage[level1-1][size-1] / 2
		stores[Radiation] = rs.Storage[level1-1][size-1] / 2
	}
	if strings.Contains(rs.Code, "Food") {
		stores[Food] = rs.Storage[level2-1][size-1]
	}
	if strings.Contains(rs.Code, "Water") {
		stores[Water] = rs.Storage[level2-1][size-1]
	}
	if strings.Contains(rs.Code, "Power") {
		stores[Power] = rs.Storage[level2-1][size-1]
	}
	if strings.Contains(rs.Code, "Health") {
		stores[Health] = rs.Storage[level2-1][size-1]
	}
	if strings.Contains(rs.Code, "Radiation") {
		stores[Radiation] = rs.Storage[level2-1][size-1]
	}
	return stores
}

func (rs RoomStat) Produces(level1, size int) map[Storage]int {
	produce := make(map[Storage]int)
	if strings.Contains(rs.Code, "Food") {
		produce[Food] = rs.Yield[level1-1][size-1]
	}
	if strings.Contains(rs.Code, "Water") {
		produce[Water] = rs.Yield[level1-1][size-1]
	}
	if strings.Contains(rs.Code, "Power") {
		produce[Power] = rs.Yield[level1-1][size-1]
	}
	if strings.Contains(rs.Code, "Health") {
		produce[Health] = rs.Yield[level1-1][size-1]
	}
	if strings.Contains(rs.Code, "Radiation") {
		produce[Radiation] = rs.Yield[level1-1][size-1]
	}

	return produce
}

func (rs RoomStat) Capacity(level1, level2, size int) int {
	switch rs.Category {
	case "Production", "Medical":
		return rs.Storage[level2-1][size-1]
	case "Storage":
		return rs.Storage[level1-1][size-1]
	}
	return 0
}

func (rs RoomStat) Production(level1, level2, size int) int {
	return rs.Yield[level1-1][size-1]
}

func (rs RoomStat) Upgrade1Cost(level1, size int) int {
	switch size {
	case 1:
		return rs.Upgrade1Price[level1-1]
	case 2:
		return rs.Upgrade1Price[level1-1] + rs.Upgrade1Price[level1-1]/2
	case 3:
		return 2 * rs.Upgrade1Price[level1-1]
	}
	return 0
}

func (rs RoomStat) Upgrade2Cost(level2, size int) int {
	switch size {
	case 1:
		return rs.Upgrade2Price[level2-1]
	case 2:
		return rs.Upgrade2Price[level2-1] + rs.Upgrade2Price[level2-1]/2
	case 3:
		return 2 * rs.Upgrade2Price[level2-1]
	}
	return 0
}

func RoomByCode(code string) (Room, RoomStat) {
	for r, rs := range Rooms {
		if rs.Code == code {
			return r, rs
		}
	}
	return -1, RoomStat{}
}

var Rooms = map[Room]RoomStat{
	// special
	0: {
		Code:     "Dirt",
		Category: "Special",
	},
	1: {
		Code:     "Elevator",
		Category: "Special",
		CostBase: 100,
		CostAdd:  25,
	},

	// food production
	4: {
		Code:          "Food1",
		Category:      "Production",
		CostBase:      100,
		CostAdd:       25,
		Upgrade1Price: [2]int{200, 600},
		Upgrade2Price: [2]int{100, 400},
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
		Code:          "Food2",
		Category:      "Production",
		CostBase:      1200,
		CostAdd:       300,
		Upgrade1Price: [2]int{2000, 6000},
		Upgrade2Price: [2]int{1000, 4000},

		CostLv2: 3000,
		CostLv3: 9000,
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
		MinDwellers: 30,
	},
	// water production
	8: {
		Code:          "Water1",
		Category:      "Production",
		CostBase:      100,
		CostAdd:       25,
		Upgrade1Price: [2]int{200, 600},
		Upgrade2Price: [2]int{100, 400},
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
	9: {
		Code:          "Water2",
		Category:      "Production",
		CostBase:      1200,
		CostAdd:       300,
		Upgrade1Price: [2]int{2000, 6000},
		Upgrade2Price: [2]int{1000, 4000},
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
		MinDwellers: 80,
	},
	// food+water
	12: {
		Code:        "Food+Water",
		Category:    "Production",
		CostBase:    2500,
		CostAdd:     500,
		Tier:        1,
		MinDwellers: 75,

		Upgrade1Price: [2]int{3000, 9000},
		Upgrade2Price: [2]int{2000, 4000},
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
	},
	// power production
	6: {
		Code:          "Power1",
		Category:      "Production",
		CostBase:      100,
		CostAdd:       25,
		Upgrade1Price: [2]int{160, 500},
		Upgrade2Price: [2]int{160, 500},
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
		Code:          "Power2",
		Category:      "Production",
		CostBase:      1200,
		CostAdd:       300,
		Upgrade1Price: [2]int{1600, 5000},
		Upgrade2Price: [2]int{1600, 5000},
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
		MinDwellers: 35,
	},

	// storage
	2: {
		Code:          "People",
		Category:      "Storage",
		CostBase:      100,
		CostAdd:       25,
		Upgrade1Price: [2]int{250, 750},
		Upgrade2Price: [2]int{2000, 6000},
		Storage: [3][3]int{
			{8, 18, 32},
			{10, 24, 38},
			{12, 28, 46},
		},
	},
	3: {
		Code:          "Warehouse",
		Category:      "Storage",
		MinDwellers:   12,
		CostBase:      100,
		CostAdd:       50,
		Upgrade1Price: [2]int{750, 2250},
		Upgrade2Price: [2]int{1000, 4000},

		Storage: [3][3]int{
			{10, 20, 30},
			{15, 30, 45},
			{20, 40, 60},
		},
	},
	// medical
	10: {
		Code:        "Health",
		Category:    "Medical",
		CostBase:    400,
		CostAdd:     100,
		Tier:        1,
		MinDwellers: 20,

		Upgrade1Price: [2]int{1000, 3000},
		Upgrade2Price: [2]int{500, 1500},
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
	},
	11: {
		Code:        "Radiation",
		Category:    "Medical",
		CostBase:    400,
		CostAdd:     100,
		Tier:        1,
		MinDwellers: 20,

		Upgrade1Price: [2]int{1000, 3000},
		Upgrade2Price: [2]int{500, 1500},
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
	},

	// service rooms
	13: {
		Code:          "Repair",
		Category:      "Misc",
		CostBase:      5000,
		CostAdd:       1000,
		Upgrade1Price: [2]int{},
		Upgrade2Price: [2]int{},
		CostLv2:       4000,
		CostLv3:       10000,
		Tier:          1,
		MinDwellers:   30,
	},
	14: {
		Code:     "Radio",
		Category: "Misc",
		CostBase: 4000,
		CostAdd:  0,
		CostLv2:  4000,
		CostLv3:  10000,
		Tier:     1,
	},
	15: {
		Code:        "Office",
		Category:    "Misc",
		CostBase:    20000,
		CostAdd:     500,
		CostLv2:     4000,
		CostLv3:     10000,
		Tier:        1,
		MinDwellers: 60,
	},
	16: {
		Code:     "Entrance",
		Category: "Misc",
		CostLv2:  500,
		CostLv3:  2000,
	},
	17: {
		Code:        "Trader",
		Category:    "Misc",
		CostBase:    10000,
		CostAdd:     0,
		CostLv2:     4000,
		CostLv3:     10000,
		Tier:        1,
		MinDwellers: 120,
	},
	18: {
		Code:        "Hanger",
		Category:    "Misc",
		CostBase:    100000,
		CostAdd:     0,
		CostLv2:     40000,
		CostLv3:     150000,
		Tier:        1,
		MinDwellers: 200,
	},
}
