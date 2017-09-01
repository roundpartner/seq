package buffer

import "testing"

func TestNewSimpleBuffer(t *testing.T) {
    e := NewSimpleBuffer()
    e.Buffer <- Message{}
}

func TestBaseBufferAdd(t *testing.T) {
    e := NewSimpleBuffer()
    ok := Add(e, "Hello World")
    if false == ok {
        t.Fail()
    }
}

func TestBaseBufferPop(t *testing.T) {
    e := NewSimpleBuffer()
    msg := Message{Content: "Hello World"}
    e.Buffer <- msg
    _, ok := Pop(e)
    if false == ok {
        t.Fail()
    }
}
