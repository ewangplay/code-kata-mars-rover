package rover

import (
    "testing"
)

func TestRoverMoveNorth(t *testing.T) {
    var err error
    
    rover := &Rover{}
            
    err = rover.Init(0,0,NH)
    if err != nil {
        t.Errorf("init rover fail: %v", err)        
    }

    err = rover.Move()
    if err != nil {
        t.Errorf("order rover fail: %v", err)        
    }
    
    if rover.PostionX != 0 && rover.PostionY != 1 {
        t.Error("rover's postion is not expect(0,1)")        
    }
}

func TestRoverMoveEAST(t *testing.T) {
    var err error
    
    rover := &Rover{}
    
    err = rover.Init(0,0,ET)
    if err != nil {
        t.Errorf("init rover fail: %v", err)        
    }

    err = rover.Move()
    if err != nil {
        t.Errorf("order rover fail: %v", err)        
    }
    
    if rover.PostionX != 1 && rover.PostionY != 0 {
        t.Error("rover's postion is not expect(1,0)")        
    }
}

func TestRoverMoveSouth(t *testing.T) {
    var err error
    
    rover := &Rover{}
    
    err = rover.Init(0,0,SH)
    if err != nil {
        t.Errorf("init rover fail: %v", err)        
    }

    err = rover.Move()
    if err != nil {
        t.Errorf("order rover fail: %v", err)        
    }
    
    if rover.PostionX != 0 && rover.PostionY != -1 {
        t.Error("rover's postion is not expect(0,-1)")        
    }
}

func TestRoverMoveWest(t *testing.T) {
    var err error
    
    rover := &Rover{}
    
    err = rover.Init(0,0,WT)
    if err != nil {
        t.Errorf("init rover fail: %v", err)        
    }

    err = rover.Move()
    if err != nil {
        t.Errorf("order rover fail: %v", err)        
    }
    
    if rover.PostionX != -1 && rover.PostionY != 0 {
        t.Error("rover's postion is not expect(-1,0)")        
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
    
    if rover.PostionX != 2 && rover.PostionY !=2 {
        t.Error("rover's postion is not expect(2,2)")
    }
}
