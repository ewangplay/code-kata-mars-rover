package rover

import (
    "testing"
)

var turn_directions = []int8{ET,SH,WT,NH}
var move_forward_positions = map[int8]Position {
    ET: NewPosition(1, 0),
    SH: NewPosition(0, -1),
    WT: NewPosition(-1, 0),
    NH: NewPosition(0, 1),
}

func TestRoverMoveForward(t *testing.T) {
    var err error
    
    rover := &Rover{}
    
    for d, p := range move_forward_positions {
        
        err = rover.Init(0,0,d)
        if err != nil {
            t.Errorf("init rover fail: %v", err)        
        }

        err = rover.MoveForward()
        if err != nil {
            t.Errorf("order rover fail: %v", err)        
        }
        
        if rover.CurrPosition != p {
            t.Error("rover's postion is not expect(-1,0)")        
        }    
    }
    
}

func GetLeftPosition(pos int) int {
    if pos == 0 {
        return 3
    }
    return pos-1
}


func GetRightPosition(pos int) int {
    if pos == 3 {
        return 0
    }
    return pos+1
}

func TestRoverTurnLeft(t *testing.T) {
    var err error
    
    rover := &Rover{}
        
    for i, d := range turn_directions {
        err = rover.Init(0,0,d)
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
        err = rover.Init(0,0,d)
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
    
    err = rover.Init(0,0,NH)
    if err != nil {
        t.Errorf("init rover fail: %v", err)        
    }
    
    err = rover.Order("ffrff")
    if err != nil {
        t.Errorf("order rover fail: %v", err)
    }
    
    if rover.CurrPosition != NewPosition(2,2) {
        t.Error("rover's postion is not expect(2,2)")
    }
}
