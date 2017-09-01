package buffer

type Message struct {
	Content string
}

type BaseBuffer interface {
    Add(Content string) bool
    Pop() (string, bool)
}

func Add(bb BaseBuffer, Content string) (bool) {
    return bb.Add(Content)
}

func Pop(bb BaseBuffer) (string, bool) {
    return bb.Pop()
}

type SimpleBuffer struct {
    Buffer chan Message
}

func NewSimpleBuffer() *SimpleBuffer {
    sb := &SimpleBuffer{
        Buffer: make(chan Message, 10),
    }
    return sb
}

func (sb SimpleBuffer) Add(Content string) (bool) {
    msg := Message{Content: Content}
    sb.Buffer <- msg
    return true
}

func (sb SimpleBuffer) Pop() (string, bool) {
    select {
        case msg := <- sb.Buffer:
            return msg.Content, true
        default:
            return "", false

    }
}

