package buffer

import (
	"testing"
)

func TestCreate(t *testing.T) {
	messages := Create(1)
	message := Message{Content: "Hello World"}
	messages <- message
	if <-messages != message {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	messages := Create(1)
	result := Add(messages, "Hello World")
	if true != result {
		t.Fail()
	}
}

func TestAddTwo(t *testing.T) {
	messages := Create(2)
	Add(messages, "Hello World")
	result := Add(messages, "Hello World")
	if true != result {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	messages := Create(1)
	Add(messages, "Hello World")
	message := Pop(messages)
	if "Hello World" != message {
		t.Fail()
	}
}