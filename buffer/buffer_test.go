package buffer

import (
	"testing"
)

func TestNew(t *testing.T) {
	sb := New(1)
	message := Message{Content: "Hello World"}
    sb.Messages <- message
	if <-sb.Messages != message {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	sb := New(1)
    result := sb.Add("Hello World")
	if true != result {
		t.Fail()
	}
}

func TestAddTwoFails(t *testing.T) {
    sb := New(1)
    sb.Add("Hello World")
    result := sb.Add("Hello World")
	if false != result {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
    sb := New(1)
    sb.Add("Hello World")
	message, _ := sb.Pop()
	if "Hello World" != message {
		t.Fail()
	}
}

func TestPopTwoFails(t *testing.T) {
    sb := New(1)
    sb.Add("Hello World")
    sb.Pop()
    _, ok := sb.Pop()
    if false != ok {
        t.Fail()
    }
}