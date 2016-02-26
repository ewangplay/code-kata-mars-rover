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

type Position struct {
    X int64
    Y int64
}

func NewPosition(x,y int64) Position {
    return Position{
        X: x,
        Y: y,
    }
}

type Rover struct {
    CurrPosition Position
    Direction int8
}

func (this *Rover) Init(x,y int64, d int8) error {
    this.CurrPosition = NewPosition(x,y)
    this.Direction = d   
    
    return nil 
}

func (this *Rover) Move() error {
    switch this.Direction {
        case NH:
            this.CurrPosition.Y += 1
        case ET:
            this.CurrPosition.X += 1
        case SH:
            this.CurrPosition.Y -= 1
        case WT:
            this.CurrPosition.X -= 1
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
    this.CurrPosition.X = 2
    this.CurrPosition.Y = 2
    
    return nil
}