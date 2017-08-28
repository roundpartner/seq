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

func Add(messages chan Message, Content string) bool {
    select {
        case messages <- Message{Content: Content}:
            return true
        default:
            return false
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
