package rover

import (
    "testing"
)

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
