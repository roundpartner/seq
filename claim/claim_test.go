package claim

import (
    "testing"
    "github.com/roundpartner/seq/buffer"
)

func TestNext(t *testing.T) {
    buf := reset(1)
    buffer.Add(buf, "Hello World")
    _, ok := Next(buf)
    if false == ok {
        t.Fail()
    }
}

func TestNextHasId(t *testing.T) {
    buf := reset(1)
    buffer.Add(buf, "Hello World")
    c, _ := Next(buf)
    if 1 != c.Id {
        t.Fail()
    }
}

func TestNextHasBody(t *testing.T) {
    buf := reset(1)
    buffer.Add(buf, "Hello World")
    c, _ := Next(buf)
    if "Hello World" != c.Body {
        t.Fail()
    }
}

func TestNextHasIncrementingId(t *testing.T) {
    buf := reset(2)
    buffer.Add(buf, "Hello World")
    buffer.Add(buf, "Hello World")
    Next(buf)
    c, _ := Next(buf)
    if 2 != c.Id {
        t.Errorf("got: %d", c.Id)
        t.Fail()
    }
}

func reset(size int) chan buffer.Message {
    id = 0
    return buffer.Create(size)
}