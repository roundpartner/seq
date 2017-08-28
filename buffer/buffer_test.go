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
	messages := New(1)
	result := Add(messages.Messages, "Hello World")
	if true != result {
		t.Fail()
	}
}

func TestAddTwo(t *testing.T) {
	messages := New(1)
	Add(messages.Messages, "Hello World")
	result := Add(messages.Messages, "Hello World")
	if false != result {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	messages := New(1)
	Add(messages.Messages, "Hello World")
	message, _ := Pop(messages.Messages)
	if "Hello World" != message {
		t.Fail()
	}
}

func TestPopTwo(t *testing.T) {
    messages := New(1)
    Add(messages.Messages, "Hello World")
    Pop(messages.Messages)
    _, ok := Pop(messages.Messages)
    if false != ok {
        t.Fail()
    }
}