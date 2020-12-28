package create_01

import . "patterns/create_prelude"

// The function CreateMaze() builds and returns a maze. One problem with this function is
// that it hard-codes the classes of maze, rooms, doors, and walls. We'll introduce factory
// methods to let subclasses choose these components.

// First we'll define factory methods in IMazeGame for creating the maze, room, wall, and
// door objects:

type IMazeGame interface {
	MakeMaze() IMaze
	MakeRoom(n int) IRoom
	MakeWall() IWall
	MakeDoor(r1 IRoom, r2 IRoom) IDoor
}

// Each factory method returns a maze component of a given type. mazeGame provides default
// implementations that return the simplest kinds of maze, rooms, walls, and doors.

type mazeGame struct{}

func NewMazeGame() *mazeGame { return &mazeGame{} }

func (_ mazeGame) MakeMaze() IMaze                   { return NewMaze() }
func (_ mazeGame) MakeRoom(n int) IRoom              { return NewRoom(n) }
func (_ mazeGame) MakeWall() IWall                   { return NewWall() }
func (_ mazeGame) MakeDoor(r1 IRoom, r2 IRoom) IDoor { return NewDoor(r1, r2) }

// Now we can rewrite CreateMaze to use these factory methods:
// MK - I can't see a way for this to be a method of IMazeGame so in Go this isn't much
//      different to Abstract Factory

func CreateMaze(g IMazeGame) IMaze {
	aMaze := g.MakeMaze()
	r1 := g.MakeRoom(1)
	r2 := g.MakeRoom(2)
	aDoor := g.MakeDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(North, g.MakeWall())
	r1.SetSide(East, aDoor)
	r1.SetSide(South, g.MakeWall())
	r1.SetSide(West, g.MakeWall())

	r2.SetSide(North, g.MakeWall())
	r2.SetSide(East, g.MakeWall())
	r2.SetSide(South, g.MakeWall())
	r2.SetSide(West, aDoor)

	return aMaze
}

// Different games can embed mazeGame and specialize parts of the maze. mazeGame derivatives
// can redefine some or all of the factory methods to specify variations in products. For
// example, a bombedMazeGame can redefine the IRoom and IWall products to return the bombed
// varieties:

type bombedMazeGame struct {
	mazeGame
}

func NewBombedMazeGame() *bombedMazeGame { return &bombedMazeGame{} }

func (_ bombedMazeGame) MakeRoom(n int) IRoom { return NewRoomWithABomb(n) }
func (_ bombedMazeGame) MakeWall() IWall      { return NewBombedWall() }

// An enchantedMazeGame variant might be defined like this:

type enchantedMazeGame struct {
	mazeGame
}

func NewEnchantedMazeGame() *enchantedMazeGame { return &enchantedMazeGame{} }

func (_ enchantedMazeGame) MakeRoom(n int) IRoom              { return NewEnchantedRoom(n) }
func (_ enchantedMazeGame) MakeDoor(r1 IRoom, r2 IRoom) IDoor { return NewDoorNeedingSpell(r1, r2) }

//------------------------------------------------------------------------------

func Main() {
	game := NewBombedMazeGame()
	maze := CreateMaze(game)
	maze.Run()
}
