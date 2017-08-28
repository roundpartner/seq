package buffer

type Message struct {
	Content string
}

type SimpleBuffer struct {
    Messages chan Message
}

func New(size int) *SimpleBuffer {
    return &SimpleBuffer{
        Messages: make(chan Message, size),
    }
}

func (sb *SimpleBuffer) Add(Content string) bool {
    select {
        case sb.Messages <- Message{Content: Content}:
            return true
        default:
            return false
    }
}

func (sb *SimpleBuffer) Pop() (string, bool) {
    select {
        case message := <- sb.Messages:
            return message.Content, true
        default:
            return "", false
    }
}

func Pop(messages chan Message) (string, bool) {
    select {
        case message := <-messages:
            return message.Content, true
        default:
            return "", false
    }

}
