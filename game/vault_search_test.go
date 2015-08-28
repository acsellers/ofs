package game

import "testing"

func TestVaultRouting(t *testing.T) {
	v := NewVault()
	v.PlaceRoom(10, 0, 1)
	v.PlaceRoom(10, 1, 1)
	to := v.PlaceRoom(11, 1, 2)
	path := v.Route(1, to)
	if len(path) == 0 {
		t.Fatal("No Path?")
	}
}

func TestRoomSprite(t *testing.T) {
	v := NewVault()
	r := v.PlaceRoom(10, 0, 4)
	sp := v.Rooms[r].Sprite()
	if sp != "food1u1x3" {
		t.Fatal("Expected: food1u1x3, Found:", sp)
	}
}
