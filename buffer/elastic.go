package buffer

type ElasticBuffer struct {
    In chan Message
    Out chan Message
    buffer []Message
}

func New() *ElasticBuffer {
    sb := &ElasticBuffer{
        In: make(chan Message, 10),
        Out: make(chan Message, 10),
        buffer: make([]Message, 0),
    }
    go sb.run()
    return sb
}

func (sb *ElasticBuffer) run() {
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

func (sb *ElasticBuffer) Add(Content string) bool {
    select {
        case sb.In <- Message{Content: Content}:
            return true
        default:
            sb.In <- Message{Content: Content}
            return true
    }
}

func (sb *ElasticBuffer) Pop() (string, bool) {
    select {
        case message := <- sb.Out:
            return message.Content, true
        default:
            return "", false
    }
}

