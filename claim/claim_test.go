package claim

import (
    "testing"
    "github.com/roundpartner/seq/buffer"
)

func TestNext(t *testing.T) {
    e, sb := reset(1)
    clm := NewC(e, sb)
    sb.Add("Hello World")
    _, ok := clm.Next()
    if false == ok {
        t.Fail()
    }
}

func TestNextHasId(t *testing.T) {
    e, sb := reset(1)
    clm := NewC(e, sb)
    sb.Add("Hello World")
    c, _ := clm.Next()
    if 1 != c.Id {
        t.Fail()
    }
}

func TestNextHasBody(t *testing.T) {
    e, sb := reset(1)
    clm := NewC(e, sb)
    sb.Add("Hello World")
    c, _ := clm.Next()
    if "Hello World" != c.Body {
        t.Fail()
    }
}

func TestNextHasIncrementingId(t *testing.T) {
    e, sb := reset(2)
    clm := NewC(e, sb)
    sb.Add("Hello World")
    sb.Add("Hello World")
    clm.Next()
    c, _ := clm.Next()
    if 2 != c.Id {
        t.Errorf("got: %d", c.Id)
        t.Fail()
    }
}

func TestNextInsertsAddsClaim(t *testing.T) {
    e, sb := reset(1)
    clm := NewC(e, sb)
    sb.Add("Hello World")
    c, _ := clm.Next()
    if <-e.Out != c {
        t.Fail()
    }
}

func reset(size int) (*Elastic, *buffer.SimpleBuffer) {
    sb := buffer.New(size)
    return New(), sb
}