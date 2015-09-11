package game

import (
	"fmt"
	"log"
	"strings"
	"time"
	"unicode"

	"github.com/pietv/astar"
)

type Vault struct {
	Rooms map[RoomID]RoomInstance
	// a RoomID usually spans more than 1 block, excepting elevators
	// not actually a map, but a table
	RoomMap [50][30]RoomID
	// The next RoomID for the vault
	RoomCnt int
	// [row][col], first 10 columns on row 0 are taken up by door and wasteland
	// normal room costs 3 columns, elevator takes 1 column
	Dwellers map[PersonID]Person
	// The next person ID for the vault
	DwellerCnt int
	// Vault time passes differently from normal time
	// start at january 1, 2177
	Time time.Time
	// places around the vault to discover
	Landmarks LandmarkList
	// factions to befriend and piss off
	Factions map[FactionID]Faction
}

type RoomID int

type RoomInstance struct {
	ID          RoomID
	Type        Room
	Size        int
	Upgrade1    int
	Upgrade2    int
	Row, Col    int
	Connections []RoomConnection
}

func (ri RoomInstance) Sprite() string {
	r := Rooms[ri.Type]
	name := strings.ToLower(
		strings.TrimFunc(
			r.Code,
			func(r rune) bool {
				return !(unicode.IsLetter(r) || unicode.IsDigit(r))
			},
		),
	)

	return fmt.Sprintf(
		"rooms/%s/level_%d/size_%d/sprite.png",
		name,
		ri.Upgrade1,
		ri.Size,
	)
}

func (ri RoomInstance) Location(offsetx, offsety float32) (float32, float32) {
	x := float32(ri.Col*40) + offsetx
	y := float32(ri.Row*80) + offsety
	return x, y
}

type RoomConnection struct {
	To, From RoomID
	// 2 per block, 1 between floors, 1 to get on elevator
	Cost int
}

func NewVault() Vault {
	v := Vault{
		// guess when this is
		Time:     time.Date(2277, time.August, 17, 9, 04, 0, 0, time.UTC),
		Dwellers: make(map[PersonID]Person),

		// add vault door
		RoomCnt: 2,
		Rooms: map[RoomID]RoomInstance{
			1: RoomInstance{
				ID:       1,
				Type:     16,
				Size:     9,
				Upgrade1: 1,
				Upgrade2: 1,
			},
		},
	}
	// mark vault door on map
	for i := 0; i < 9; i++ {
		v.RoomMap[0][i] = 1
	}
	return v
}

// PlaceRoom adds a new room of type r at col and row
// It will panic if there is an active room already at that point
func (v *Vault) PlaceRoom(col, row int, r Room) RoomID {
	if v.RoomMap[row][col] != 0 {
		log.Println("Placing new room at", col, row, "over existing room")
		panic(v)
	}

	ri := RoomInstance{
		ID:       RoomID(v.RoomCnt),
		Type:     r,
		Size:     3,
		Upgrade1: 1,
		Upgrade2: 1,
		Col:      col,
		Row:      row,
	}
	v.RoomCnt++
	// Elevator sizing
	if r == 1 {
		ri.Size = 1
		ri.Upgrade1 = 0
		ri.Upgrade2 = 0

		// look for elevator connections
		for i := 0; i < 40; i++ {
			ori := v.RoomMap[i][col]
			if ori != 0 && v.Rooms[ori].Type == 1 {
				other := v.Rooms[ori]
				diff := row - i
				if diff < 0 {
					diff = -diff
				}

				// direct connect to the other elevator
				other.Connections = append(other.Connections, RoomConnection{
					To:   ri.ID,
					From: ori,
					Cost: 1 + diff,
				})
				v.Rooms[ori] = other

				// add connection to self
				ri.Connections = append(ri.Connections, RoomConnection{
					To:   ori,
					From: ri.ID,
					Cost: 1 + diff,
				})
			}
		}
	}

	// connect to left, if it exists
	lri := v.RoomMap[row][col-1]
	if lri != 0 {
		left := v.Rooms[lri]

		// add connection to us on the neighbor
		left.Connections = append(left.Connections, RoomConnection{
			From: lri,
			To:   ri.ID,
			Cost: left.Size + ri.Size,
		})
		v.Rooms[lri] = left

		// add connection to self
		ri.Connections = append(ri.Connections, RoomConnection{
			From: ri.ID,
			To:   lri,
			Cost: left.Size + ri.Size,
		})
	}

	// connect to left, if it exists
	rri := v.RoomMap[row][col+ri.Size]
	if rri != 0 {
		right := v.Rooms[rri]

		// add connection to us on the neighbor
		right.Connections = append(right.Connections, RoomConnection{
			From: rri,
			To:   ri.ID,
			Cost: right.Size + ri.Size,
		})
		v.Rooms[rri] = right

		// add connection to self
		ri.Connections = append(ri.Connections, RoomConnection{
			From: ri.ID,
			To:   rri,
			Cost: right.Size + ri.Size,
		})
	}

	// add us to the vault
	v.Rooms[ri.ID] = ri
	for i := 0; i < ri.Size; i++ {

		v.RoomMap[row][col+i] = ri.ID
	}

	return ri.ID
}

// Route will return a quick, valid path from the room with the
// RoomID from to the room with RoomID to.
func (v Vault) Route(from, to RoomID) []RoomID {
	s := &searchVault{to, from, &v, from}
	path, _, err := astar.Search(s)
	if err != nil {
		v.PrintVault()
		log.Fatal("Path Search Error", err)
	}
	rooms := []RoomID{}
	for _, it := range path {
		rooms = append(rooms, it.(RoomID))
	}
	return rooms
}

func (v Vault) PrintVault() {
	for _, row := range v.RoomMap {
		fmt.Println(row)
	}
}

type searchVault struct {
	to, from RoomID
	v        *Vault
	curr     RoomID
}

func (sv searchVault) Start() interface{} {
	return sv.from
}

func (sv searchVault) Finish() bool {
	return sv.curr == sv.to
}

func (sv *searchVault) Move(state interface{}) {
	sv.curr = state.(RoomID)
}

func (sv searchVault) Successors() []interface{} {
	successors := []interface{}{}
	for _, conn := range sv.v.Rooms[sv.curr].Connections {
		successors = append(successors, conn.To)
	}
	return successors
}

func (sv searchVault) Cost(given interface{}) float64 {
	gri := given.(RoomID)
	for _, conn := range sv.v.Rooms[sv.curr].Connections {
		if gri == conn.To {
			return float64(conn.Cost)
		}
	}
	return 999999
}

func (sv searchVault) Estimate(given interface{}) float64 {
	ri := given.(RoomID)
	rows := float64(sv.v.Rooms[sv.from].Row-sv.v.Rooms[ri].Row) * 2
	cols := float64(sv.v.Rooms[sv.from].Col - sv.v.Rooms[ri].Col)
	if cols != 0 {
		cols += 1
	}

	return rows + cols
}

type DwellerFilter int

const (
	// Sort and Restriction
	Name DwellerFilter = iota
	Level
	HP
	Mood
	DPS
	CurrentJob

	// special
	StatCharm
	StatAwareness
	StatPower
	StatSuccess
	StatTenacity
	StatOutgoing
	StatNimble
	StatErudite

	// Restriction only
	Dead
)

type Operation int

const (
	Eq Operation = iota
	Neq
	Gt
	Lt
)

type DwellerRestriction struct {
	Filter    DwellerFilter
	Op        Operation
	BoolValue bool
	StrValue  string
	IntValue  int
}

func (dr DwellerRestriction) Matches(p Person) bool {
	switch dr.Filter {
	case Name:
		return dr.MatchName(p)
	case HP:
		return dr.MatchHP(p)
	case Dead:
		switch dr.Op {
		case Eq:
			return p.IsDead == dr.BoolValue
		case Neq:
			return p.IsDead != dr.BoolValue
		}
	case CurrentJob:
		switch dr.Op {
		case Eq:
			return p.CurrentJob == Job(dr.IntValue)
		case Neq:
			return p.CurrentJob != Job(dr.IntValue)
		}
	case Level:
		return dr.MatchStat(p.Level)
	case Mood:
		return dr.MatchStat(p.Mood)
	case DPS:
		return dr.MatchStat(p.SimpleDPS())
	case StatCharm:
		return dr.MatchStat(p.Stats.Charm)
	case StatAwareness:
		return dr.MatchStat(p.Stats.Awareness)
	case StatPower:
		return dr.MatchStat(p.Stats.Power)
	case StatSuccess:
		return dr.MatchStat(p.Stats.Success)
	case StatTenacity:
		return dr.MatchStat(p.Stats.Tenacity)
	case StatOutgoing:
		return dr.MatchStat(p.Stats.Outgoing)
	case StatNimble:
		return dr.MatchStat(p.Stats.Nimble)
	case StatErudite:
		return dr.MatchStat(p.Stats.Erudite)

	}
	return false
}

func (dr DwellerRestriction) MatchStat(v int) bool {
	switch dr.Op {
	case Gt:
		return dr.IntValue > v
	case Lt:
		return dr.IntValue < v
	case Eq:
		return dr.IntValue == v
	case Neq:
		return dr.IntValue != v
	}
	return false
}

func (dr DwellerRestriction) MatchName(p Person) bool {
	switch dr.Op {
	case Gt:
		return strings.ToLower(p.Name) > strings.ToLower(dr.StrValue)
	case Lt:
		return strings.ToLower(p.Name) < strings.ToLower(dr.StrValue)
	case Eq:
		return strings.HasPrefix(strings.ToLower(p.Name), strings.ToLower(dr.StrValue))
	case Neq:
		return strings.ToLower(p.Name) != strings.ToLower(dr.StrValue)
	}
	return false
}

// MatchHP matches against the integer percenage of HP remaining (0-100)
func (dr DwellerRestriction) MatchHP(p Person) bool {
	switch dr.Op {
	case Gt:
		return p.CurrentHP*100/p.TotalHP > dr.IntValue
	case Eq:
		return p.CurrentHP*100/p.TotalHP == dr.IntValue
	case Neq:
		return p.CurrentHP*100/p.TotalHP != dr.IntValue
	case Lt:
		return p.CurrentHP*100/p.TotalHP < dr.IntValue
	}
	return false
}

type DwellerList struct {
	Dwellers  []Person
	SortKey   DwellerFilter
	Ascending bool
}

func (dl DwellerList) Len() int {
	return len(dl.Dwellers)
}

func (dl DwellerList) Less(i, j int) bool {
	di := dl.Dwellers[i]
	dj := dl.Dwellers[j]

	if dl.Ascending {
		dj = dl.Dwellers[i]
		di = dl.Dwellers[j]
	}
	// act like i to j is highest to lowest
	switch dl.SortKey {
	case Name:
		return di.Name > dj.Name
	case Level:
		return di.Level > dj.Level
	case HP:
		return di.CurrentHP*100/di.TotalHP > dj.CurrentHP*100/dj.TotalHP
	case CurrentJob:
		return di.CurrentJob > dj.CurrentJob
	case StatCharm:
		return di.Stats.Charm > dj.Stats.Charm
	case StatAwareness:
		return di.Stats.Awareness > dj.Stats.Awareness
	case StatPower:
		return di.Stats.Power > dj.Stats.Power
	case StatSuccess:
		return di.Stats.Success > dj.Stats.Success
	case StatTenacity:
		return di.Stats.Tenacity > dj.Stats.Tenacity
	case StatOutgoing:
		return di.Stats.Outgoing > dj.Stats.Outgoing
	case StatNimble:
		return di.Stats.Nimble > dj.Stats.Nimble
	case StatErudite:
		return di.Stats.Erudite > dj.Stats.Erudite
	case Mood:
		return di.Mood > dj.Mood
	case DPS:
		return di.SimpleDPS() > dj.SimpleDPS()
	}
	return false
}

func (dl DwellerList) Swap(i, j int) {
	dl.Dwellers[i], dl.Dwellers[j] = dl.Dwellers[j], dl.Dwellers[i]
}
func (v Vault) FilterDwellers(filter DwellerFilter, ascending bool, restrictions []DwellerRestriction) []Person {
	dwellers := []Person{}
DwellCheck:
	for _, p := range v.Dwellers {
		for _, r := range restrictions {
			if !r.Matches(p) {
				continue DwellCheck
			}
		}
		dwellers = append(dwellers, p)
	}
	return dwellers
}
