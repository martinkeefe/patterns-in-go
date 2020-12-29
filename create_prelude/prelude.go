package create_prelude

import "fmt"

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

//------------------------------------------------------------------------------

type IMapSite interface {
	Enter()
}

type IRoom interface {
	IMapSite
	GetSide(d Direction) IMapSite
	SetSide(d Direction, s IMapSite)
	RoomNumber() int
	// For Prototype Pattern
	Clone() IRoom
	Initialize(roomNo int)
}

type IWall interface {
	IMapSite
	// For Prototype Pattern
	Clone() IWall
}

type IDoor interface {
	IMapSite
	OtherSideFrom(r IRoom) IRoom
	// For Prototype Pattern
	Clone() IDoor
	Initialize(r1 IRoom, r2 IRoom)
}

type IMaze interface {
	AddRoom(r IRoom)
	RoomNo(no int) (IRoom, bool)
	Run()
	// For Prototype Pattern
	Clone() IMaze
}

//------------------------------------------------------------------------------

type room struct {
	_sides      [4]IMapSite
	_roomNumber int
}

func NewRoom(roomNo int) *room { return &room{[4]IMapSite{}, roomNo} }

func (r *room) Enter()                          {}
func (r *room) GetSide(d Direction) IMapSite    { return r._sides[d] }
func (r *room) SetSide(d Direction, s IMapSite) { r._sides[d] = s }
func (r *room) RoomNumber() int                 { return r._roomNumber }
func (r *room) Clone() IRoom                    { return NewRoom(r._roomNumber) }
func (r *room) Initialize(roomNo int)           { r._roomNumber = roomNo }

//------------------------------------------------------------------------------

type wall struct{}

func NewWall() *wall { return &wall{} }

func (_ wall) Enter()       {}
func (w wall) Clone() IWall { return NewWall() }

//------------------------------------------------------------------------------

type door struct {
	_room1  IRoom
	_room2  IRoom
	_isOpen bool
}

func NewDoor(r1 IRoom, r2 IRoom) *door { return &door{r1, r2, false} }

func (d *door) Enter() {}
func (d *door) OtherSideFrom(r IRoom) IRoom {
	if r == d._room1 {
		return d._room2
	}
	return d._room1
}
func (d *door) Clone() IDoor                  { return NewDoor(d._room1, d._room2) }
func (d *door) Initialize(r1 IRoom, r2 IRoom) { d._room1 = r1; d._room2 = r2 }

//------------------------------------------------------------------------------

type maze map[int]IRoom

func NewMaze() maze { return make(maze) }

func (m maze) AddRoom(r IRoom)             { m[r.RoomNumber()] = r }
func (m maze) RoomNo(no int) (IRoom, bool) { room, ok := m[no]; return room, ok }
func (m maze) Run() {
	room, ok := m.RoomNo(1)
	if ok {
		room.Enter()
	}
}
func (m maze) Clone() IMaze { return NewMaze() }

//------------------------------------------------------------------------------
// Bombed Maze

type roomWithABomb struct{ room }

func NewRoomWithABomb(roomNo int) *roomWithABomb {
	fmt.Println("newRoomWithABomb")
	return &roomWithABomb{room{[4]IMapSite{}, roomNo}}
}

func (r *roomWithABomb) Enter() {
	fmt.Println("Bang!")
	for i := range r._sides {
		wall, ok := r._sides[i].(*bombedWall)
		if ok {
			wall.Enter()
		}
	}
}
func (r *roomWithABomb) Clone() IRoom { return NewRoomWithABomb(r._roomNumber) }

type bombedWall struct{ wall }

func NewBombedWall() *bombedWall   { fmt.Println("newBombedWall"); return &bombedWall{} }
func (_ *bombedWall) Enter()       { fmt.Println("Crash!") }
func (w *bombedWall) Clone() IWall { return NewBombedWall() }

//------------------------------------------------------------------------------
// Enchanted Maze

type enchantedRoom struct{ room }

func NewEnchantedRoom(roomNo int) *enchantedRoom {
	fmt.Println("newEnchantedRoom")
	return &enchantedRoom{room{[4]IMapSite{}, roomNo}}
}

type doorNeedingSpell struct{ door }

func NewDoorNeedingSpell(r1 IRoom, r2 IRoom) *doorNeedingSpell {
	fmt.Println("newDoorNeedingSpell")
	return &doorNeedingSpell{door{r1, r2, false}}
}
