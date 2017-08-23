package buffer

type Message struct {
	Content string
}

func Create(size int) chan Message {
	messages := make(chan Message, size)
	return messages
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
