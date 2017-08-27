package claim

import (
    "testing"
    "github.com/roundpartner/seq/buffer"
)

func TestNext(t *testing.T) {
    e, buf := reset(1)
    buffer.Add(buf, "Hello World")
    _, ok := Next(e, buf)
    if false == ok {
        t.Fail()
    }
}

func TestNextHasId(t *testing.T) {
    e, buf := reset(1)
    buffer.Add(buf, "Hello World")
    c, _ := Next(e, buf)
    if 1 != c.Id {
        t.Fail()
    }
}

func TestNextHasBody(t *testing.T) {
    e, buf := reset(1)
    buffer.Add(buf, "Hello World")
    c, _ := Next(e, buf)
    if "Hello World" != c.Body {
        t.Fail()
    }
}

func TestNextHasIncrementingId(t *testing.T) {
    e, buf := reset(2)
    buffer.Add(buf, "Hello World")
    buffer.Add(buf, "Hello World")
    Next(e, buf)
    c, _ := Next(e, buf)
    if 2 != c.Id {
        t.Errorf("got: %d", c.Id)
        t.Fail()
    }
}

func TestNextInsertsAddsClaim(t *testing.T) {
    e, buf := reset(1)
    buffer.Add(buf, "Hello World")
    c, _ := Next(e, buf)
    if <-e.Out != c {
        t.Fail()
    }
}

func reset(size int) (*Elastic, chan buffer.Message) {
    id = 0
    return New(), buffer.Create(size)
}