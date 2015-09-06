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

func TestClothesBalance(t *testing.T) {
	for _, c := range Clothes {
		capsum := c.Stats.Charm +
			c.Stats.Awareness +
			c.Stats.Power +
			c.Stats.Success +
			c.Stats.Tenacity +
			c.Stats.Outgoing +
			c.Stats.Nimble +
			c.Stats.Erudite

		if capsum != c.Tier*3 {
			t.Fatalf("Bad stats for clothing %s: %d", c.Name, capsum)
		}
	}
}
