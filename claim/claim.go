package claim

import (
    "github.com/roundpartner/seq/buffer"
    "encoding/json"
    "bytes"
)

type Item struct {
    Id int `json:"id"`
    Body interface{} `json:"body"`
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

    b := bytes.NewBufferString(body).Bytes()
    var i interface{}
    err := json.Unmarshal(b, &i)
    if nil != err {
        return Item{}, false
    }

    item := Item{Id: claim.counter, Body: i}
    claim.elastic.In <- item
    return item, true
}
