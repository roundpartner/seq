package claim

import (
    "testing"
    "github.com/roundpartner/seq/buffer"
)

func TestNext(t *testing.T) {
    e, sb := reset(1)
    sb.Add("Hello World")
    _, ok := Next(e, sb)
    if false == ok {
        t.Fail()
    }
}

func TestNextHasId(t *testing.T) {
    e, sb := reset(1)
    sb.Add("Hello World")
    c, _ := Next(e, sb)
    if 1 != c.Id {
        t.Fail()
    }
}

func TestNextHasBody(t *testing.T) {
    e, sb := reset(1)
    sb.Add("Hello World")
    c, _ := Next(e, sb)
    if "Hello World" != c.Body {
        t.Fail()
    }
}

func TestNextHasIncrementingId(t *testing.T) {
    e, sb := reset(2)
    sb.Add("Hello World")
    sb.Add("Hello World")
    Next(e, sb)
    c, _ := Next(e, sb)
    if 2 != c.Id {
        t.Errorf("got: %d", c.Id)
        t.Fail()
    }
}

func TestNextInsertsAddsClaim(t *testing.T) {
    e, sb := reset(1)
    sb.Add("Hello World")
    c, _ := Next(e, sb)
    if <-e.Out != c {
        t.Fail()
    }
}

func reset(size int) (*Elastic, *buffer.SimpleBuffer) {
    id = 0
    sb := buffer.New(size)
    return New(), sb
}