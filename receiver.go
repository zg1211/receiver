package receiver

type Receiver interface {
	Message() <-chan Message
	Close()
}
