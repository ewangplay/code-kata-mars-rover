package rover

import (
    "fmt"
)

type CoordData struct {
	Is_obstacle   bool
	Obstacle_desc string
}

type Map struct {
	length int64
	width  int64
	coords map[Position]CoordData
}

func (this *Map) Init(l, w int64) error {
	this.length = l
	this.width = w
	this.coords = make(map[Position]CoordData, l*w)

	//Init map coords
	for i := -(l / 2); i <= l/2; i++ {
		for j := -(w / 2); j <= w/2; j++ {
			this.coords[NewPosition(i, j)] = CoordData{false, ""}
		}
	}

	return nil
}

func (this *Map) GetEastEdge() int64 {
	return this.length / 2
}

func (this *Map) GetWestEdge() int64 {
	return -(this.length / 2)
}

func (this *Map) GetNorthEdge() int64 {
	return this.width / 2
}

func (this *Map) GetSouthEdge() int64 {
	return -(this.width / 2)
}

func (this *Map) Print() {
	for i := -(this.length / 2); i <= this.length/2; i++ {
		for j := -(this.width / 2); j <= this.width/2; j++ {
			data, ok := this.coords[NewPosition(i, j)]
			if ok {
				if data.Is_obstacle {
					fmt.Print("D")
				} else {
					fmt.Print("O")
				}
			} else {
				//unknown area
				fmt.Print("U")
			}
		}
        fmt.Println()
	}
}
