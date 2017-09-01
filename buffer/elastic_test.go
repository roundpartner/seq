package buffer

import (
    "testing"
    "runtime"
    "time"
)

func TestNew(t *testing.T) {
    sb := New()
    message := Message{Content: "Test New"}
    sb.In <- message
    runtime.Gosched()
    if <-sb.Out != message {
        t.Fail()
    }
}

func TestAdd(t *testing.T) {
    sb := New()
    result := sb.Add("Test Add")
    if true != result {
        t.Fail()
    }
}

func TestAddTwo(t *testing.T) {
    sb := New()
    sb.Add("Test Add Two")
    result := sb.Add("Test Add Two")
    if true != result {
        t.Fail()
    }
}

func TestPopOut(t *testing.T) {
    sb := New()
    msg := Message{Content: "Test Pop Out"}
    sb.Out <- msg
    message, _ := sb.Pop()
    if "Test Pop Out" != message {
        t.Fail()
    }
}

func TestPopIn(t *testing.T) {
    sb := New()
    msg := Message{Content: "Test Pop In"}
    sb.In <- msg
    runtime.Gosched()
    message, _ := sb.Pop()
    if "Test Pop In" != message {
        t.Fail()
    }
}

func TestPop(t *testing.T) {
    sb := New()
    sb.Add("Test Pop")
    runtime.Gosched()
    message, _ := sb.Pop()
    if "Test Pop" != message {
        t.Fail()
    }
}

func TestPopTwoFails(t *testing.T) {
    sb := New()
    sb.Add("Hello World")
    sb.Pop()
    _, ok := sb.Pop()
    if false != ok {
        t.Fail()
    }
}

func TestPopMany(t *testing.T) {
    sb := New()
    for i := 1; i <= 30; i++ {
        runtime.Gosched()
        sb.Add("Test Pop Many")
        time.Sleep(time.Millisecond * 10)
    }
    for len(sb.buffer) != 0 {
        runtime.Gosched()
        message, _ := sb.Pop()
        if "Test Pop Many" != message {
            t.Fail()
        }
    }
}