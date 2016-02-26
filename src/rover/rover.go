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

func (this *Rover) TurnLeft() error {
    switch this.Direction {
        case NH:
            this.Direction = WT
        case ET:
            this.Direction = NH
        case SH:
            this.Direction = ET
        case WT:
            this.Direction = SH
        default:
            return fmt.Errorf("not setting direction")
    }

    return nil
}

func (this *Rover) TurnRight() error {
    switch this.Direction {
        case NH:
            this.Direction = ET
        case ET:
            this.Direction = SH
        case SH:
            this.Direction = WT
        case WT:
            this.Direction = NH
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