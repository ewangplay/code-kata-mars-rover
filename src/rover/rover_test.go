package rover

import (
    "testing"
)

func TestRoverOrder(t *testing.T) {
    
    rover := &Rover{}
    
    rover.Order("ffrff")
    
    if rover.PostionX != 2 && rover.PostionY !=2 {
        t.Error("rover's postion is not expect(2,2)")
    }
}
