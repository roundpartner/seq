package buffer

import (
	"testing"
    "runtime"
)

func TestNew(t *testing.T) {
	sb := New()
	message := Message{Content: "Test New"}
    sb.In <- message
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