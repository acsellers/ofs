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
