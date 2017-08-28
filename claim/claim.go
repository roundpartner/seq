package claim

import "github.com/roundpartner/seq/buffer"

type Item struct {
    Id int `json:"id"`
    Body string `json:"body"`
}

var id int = 0

type C struct {
    elastic *Elastic
    buf chan buffer.Message
}

func NewC(elastic *Elastic, sb *buffer.SimpleBuffer) *C {
    c := &C{elastic, sb.Messages}
    id = 0
    return c
}

func Next(elastic *Elastic, sb *buffer.SimpleBuffer) (Item, bool) {
    body, ok := sb.Pop()
    if false == ok {
        return Item{}, false
    }
    id++
    item := Item{Id: id, Body: body}
    elastic.In <- item
    return item, true
}
