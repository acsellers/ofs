package game

import "testing"

func TestRoomCapacity(t *testing.T) {
	testRooms := []struct {
		ID       Room
		Level    int
		Size     int
		Expected int
		Type     Storage
	}{
		// food 1
		{4, 1, 1, 50, Food},
		{4, 1, 2, 100, Food},
		{4, 1, 3, 150, Food},
		{4, 2, 1, 75, Food},
		{4, 2, 2, 150, Food},
		{4, 2, 3, 225, Food},
		{4, 3, 1, 100, Food},
		{4, 3, 2, 200, Food},
		{4, 3, 3, 300, Food},

		// food2
		{5, 1, 1, 60, Food},
		{5, 1, 2, 120, Food},
		{5, 1, 3, 180, Food},
		{5, 2, 1, 90, Food},
		{5, 2, 2, 180, Food},
		{5, 2, 3, 270, Food},
		{5, 3, 1, 120, Food},
		{5, 3, 2, 240, Food},
		{5, 3, 3, 360, Food},

		// water1
		{8, 1, 1, 50, Water},
		{8, 1, 2, 100, Water},
		{8, 1, 3, 150, Water},
		{8, 2, 1, 75, Water},
		{8, 2, 2, 150, Water},
		{8, 2, 3, 225, Water},
		{8, 3, 1, 100, Water},
		{8, 3, 2, 200, Water},
		{8, 3, 3, 300, Water},

		// food2
		{9, 1, 1, 60, Water},
		{9, 1, 2, 120, Water},
		{9, 1, 3, 180, Water},
		{9, 2, 1, 90, Water},
		{9, 2, 2, 180, Water},
		{9, 2, 3, 270, Water},
		{9, 3, 1, 120, Water},
		{9, 3, 2, 240, Water},
		{9, 3, 3, 360, Water},

		// food+water
		{12, 1, 1, 50, Food},
		{12, 1, 2, 100, Food},
		{12, 1, 3, 150, Food},
		{12, 2, 1, 75, Food},
		{12, 2, 2, 150, Food},
		{12, 2, 3, 225, Food},
		{12, 3, 1, 100, Food},
		{12, 3, 2, 200, Food},
		{12, 3, 3, 300, Food},
		{12, 1, 1, 50, Water},
		{12, 1, 2, 100, Water},
		{12, 1, 3, 150, Water},
		{12, 2, 1, 75, Water},
		{12, 2, 2, 150, Water},
		{12, 2, 3, 225, Water},
		{12, 3, 1, 100, Water},
		{12, 3, 2, 200, Water},
		{12, 3, 3, 300, Water},

		// power1
		{6, 1, 1, 75, Power},
		{6, 1, 2, 150, Power},
		{6, 1, 3, 225, Power},
		{6, 2, 1, 112, Power},
		{6, 2, 2, 224, Power},
		{6, 2, 3, 336, Power},
		{6, 3, 1, 150, Power},
		{6, 3, 2, 300, Power},
		{6, 3, 3, 450, Power},

		// power2
		{7, 1, 1, 200, Power},
		{7, 1, 2, 400, Power},
		{7, 1, 3, 600, Power},
		{7, 2, 1, 300, Power},
		{7, 2, 2, 600, Power},
		{7, 2, 3, 900, Power},
		{7, 3, 1, 400, Power},
		{7, 3, 2, 800, Power},
		{7, 3, 3, 1200, Power},

		// people
		{2, 1, 1, 8, People},
		{2, 1, 2, 18, People},
		{2, 1, 3, 32, People},
		{2, 2, 1, 10, People},
		{2, 2, 2, 24, People},
		{2, 2, 3, 38, People},
		{2, 3, 1, 12, People},
		{2, 3, 2, 28, People},
		{2, 3, 3, 46, People},

		// warehouse
		{3, 1, 1, 10, Items},
		{3, 1, 2, 20, Items},
		{3, 1, 3, 30, Items},
		{3, 2, 1, 15, Items},
		{3, 2, 2, 30, Items},
		{3, 2, 3, 45, Items},
		{3, 3, 1, 20, Items},
		{3, 3, 2, 40, Items},
		{3, 3, 3, 60, Items},

		{3, 1, 1, 5, Health},
		{3, 1, 2, 10, Radiation},
		{3, 1, 3, 15, Health},
		{3, 2, 1, 7, Radiation},
		{3, 2, 2, 15, Health},
		{3, 2, 3, 22, Radiation},
		{3, 3, 1, 10, Health},
		{3, 3, 2, 20, Radiation},
		{3, 3, 3, 30, Health},
	}

	for _, test := range testRooms {
		cp := Rooms[test.ID].Stores(test.Level, test.Level, test.Size)
		if cp[test.Type] != test.Expected {
			t.Fatalf(
				"%s Stores (L%d, S%d, T%d): Expected %d, got %d",
				Rooms[test.ID].Name(test.Level),
				test.Level,
				test.Size,
				test.Type,
				test.Expected,
				cp,
			)
		}
	}
}

func TestRoomCost(t *testing.T) {
	testRooms := []struct {
		ID        Room
		Level     int
		Size      int
		Expected1 int
		Expected2 int
	}{
		// living quarters
		{2, 1, 1, 250, 2000},
		{2, 1, 2, 375, 3000},
		{2, 1, 3, 500, 4000},
		{2, 2, 1, 750, 6000},
		{2, 2, 2, 1125, 9000},
		{2, 2, 3, 1500, 12000},

		// storage room
		{3, 1, 1, 750, 1000},
		{3, 1, 2, 1125, 1500},
		{3, 1, 3, 1500, 2000},
		{3, 2, 1, 2250, 4000},
		{3, 2, 2, 3375, 6000},
		{3, 2, 3, 4500, 8000},

		// food1
		{4, 1, 1, 200, 100},
		{4, 1, 2, 300, 150},
		{4, 1, 3, 400, 200},
		{4, 2, 1, 600, 400},
		{4, 2, 2, 900, 600},
		{4, 2, 3, 1200, 800},

		// food2
		{5, 1, 1, 2000, 1000},
		{5, 1, 2, 3000, 1500},
		{5, 1, 3, 4000, 2000},
		{5, 2, 1, 6000, 4000},
		{5, 2, 2, 9000, 6000},
		{5, 2, 3, 12000, 8000},

		// power1
		{6, 1, 1, 160, 160},
		{6, 1, 2, 240, 240},
		{6, 1, 3, 320, 320},
		{6, 2, 1, 500, 500},
		{6, 2, 2, 750, 750},
		{6, 2, 3, 1000, 1000},

		// power2
		{7, 1, 1, 1600, 1600},
		{7, 1, 2, 2400, 2400},
		{7, 1, 3, 3200, 3200},
		{7, 2, 1, 5000, 5000},
		{7, 2, 2, 7500, 7500},
		{7, 2, 3, 10000, 10000},

		// water1
		{8, 1, 1, 200, 100},
		{8, 1, 2, 300, 150},
		{8, 1, 3, 400, 200},
		{8, 2, 1, 600, 400},
		{8, 2, 2, 900, 600},
		{8, 2, 3, 1200, 800},

		// water2
		{9, 1, 1, 2000, 1000},
		{9, 1, 2, 3000, 1500},
		{9, 1, 3, 4000, 2000},
		{9, 2, 1, 6000, 4000},
		{9, 2, 2, 9000, 6000},
		{9, 2, 3, 12000, 8000},

		// food+water
		{12, 1, 1, 3000, 2000},
		{12, 1, 2, 4500, 3000},
		{12, 1, 3, 6000, 4000},
		{12, 2, 1, 9000, 4000},
		{12, 2, 2, 13500, 6000},
		{12, 2, 3, 18000, 8000},

		// health
		{10, 1, 1, 1000, 500},
		{10, 1, 2, 1500, 750},
		{10, 1, 3, 2000, 1000},
		{10, 2, 1, 3000, 1500},
		{10, 2, 2, 4500, 2250},
		{10, 2, 3, 6000, 3000},

		// radiation
		{11, 1, 1, 1000, 500},
		{11, 1, 2, 1500, 750},
		{11, 1, 3, 2000, 1000},
		{11, 2, 1, 3000, 1500},
		{11, 2, 2, 4500, 2250},
		{11, 2, 3, 6000, 3000},

		// repair
		// radio
		// office
		// entrance
		// strength
		// perception
		// endurance
		// charisma
		// intelligence
		// agility
		// luck
		// flame trap
		// drop trap
		// gun trap
		// gas trap
	}

	for _, test := range testRooms {
		p1 := Rooms[test.ID].Upgrade1Cost(test.Level, test.Size)
		if p1 != test.Expected1 {
			t.Fatalf(
				"%s Upgrade1Cost (L%d, S%d): Expected %d, got %d",
				Rooms[test.ID].Name(test.Level),
				test.Level,
				test.Size,
				test.Expected1,
				p1,
			)
		}

		p2 := Rooms[test.ID].Upgrade2Cost(test.Level, test.Size)
		if p2 != test.Expected2 {
			t.Fatalf(
				"%s Upgrade2Cost (L%d, S%d): Expected %d, got %d",
				Rooms[test.ID].Name(test.Level),
				test.Level,
				test.Size,
				test.Expected2,
				p2,
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
		Type     Storage
	}{
		// food1
		{4, 1, 1, 8, Food},
		{4, 1, 2, 18, Food},
		{4, 1, 3, 30, Food},
		{4, 2, 1, 10, Food},
		{4, 2, 2, 22, Food},
		{4, 2, 3, 36, Food},
		{4, 3, 1, 12, Food},
		{4, 3, 2, 26, Food},
		{4, 3, 3, 42, Food},

		// food2
		{5, 1, 1, 10, Food},
		{5, 1, 2, 23, Food},
		{5, 1, 3, 36, Food},
		{5, 2, 1, 13, Food},
		{5, 2, 2, 29, Food},
		{5, 2, 3, 45, Food},
		{5, 3, 1, 16, Food},
		{5, 3, 2, 35, Food},
		{5, 3, 3, 54, Food},

		// water1
		{8, 1, 1, 8, Water},
		{8, 1, 2, 18, Water},
		{8, 1, 3, 30, Water},
		{8, 2, 1, 10, Water},
		{8, 2, 2, 22, Water},
		{8, 2, 3, 36, Water},
		{8, 3, 1, 12, Water},
		{8, 3, 2, 26, Water},
		{8, 3, 3, 42, Water},

		// water2
		{9, 1, 1, 10, Water},
		{9, 1, 2, 23, Water},
		{9, 1, 3, 36, Water},
		{9, 2, 1, 13, Water},
		{9, 2, 2, 29, Water},
		{9, 2, 3, 45, Water},
		{9, 3, 1, 16, Water},
		{9, 3, 2, 35, Water},
		{9, 3, 3, 54, Water},

		// food+water
		{12, 1, 1, 8, Food},
		{12, 1, 2, 18, Food},
		{12, 1, 3, 30, Food},
		{12, 2, 1, 10, Food},
		{12, 2, 2, 22, Food},
		{12, 2, 3, 36, Food},
		{12, 3, 1, 12, Food},
		{12, 3, 2, 26, Food},
		{12, 3, 3, 42, Food},

		{12, 1, 1, 8, Water},
		{12, 1, 2, 18, Water},
		{12, 1, 3, 30, Water},
		{12, 2, 1, 10, Water},
		{12, 2, 2, 22, Water},
		{12, 2, 3, 36, Water},
		{12, 3, 1, 12, Water},
		{12, 3, 2, 26, Water},
		{12, 3, 3, 42, Water},

		// power1
		{6, 1, 1, 14, Power},
		{6, 1, 2, 30, Power},
		{6, 1, 3, 44, Power},
		{6, 2, 1, 17, Power},
		{6, 2, 2, 36, Power},
		{6, 2, 3, 55, Power},
		{6, 3, 1, 20, Power},
		{6, 3, 2, 42, Power},
		{6, 3, 3, 66, Power},

		// power2
		{7, 1, 1, 20, Power},
		{7, 1, 2, 44, Power},
		{7, 1, 3, 66, Power},
		{7, 2, 1, 24, Power},
		{7, 2, 2, 55, Power},
		{7, 2, 3, 77, Power},
		{7, 3, 1, 28, Power},
		{7, 3, 2, 66, Power},
		{7, 3, 3, 88, Power},
		// health
		// radiation
	}

	for _, test := range testRooms {
		cp := Rooms[test.ID].Produces(test.Level, test.Size)
		if cp[test.Type] != test.Expected {
			t.Fatalf(
				"%s Production (L%d, S%d, T%d): Expected %d, got %d",
				Rooms[test.ID].Name(test.Level),
				test.Level,
				test.Size,
				test.Type,
				test.Expected,
				cp,
			)
		}
	}
}
