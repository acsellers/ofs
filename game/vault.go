package game

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/pietv/astar"
)

type Vault struct {
	Rooms map[RoomID]RoomInstance
	next  RoomID
	// a RoomID usually spans more than 1 block, excepting elevators
	// not actually a map, but a table
	RoomMap [40][50]RoomID
	// [row][col], first 10 columns on row 0 are taken up by door and wasteland
	// normal room costs 3 columns, elevator takes 1 column
}

type RoomID int
type RoomInstance struct {
	ID           RoomID
	Type         Room
	Size         int
	UpgradeLevel int
	Row, Col     int
	Connections  []RoomConnection
}

func (ri RoomInstance) Sprite() string {
	r := Rooms[ri.Type]
	var name string
	switch ri.UpgradeLevel {
	case 1:
		name = r.Lv1Name
	case 2:
		name = r.Lv2Name
	case 3:
		name = r.Lv3Name
	}
	name = fmt.Sprintf("%sx%d", name, ri.Size)

	return strings.ToLower(
		strings.TrimFunc(
			name,
			func(r rune) bool {
				return !(unicode.IsLetter(r) || unicode.IsDigit(r))
			},
		),
	)

}

type RoomConnection struct {
	To, From RoomID
	// 2 per block, 1 between floors, 1 to get on elevator
	Cost int
}

func NewVault() Vault {
	v := Vault{
		Rooms: map[RoomID]RoomInstance{
			1: RoomInstance{
				ID:   1,
				Type: -1,
				Size: 9,
			},
		},
		next: 2,
	}
	// mark vault door on map
	for i := 0; i < 10; i++ {
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
		ID:           v.next,
		Type:         r,
		Size:         3,
		UpgradeLevel: 1,
		Col:          col,
		Row:          row,
	}
	v.next += 1
	// Elevator sizing
	if r == 1 {
		ri.Size = 1
		ri.UpgradeLevel = 0

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
		for _, row := range v.RoomMap {
			fmt.Println(row)
		}
		log.Fatal("Path Search Error", err)
	}
	rooms := []RoomID{}
	for _, it := range path {
		rooms = append(rooms, it.(RoomID))
	}
	return rooms
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
