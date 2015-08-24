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
