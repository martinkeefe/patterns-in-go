package create_04

import (
	"patterns/create_02"
	. "patterns/create_prelude"
)

type IMazePrototypeFactory interface{ create_02.IMazeFactory }

type mazePrototypeFactory struct {
	_prototypeMaze IMaze
	_prototypeRoom IRoom
	_prototypeWall IWall
	_prototypeDoor IDoor
}

func newMazePrototypeFactory(m IMaze, w IWall, r IRoom, d IDoor) mazePrototypeFactory {
	return mazePrototypeFactory{m, r, w, d}
}

func (f mazePrototypeFactory) MakeMaze() IMaze { return f._prototypeMaze.Clone() }
func (f mazePrototypeFactory) MakeRoom(n int) IRoom {
	r := f._prototypeRoom.Clone()
	r.Initialize(n)
	return r
}
func (f mazePrototypeFactory) MakeWall() IWall { return f._prototypeWall.Clone() }
func (f mazePrototypeFactory) MakeDoor(r1 IRoom, r2 IRoom) IDoor {
	d := NewDoor(r1, r2)
	d.Initialize(r1, r2)
	return d
}

//------------------------------------------------------------------------------

func Main() {
	game := create_02.NewMazeGame()
	// simpleMazeFactory := newMazePrototypeFactory(NewMaze(), NewWall(), NewRoom(0), NewDoor(nil, nil))
	bombedMazeFactory := newMazePrototypeFactory(NewMaze(), NewBombedWall(), NewRoomWithABomb(0), NewDoor(nil, nil))
	// maze := game.CreateMaze(simpleMazeFactory)
	maze := game.CreateMaze(bombedMazeFactory)
	maze.Run()
}
