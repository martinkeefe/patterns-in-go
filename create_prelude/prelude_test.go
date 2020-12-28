package create_prelude

import "testing"

func TestGetSide(t *testing.T) {
	w1 := NewWall()
	r := NewRoom(1)
	r.SetSide(North, w1)

	w2 := r.GetSide(North).(*wall)
	if w2 != w1 {
		t.Errorf("TestGetSide : %p != %p", w1, w2)
	}
}

func TestOtherSide(t *testing.T) {
	r1 := NewRoom(1)
	r2 := NewRoom(2)
	d := NewDoor(r1, r2)

	r := d.OtherSideFrom(r1) //.(*room)
	p1 := r2
	p2 := r.(*room)
	if p1 != p2 {
		t.Errorf("TestOtherSide : %p != %p", p1, p2)
	}
}

func TestRoomNo(t *testing.T) {
	m := NewMaze()
	r1 := NewRoom(1)

	m.AddRoom(r1)
	ir, ok := m.RoomNo(1)
	if ok {
		p1 := r1
		p2 := ir.(*room)
		if p1 != p2 {
			t.Errorf("TestRoomNo : %p != %p", p1, p2)
		}
	} else {
		t.Errorf("TestRoomNo : ok == %t", ok)
	}
}
