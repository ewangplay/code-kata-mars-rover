package rover

import (
	"fmt"
	"os"
	"testing"
)

var turn_directions = []int8{ET, SH, WT, NH}
var move_forward_positions = map[int8]Position{
	ET: NewPosition(1, 0),
	SH: NewPosition(0, -1),
	WT: NewPosition(-1, 0),
	NH: NewPosition(0, 1),
}
var move_backward_positions = map[int8]Position{
	ET: NewPosition(-1, 0),
	SH: NewPosition(0, 1),
	WT: NewPosition(1, 0),
	NH: NewPosition(0, -1),
}

var mars_map *Map

func TestMain(m *testing.M) {
	var err error

	mars_map = &Map{}
	err = mars_map.LoadData("mars_map_file.data")
	if err != nil {
		fmt.Printf("load mars map file fail: %v", err)
		return
	}

	os.Exit(m.Run())
}

func TestRoverMoveForward(t *testing.T) {
	var err error

	rover := &Rover{}

	for d, p := range move_forward_positions {

		err = rover.Init(mars_map, NewPosition(0, 0), d)
		if err != nil {
			t.Errorf(" rover fail: %v", err)
		}

		err = rover.MoveForward()
		if err != nil {
			t.Errorf("order rover fail: %v", err)
		}

		if rover.CurrPosition != p {
			t.Errorf("rover's postion should be %v", p)
		}
	}

}

func TestRoverMoveBackward(t *testing.T) {
	var err error

	rover := &Rover{}

	for d, p := range move_backward_positions {

		err = rover.Init(mars_map, NewPosition(0, 0), d)
		if err != nil {
			t.Errorf("init rover fail: %v", err)
		}

		err = rover.MoveBackward()
		if err != nil {
			t.Errorf("order rover fail: %v", err)
		}

		if rover.CurrPosition != p {
			t.Error("rover's postion should be %v", p)
		}
	}

}

func GetLeftPosition(pos int) int {
	if pos == 0 {
		return 3
	}
	return pos - 1
}

func GetRightPosition(pos int) int {
	if pos == 3 {
		return 0
	}
	return pos + 1
}

func TestRoverTurnLeft(t *testing.T) {
	var err error

	rover := &Rover{}

	for i, d := range turn_directions {
		err = rover.Init(mars_map, NewPosition(0, 0), d)
		if err != nil {
			t.Errorf("init rover fail: %v", err)
		}

		err = rover.TurnLeft()
		if err != nil {
			t.Errorf("rover turn left fail: %v", err)
		}

		if rover.Direction != turn_directions[GetLeftPosition(i)] {
			t.Errorf("rover facing to %v turn left will face to %v.", d, turn_directions[GetLeftPosition(i)])
		}
	}
}

func TestRoverTurnRight(t *testing.T) {
	var err error

	rover := &Rover{}

	for i, d := range turn_directions {
		err = rover.Init(mars_map, NewPosition(0, 0), d)
		if err != nil {
			t.Errorf("init rover fail: %v", err)
		}

		err = rover.TurnRight()
		if err != nil {
			t.Errorf("rover turn right fail: %v", err)
		}

		if rover.Direction != turn_directions[GetRightPosition(i)] {
			t.Errorf("rover facing to %v turn right will face to %v.", d, turn_directions[GetRightPosition(i)])
		}
	}
}

func TestRoverOrder(t *testing.T) {
	var err error

	rover := &Rover{}

	err = rover.Init(mars_map, NewPosition(0, 0), NH)
	if err != nil {
		t.Errorf("init rover fail: %v", err)
	}

	err = rover.Order("ffrff")
	if err != nil {
		t.Errorf("order rover fail: %v", err)
	}

	if rover.CurrPosition != NewPosition(2, 2) {
		t.Error("rover's postion is not expect(2,2)")
	}
}
