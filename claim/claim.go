package claim

import "github.com/roundpartner/seq/buffer"

type Item struct {
    Id int `json:"id"`
    Body string `json:"body"`
}

type Clm struct {
    elastic *Elastic
    sb buffer.BaseBuffer
    counter int
}

func NewC(elastic *Elastic, sb buffer.BaseBuffer) *Clm {
    c := &Clm{elastic, sb, 0}
    return c
}

func (claim *Clm) Next() (Item, bool) {
    body, ok := claim.sb.Pop()
    if false == ok {
        return Item{}, false
    }
    claim.counter++
    item := Item{Id: claim.counter, Body: body}
    claim.elastic.In <- item
    return item, true
}
