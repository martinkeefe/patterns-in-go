package create_02

import . "patterns/create_prelude"

// An IMazeFactory can create components of mazes. It builds rooms, walls, and
// doors between rooms. It might be used by a program that reads plans for mazes
// from a file and builds the corresponding maze. Or it might be used by a program
// that builds mazes randomly. Programs that build mazes take a MazeFactory as an
// argument so that the programmer can specify the classes of rooms, walls, and
// doors to construct.

type IMazeFactory interface {
	MakeMaze() IMaze
	MakeRoom(n int) IRoom
	MakeWall() IWall
	MakeDoor(r1 IRoom, r2 IRoom) IDoor
}

type mazeFactory struct{}

func NewMazeFactory() mazeFactory { return mazeFactory{} }

func (_ mazeFactory) MakeMaze() IMaze                   { return NewMaze() }
func (_ mazeFactory) MakeRoom(n int) IRoom              { return NewRoom(n) }
func (_ mazeFactory) MakeWall() IWall                   { return NewWall() }
func (_ mazeFactory) MakeDoor(r1 IRoom, r2 IRoom) IDoor { return NewDoor(r1, r2) }

//------------------------------------------------------------------------------

type mazeGame struct{}

func NewMazeGame() mazeGame { return mazeGame{} }

// Recall that the member function CreateMaze () builds a small maze consisting
// of two rooms with a door between them. CreateMaze hard-codes the class names,
// making it difficult to create mazes with different components.

// Here's a version of CreateMaze that remedies that shortcoming by taking an
// IMazeFactory as a parameter:

func (_ mazeGame) CreateMaze(factory IMazeFactory) IMaze {
	aMaze := factory.MakeMaze()
	r1 := factory.MakeRoom(1)
	r2 := factory.MakeRoom(2)
	aDoor := factory.MakeDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(North, factory.MakeWall())
	r1.SetSide(East, aDoor)
	r1.SetSide(South, factory.MakeWall())
	r1.SetSide(West, factory.MakeWall())

	r2.SetSide(North, factory.MakeWall())
	r2.SetSide(East, factory.MakeWall())
	r2.SetSide(South, factory.MakeWall())
	r2.SetSide(West, aDoor)

	return aMaze
}

// We can create enchantedMazeFactory, a factory for enchanted mazes, by embedding
// mazeFactory. enchantedMazeFactory will override different member functions and
// return different subclasses of Room, Wall, etc.

type enchantedMazeFactory struct{ mazeFactory }

func NewEnchantedMazeFactory() enchantedMazeFactory { return enchantedMazeFactory{mazeFactory{}} }

func (_ enchantedMazeFactory) MakeRoom(n int) IRoom              { return NewEnchantedRoom(n) }
func (_ enchantedMazeFactory) MakeDoor(r1 IRoom, r2 IRoom) IDoor { return NewDoorNeedingSpell(r1, r2) }

// Now suppose we want to make a maze game in which a room can have a bomb set in
// it. If the bomb goes off, it will damage the walls (at least). We can make a
// derivative of Room keep track of whether the room has a bomb in it and whether
// the bomb has gone off. We'll also need a derivative of Wall to keep track of
// the damage done to the wall. We'll call these roomWithABomb and bombedWall.

// The last type we'll define is bombedMazeFactory, a derivative of mazeFactory
// that ensures walls are of type bombedWall and rooms are of type roomWithABomb.
// bombedMazeFactory only needs to override two functions:

type bombedMazeFactory struct{ mazeFactory }

func NewBombedMazeFactory() bombedMazeFactory { return bombedMazeFactory{mazeFactory{}} }

func (_ bombedMazeFactory) MakeRoom(n int) IRoom { return NewRoomWithABomb(n) }
func (_ bombedMazeFactory) MakeWall() IWall      { return NewBombedWall() }

// CreateMaze used the SetSide operation on rooms to specify their sides. If it
// creates rooms with a bombedMazeFactory, then the maze will be made up of
// roomWithABomb objects with bombedWall sides. If roomWithABomb had to access a
// type-specific member of bombedWall, then it would have to cast a reference to
// its walls from IWall to bombedWall.

// func (r roomWithABomb) Enter() {
// 	fmt.Println("Bang!")
// 	for i := range r._sides {
// 		wall, ok := r._sides[i].(bombedWall)
// 		if ok {
// 			wall.Enter()
// 		}
// 	}
// }

// func (m maze) Run() {
// 	room, ok := m.RoomNo(1)
// 	if ok {
// 		room.Enter()
// 	}
// }

//------------------------------------------------------------------------------

func Main() {
	game := NewMazeGame()
	// factory := NewMazeFactory()
	factory := NewEnchantedMazeFactory()
	// factory := NewBombedMazeFactory()
	maze := game.CreateMaze(factory)
	maze.Run()
}
