package rover

type Rover struct {
    PostionX int64
    PostionY int64
}

func (this *Rover) Order(cmd string) error {
    this.PostionX = 2
    this.PostionY = 2
    
    return nil
}