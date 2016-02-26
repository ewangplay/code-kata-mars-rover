package rover

import (
    "fmt"
)

const (
    NH int8 = 0
    ET int8 = 1
    SH int8 = 2
    WT int8 = 3
)
type Rover struct {
    PostionX int64
    PostionY int64
    Direction int8
}

func (this *Rover) Init(x,y int64, d int8) error {
    this.PostionX = x
    this.PostionY = y
    this.Direction = d   
    
    return nil 
}

func (this *Rover) Move() error {
    switch this.Direction {
        case NH:
            this.PostionY += 1
        case ET:
            this.PostionX += 1
        case SH:
            this.PostionY -= 1
        case WT:
            this.PostionX -= 1
        default:
            return fmt.Errorf("not setting direction")
    }
    
    return nil
}
func (this *Rover) Order(cmd string) error {
    this.PostionX = 2
    this.PostionY = 2
    
    return nil
}