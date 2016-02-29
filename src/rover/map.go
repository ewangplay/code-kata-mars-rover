package rover

type Map struct {

}

func (this *Map)LoadData(filename string) error {
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
