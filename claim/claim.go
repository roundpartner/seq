package claim

import "github.com/roundpartner/seq/buffer"

type Item struct {
    Id int `json:"id"`
    Body string `json:"body"`
}

var id int = 0

func Next(elastic *Elastic, buf chan buffer.Message) (Item, bool) {
    body, ok := buffer.Pop(buf)
    if false == ok {
        return Item{}, false
    }
    id++
    item := Item{Id: id, Body: body}
    elastic.In <- item
    return item, true
}
