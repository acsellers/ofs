package game

import "testing"

func TestItemIds(t *testing.T) {
	for k, c := range Clothes {
		if int(k) != int(c.ID) {
			t.Fatal("%s Mismatched ID: %d vs %d", c.Name, k, c.ID)
		}
	}

	for k, c := range Weapons {
		if int(k)+1000 != int(c.ID) {
			t.Fatal("%s Mismatched ID: %d vs %d", c.Name, k, c.ID)
		}
	}

}
