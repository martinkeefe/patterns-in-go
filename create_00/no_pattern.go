package create_00

// Each room has four sides.

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

// The common interface for all the components of a maze.

type IMapSite interface {
	Enter()
}

// Concrete implementation of MapSite that defines the key
// relationships between components in the maze.

type room struct {
	_sides      [4]IMapSite
	_roomNumber int
}

func newRoom(roomNo int) *room {
	var sides [4]IMapSite
	return &room{sides, roomNo}
}

func (r *room) GetSide(d Direction) IMapSite    { return r._sides[d] }
func (r *room) SetSide(d Direction, s IMapSite) { r._sides[d] = s }
func (r *room) Enter()                          {}

// The following represent the wall or door that occurs on each side of a room.

type wall struct{}

func (_ wall) Enter() {}

type door struct {
	_room1  *room
	_room2  *room
	_isOpen bool
}

func newDoor(r1 *room, r2 *room) *door { return &door{r1, r2, false} }

func (_ door) Enter() {}
func (d door) OtherSideFrom(r *room) *room {
	if r == d._room1 {
		return d._room2
	}
	return d._room1
}

type maze map[int]*room

func (m maze) AddRoom(room *room)  { m[room._roomNumber] = room }
func (m maze) RoomNo(no int) *room { return m[no] }

// One straightforward way to create a maze is with a series of operations that
// add components to a maze and then interconnect them. For example, the following
// function will create a maze consisting of two rooms with a door between them:

type mazeGame struct{}

func newMazeGame() mazeGame { return mazeGame{} }

func (_ mazeGame) CreateMaze() maze {
	aMaze := make(maze)
	r1 := newRoom(1)
	r2 := newRoom(2)
	aDoor := newDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(North, wall{})
	r1.SetSide(East, aDoor)
	r1.SetSide(South, wall{})
	r1.SetSide(West, wall{})

	r2.SetSide(North, wall{})
	r2.SetSide(East, wall{})
	r2.SetSide(South, wall{})
	r2.SetSide(West, aDoor)

	return aMaze
}

func Main() {
	game := newMazeGame()
	game.CreateMaze()
}
