package claim

type Elastic struct {
    In chan Item
    Out chan Item
    buffer []Item
}

func New() *Elastic {
    e := &Elastic{
        In: make(chan Item),
        Out: make(chan Item),
        buffer: make([]Item, 0),
    }
    go e.run()
    return e
}

func (b *Elastic) run() {
    for {
        if len(b.buffer) > 0 {
            select {
            case b.Out<-b.buffer[0]:
                b.buffer = b.buffer[1:]
            case value := <-b.In:
                b.buffer = append(b.buffer, value)
            }
        } else {
            value := <-b.In
            b.buffer = append(b.buffer, value)
        }
    }
}