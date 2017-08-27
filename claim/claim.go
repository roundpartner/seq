package claim

import "github.com/roundpartner/seq/buffer"

type Item struct {
    Id int
    Body string `json:"body"`
}

var id int = 0

func Next(buf chan buffer.Message) (Item, bool) {
    body, ok := buffer.Pop(buf)
    if false == ok {
        return Item{}, false
    }
    id++
    item := Item{Id: id, Body: body}
    return item, true
}
