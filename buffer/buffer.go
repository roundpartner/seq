package buffer

type Message struct {
	Content string
}

type SimpleBuffer struct {
    In chan Message
    Out chan Message
    buffer []Message
}

func New() *SimpleBuffer {
    sb := &SimpleBuffer{
        In: make(chan Message),
        Out: make(chan Message),
        buffer: make([]Message, 0),
    }
    go sb.run()
    return sb
}

func (sb *SimpleBuffer) run() {
    for {
        if len(sb.buffer) > 0 {
            select {
                case sb.Out <- sb.buffer[0]:
                    sb.buffer = sb.buffer[1:]
                case value := <- sb.In:
                    sb.buffer = append(sb.buffer, value)
            }
        } else {
            value := <-sb.In
            sb.buffer = append(sb.buffer, value)
        }
    }
}

func (sb *SimpleBuffer) Add(Content string) bool {
    select {
        case sb.In <- Message{Content: Content}:
            return true
        default:
            return false
    }
}

func (sb *SimpleBuffer) Pop() (string, bool) {
    select {
        case message := <- sb.Out:
            return message.Content, true
        default:
            return "", false
    }
}

