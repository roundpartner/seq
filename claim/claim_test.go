package claim

import (
    "testing"
    "github.com/roundpartner/seq/buffer"
    "runtime"
)

func TestNext(t *testing.T) {
    e, sb := reset()
    clm := NewC(e, sb)
    sb.Add("Test Next")
    _, ok := clm.Next()
    if false == ok {
        t.Fail()
    }
}

func TestNextHasId(t *testing.T) {
    e, sb := reset()
    clm := NewC(e, sb)
    sb.Add("Test Next")
    runtime.Gosched()
    c, _ := clm.Next()
    if 1 != c.Id {
        t.Fail()
    }
}

func TestNextHasBody(t *testing.T) {
    e, sb := reset()
    clm := NewC(e, sb)
    sb.Add("Test Next Has Body")
    runtime.Gosched()
    c, _ := clm.Next()
    if "Test Next Has Body" != c.Body {
        t.Fail()
    }
}

func TestNextHasIncrementingId(t *testing.T) {
    e, sb := reset()
    clm := NewC(e, sb)
    sb.Add("Test Next Has Incrementing Id")
    sb.Add("Test Next Has Incrementing Id")
    runtime.Gosched()
    clm.Next()
    c, _ := clm.Next()
    if 2 != c.Id {
        t.Errorf("got: %d", c.Id)
        t.Fail()
    }
}

func TestNextInsertsAddsClaim(t *testing.T) {
    e, sb := reset()
    clm := NewC(e, sb)
    sb.Add("Test Next Inserts Adds Claim")
    runtime.Gosched()
    c, _ := clm.Next()
    runtime.Gosched()
    if <-e.Out != c {
        t.Fail()
    }
}

func reset() (*Elastic, buffer.BaseBuffer) {
    sb := buffer.NewSimpleBuffer()
    return New(), sb
}