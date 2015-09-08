package game

import (
	"math/rand"
	"time"
)

type Walkabout struct {
	CurrentVault *Vault
	Dweller      PersonID
	Items        [99]ItemID
	Money        int
	Distance     time.Duration
	LastRun      time.Time
	Log          []WalkaboutEvent
}

func (w Walkabout) Step() {
	if w.CurrentVault.Dwellers[w.Dweller].IsDead {
		return
	}
	change := w.CurrentVault.Time.Sub(w.LastRun)
	w.LastRun = w.CurrentVault.Time

	// run through the dweller's encounters
	for change > 30*time.Second {
		w.Distance += time.Minute
		change -= time.Minute

		if we, ok := w.rollEncounter(); ok {
			w.ResolveEvent(we)
			if w.CurrentVault.Dwellers[w.Dweller].IsDead {
				break
			}
		}
	}

}

func (w Walkabout) rollEncounter() (WalkaboutEvent, bool) {
	if rand.Intn(5) == 0 {
		day := 24 * time.Hour
		switch {
		case w.Distance < day/2:
		case w.Distance < day:
		case w.Distance < day*2:
		case w.Distance < day*3:
		case w.Distance < day*4:
		}
	}
	return nil, false
}

func (w Walkabout) ResolveEvent(we WalkaboutEvent) {
}

func (w Walkabout) Return() {
}
func (w Walkabout) Collect() {

}

type WalkaboutEvent interface {
}
