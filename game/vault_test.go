package game

import "testing"

func TestRoomSprite(t *testing.T) {
	v := NewVault()
	r := v.PlaceRoom(10, 0, 4)
	sp := v.Rooms[r].Sprite()
	if sp != "food1/level_1/size_3/sprite.png" {
		t.Fatal("Expected: food1u1x3, Found:", sp)
	}
}
