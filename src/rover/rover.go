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

func NewPosition(x, y int64) Position {
	return Position{
		X: x,
		Y: y,
	}
}

type Rover struct {
	position  Position
	direction int8
}

func (this *Rover) Init(m *Map, p Position, d int8) error {
	this.position = p
	this.direction = d

	return nil
}

func (this *Rover) GetPosition() Position {
	return this.position
}

func (this *Rover) GetDirection() int8 {
	return this.direction
}

func (this *Rover) MoveForward() error {
	switch this.direction {
	case NH:
		this.position.Y += 1
	case ET:
		this.position.X += 1
	case SH:
		this.position.Y -= 1
	case WT:
		this.position.X -= 1
	default:
		return fmt.Errorf("not setting direction")
	}

	return nil
}

func (this *Rover) MoveBackward() error {
	switch this.direction {
	case NH:
		this.position.Y -= 1
	case ET:
		this.position.X -= 1
	case SH:
		this.position.Y += 1
	case WT:
		this.position.X += 1
	default:
		return fmt.Errorf("not setting direction")
	}

	return nil
}

func (this *Rover) TurnLeft() error {
	switch this.direction {
	case NH:
		this.direction = WT
	case ET:
		this.direction = NH
	case SH:
		this.direction = ET
	case WT:
		this.direction = SH
	default:
		return fmt.Errorf("not setting direction")
	}

	return nil
}

func (this *Rover) TurnRight() error {
	switch this.direction {
	case NH:
		this.direction = ET
	case ET:
		this.direction = SH
	case SH:
		this.direction = WT
	case WT:
		this.direction = NH
	default:
		return fmt.Errorf("not setting direction")
	}

	return nil
}

func (this *Rover) Order(cmd string) error {
	var err error

	for i := 0; i < len(cmd); i++ {
		switch cmd[i] {
		case 'f':
			err = this.MoveForward()
		case 'b':
			err = this.MoveBackward()
		case 'r':
			err = this.TurnRight()
		case 'l':
			err = this.TurnLeft()
		default:
			err = fmt.Errorf("incorrect order")

		}

		if err != nil {
			return err
		}
	}

	return nil
}
