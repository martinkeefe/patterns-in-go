package create_03

import (
	"fmt"
	. "patterns/create_prelude"
)

// IMazeBuilder defines the following interface for building mazes:

type IMazeBuilder interface {
	BuildMaze()
	BuildRoom(room int)
	BuildDoor(roomFrom, roomTo int)
	GetMaze() IMaze
}

// This interface can create three things: (1) the maze, (2) rooms with a
// particular room number, and (3) doors between numbered rooms. The GetMaze
// operation returns the maze to the client. Implementations of IMazeBuilder
// implement this operation to return the maze that they build.

type mazeGame struct{}

func NewMazeGame() mazeGame { return mazeGame{} }

// Given the IMazeBuilder interface, we can change the CreateMaze member
// function to take this builder as a parameter.

func (_ mazeGame) CreateMaze(builder IMazeBuilder) IMaze {
	builder.BuildMaze()

	builder.BuildRoom(1)
	builder.BuildRoom(2)
	builder.BuildDoor(1, 2)

	return builder.GetMaze()
}

// The type standardMazeBuilder is an implementation that builds simple mazes.
// It keeps track of the maze it's building in the variable _currentMaze.

type standardMazeBuilder struct {
	_currentMaze IMaze
}

// The standardMazeBuilder constructor simply initializes _currentMaze.

func newStandardMazeBuilder() *standardMazeBuilder {
	return &standardMazeBuilder{nil}
}

// CommonWall is a utility operation that determines the direction of the common
// wall between two rooms.

func (_ *standardMazeBuilder) CommonWall(r1 IRoom, r2 IRoom) Direction {
	// MK - Not enough info available for a proper implementation so...
	if r1.RoomNumber() == 1 {
		return East
	} else {
		return West
	}
}

// BuildMaze instantiates an IMaze that other operations will assemble and
// eventually return to the client (with GetMaze).

func (b *standardMazeBuilder) BuildMaze()     { b._currentMaze = NewMaze() }
func (b *standardMazeBuilder) GetMaze() IMaze { return b._currentMaze }

// The BuildRoom operation creates a room and builds the walls around it:

func (b *standardMazeBuilder) BuildRoom(n int) {
	_, ok := b._currentMaze.RoomNo(n)
	if !ok {
		room := NewRoom(n)
		b._currentMaze.AddRoom(room)

		room.SetSide(North, NewWall())
		room.SetSide(South, NewWall())
		room.SetSide(East, NewWall())
		room.SetSide(West, NewWall())
	}
}

// To build a door between two rooms, standardMazeBuilder looks up both rooms in
// the maze and finds their adjoining wall:

func (b *standardMazeBuilder) BuildDoor(n1, n2 int) {
	r1, ok1 := b._currentMaze.RoomNo(n1)
	r2, ok2 := b._currentMaze.RoomNo(n2)

	if ok1 && ok2 {
		d := NewDoor(r1, r2)
		r1.SetSide(b.CommonWall(r1, r2), d)
		r2.SetSide(b.CommonWall(r2, r1), d)
	}
}

// We could have put all the standardMazeBuilder operations in maze and let each
// maze build itself. But making maze smaller makes it easier to understand and
// modify, and standardMazeBuilder is easy to separate from maze. Most
// importantly, separating the two lets you have a variety of MazeBuilders, each
// using different types for rooms, walls, and doors.

//------------------------------------------------------------------------------

// A more exotic IMazeBuilder is countingMazeBuilder. This builder doesn't
// create a maze at all; it just counts the different kinds of components that
// would have been created.

type countingMazeBuilder struct {
	_doors int
	_rooms int
}

// The constructor initializes the counters, and the overridden IMazeBuilder
// operations increment them accordingly.

func newCountingMazeBuilder() *countingMazeBuilder { return &countingMazeBuilder{0, 0} }

func (b *countingMazeBuilder) BuildMaze()                     {}
func (b *countingMazeBuilder) BuildRoom(room int)             { b._rooms++ }
func (b *countingMazeBuilder) BuildDoor(roomFrom, roomTo int) { b._doors++ }
func (b *countingMazeBuilder) GetMaze() IMaze                 { return nil }
func (b *countingMazeBuilder) GetCounts() (int, int)          { return b._rooms, b._doors }

//------------------------------------------------------------------------------

func Main() {
	game := NewMazeGame()
	if false {
		builder := newStandardMazeBuilder()
		maze := game.CreateMaze(builder)
		maze.Run()
	} else {
		builder := newCountingMazeBuilder()
		game.CreateMaze(builder)
		room, doors := builder.GetCounts()
		fmt.Printf("The maze has %d rooms and %d doors\n", room, doors)
	}
}
