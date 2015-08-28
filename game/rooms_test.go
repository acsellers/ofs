package game

import "testing"

func TestRoomCapacity(t *testing.T) {
	testRooms := []struct {
		ID       Room
		Level    int
		Size     int
		Expected int
	}{
		// living quarters
		{2, 1, 1, 8},
		{2, 1, 2, 18},
		{2, 1, 3, 28},
		{2, 2, 1, 10},
		{2, 2, 2, 22},
		{2, 2, 3, 34},
		{2, 3, 1, 12},
		{2, 3, 2, 26},
		{2, 3, 3, 40},

		// storage room
		{3, 1, 1, 10},
		{3, 1, 2, 20},
		{3, 1, 3, 30},
		{3, 2, 1, 15},
		{3, 2, 2, 30},
		{3, 2, 3, 45},
		{3, 3, 1, 20},
		{3, 3, 2, 40},
		{3, 3, 3, 60},

		// diner
		{4, 1, 1, 50},
		{4, 1, 2, 100},
		{4, 1, 3, 150},
		{4, 2, 1, 75},
		{4, 2, 2, 150},
		{4, 2, 3, 225},
		{4, 3, 1, 100},
		{4, 3, 2, 200},
		{4, 3, 3, 300},

		// garden
		{5, 1, 1, 60},
		{5, 1, 2, 120},
		{5, 1, 3, 180},
		{5, 2, 1, 90},
		{5, 2, 2, 180},
		{5, 2, 3, 270},
		{5, 3, 1, 120},
		{5, 3, 2, 240},
		{5, 3, 3, 360},

		// power generator
		{6, 1, 1, 75},
		{6, 1, 2, 150},
		{6, 1, 3, 225},
		{6, 2, 1, 112},
		{6, 2, 2, 224},
		{6, 2, 3, 336},
		{6, 3, 1, 150},
		{6, 3, 2, 300},
		{6, 3, 3, 450},

		// nuke-you-ler reactor
		{7, 1, 1, 200},
		{7, 1, 2, 400},
		{7, 1, 3, 600},
		{7, 2, 1, 300},
		{7, 2, 2, 600},
		{7, 2, 3, 900},
		{7, 3, 1, 400},
		{7, 3, 2, 800},
		{7, 3, 3, 1200},
	}

	for _, test := range testRooms {
		cp := Rooms[test.ID].Capacity(test.Level, test.Size)
		if cp != test.Expected {
			t.Fatalf(
				"%s Capacity (L%d, S%d): Expected %d, got %d",
				Rooms[test.ID].Name(test.Level),
				test.Level,
				test.Size,
				test.Expected,
				cp,
			)
		}
	}
}

func TestRoomCost(t *testing.T) {
	testRooms := []struct {
		ID       Room
		Level    int
		Size     int
		Expected int
	}{
		// living quarters
		{2, 1, 1, 250},
		{2, 1, 2, 375},
		{2, 1, 3, 500},
		{2, 2, 1, 750},
		{2, 2, 2, 1125},
		{2, 2, 3, 1500},

		// storage room
		{3, 1, 1, 750},
		{3, 1, 2, 1125},
		{3, 1, 3, 1500},
		{3, 2, 1, 2250},
		{3, 2, 2, 3375},
		{3, 2, 3, 4500},

		// diner
		{4, 1, 1, 250},
		{4, 1, 2, 375},
		{4, 1, 3, 500},
		{4, 2, 1, 750},
		{4, 2, 2, 1125},
		{4, 2, 3, 1500},

		// garden
		{5, 1, 1, 3000},
		{5, 1, 2, 4500},
		{5, 1, 3, 6000},
		{5, 2, 1, 9000},
		{5, 2, 2, 13500},
		{5, 2, 3, 18000},

		// power generator
		{6, 1, 1, 250},
		{6, 1, 2, 375},
		{6, 1, 3, 500},
		{6, 2, 1, 750},
		{6, 2, 2, 1125},
		{6, 2, 3, 1500},

		// nucleur reactor
		{7, 1, 1, 3000},
		{7, 1, 2, 4500},
		{7, 1, 3, 6000},
		{7, 2, 1, 9000},
		{7, 2, 2, 13500},
		{7, 2, 3, 18000},
	}

	for _, test := range testRooms {
		cp := Rooms[test.ID].UpgradeCost(test.Level, test.Size)
		if cp != test.Expected {
			t.Fatalf(
				"%s UpgradeCost (L%d, S%d): Expected %d, got %d",
				Rooms[test.ID].Name(test.Level),
				test.Level,
				test.Size,
				test.Expected,
				cp,
			)
		}
	}
}

func TestRoomProduction(t *testing.T) {
	testRooms := []struct {
		ID       Room
		Level    int
		Size     int
		Expected int
	}{
		// diner
		{4, 1, 1, 8},
		{4, 1, 2, 18},
		{4, 1, 3, 28},
		{4, 2, 1, 10},
		{4, 2, 2, 22},
		{4, 2, 3, 34},
		{4, 3, 1, 12},
		{4, 3, 2, 26},
		{4, 3, 3, 40},

		// garden
		{5, 1, 1, 10},
		{5, 1, 2, 23},
		{5, 1, 3, 36},
		{5, 2, 1, 13},
		{5, 2, 2, 29},
		{5, 2, 3, 45},
		{5, 3, 1, 16},
		{5, 3, 2, 35},
		{5, 3, 3, 54},

		// power generator
		{6, 1, 1, 14},
		{6, 1, 2, 30},
		{6, 1, 3, 46},
		{6, 2, 1, 17},
		{6, 2, 2, 36},
		{6, 2, 3, 55},
		{6, 3, 1, 20},
		{6, 3, 2, 42},
		{6, 3, 3, 64},

		// nucleur reactor
		{7, 1, 1, 20},
		{7, 1, 2, 43},
		{7, 1, 3, 66},
		{7, 2, 1, 24},
		{7, 2, 2, 51},
		{7, 2, 3, 78},
		{7, 3, 1, 28},
		{7, 3, 2, 59},
		{7, 3, 3, 90},
	}

	for _, test := range testRooms {
		cp := Rooms[test.ID].Production(test.Level, test.Size)
		if cp != test.Expected {
			t.Fatalf(
				"%s Production (L%d, S%d): Expected %d, got %d",
				Rooms[test.ID].Name(test.Level),
				test.Level,
				test.Size,
				test.Expected,
				cp,
			)
		}
	}
}
