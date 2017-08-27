package claim

type Elastic struct {
    In chan Item
    Out chan Item
    Query chan Query
    buffer []Item
}

type Query struct {
    Id int
    Out chan Item
    Delete bool
}

func New() *Elastic {
    e := &Elastic{
        In: make(chan Item),
        Out: make(chan Item),
        Query: make(chan Query),
        buffer: make([]Item, 0),
    }
    go e.run()
    return e
}

func (b *Elastic) run() {
    for {
        if len(b.buffer) > 0 {
            select {
                case qry := <-b.Query:
                    query(b, qry)
                case b.Out<-b.buffer[0]:
                    b.buffer = b.buffer[1:]
                case value := <-b.In:
                    b.buffer = append(b.buffer, value)
            }
        } else {
            select {
                case qry := <-b.Query:
                    qry.Out <- Item{}
                case value := <-b.In:
                    b.buffer = append(b.buffer, value)
            }
        }
    }
}

func query(b *Elastic, qry Query) {
    for i := range b.buffer {
        if b.buffer[i].Id == qry.Id {
            qry.Out <- b.buffer[i]
            if qry.Delete {
                b.buffer = append(b.buffer[:i], b.buffer[i+1:]...)
            }
            return
        }
    }
    qry.Out <- Item{}
}
