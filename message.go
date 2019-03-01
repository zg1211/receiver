package receiver

type Message interface {
	Ack()
	Topic() string
	Body() []byte
}
