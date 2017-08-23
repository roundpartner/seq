package buffer

type Message struct {
	Content string
}

func Create(size int) chan Message {
	messages := make(chan Message, size)
	return messages
}

func Add(messages chan Message, Content string) bool {
	messages <- Message{Content: Content}
	return true
}

func Pop(messages chan Message) string {
	message := <-messages
	return message.Content
}
