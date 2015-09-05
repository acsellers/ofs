package game

import "testing"

func TestRoomSprite(t *testing.T) {
	v := NewVault()
	r := v.PlaceRoom(9, 0, 4)
	sp := v.Rooms[r].Sprite()
	if sp != "rooms/food1/level_1/size_3/sprite.png" {
		t.Fatal("Expected: 'rooms/food1/level_1/size_3/sprite.png', Found:", sp)
	}
}
