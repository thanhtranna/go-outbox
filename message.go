package outbox

// Message encapsulates the contents of the message to be sent
type Message struct {
	Key     string
	Headers map[string]string
	Body    []byte
	Topic   string
}
