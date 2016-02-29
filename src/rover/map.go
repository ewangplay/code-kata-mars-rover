package rover

type Map struct {
    length int64
    width int64
}

func (this *Map)Init(l, w int64) error {
    this.length = l
    this.width = w
    return nil
}

func (this *Map)GetEastEdge() int64 {
    return 100
}

func (this *Map)GetWestEdge() int64 {
    return -100
}

func (this *Map)GetNorthEdge() int64 {
    return 100
}

func (this *Map)GetSouthEdge() int64 {
    return -100
}
