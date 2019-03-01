package mqtt

import (
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	recv "github.com/vcamvr/receiver"
)

var (
	qos              = byte(1)
	tokenWaitTimeout = 3 * time.Second
)

type receiver struct {
	c       mqtt.Client
	topic   string
	message chan recv.Message
}

func NewReceiver(client mqtt.Client, topic string) (*receiver, error) {
	receiver := &receiver{
		c:       client,
		topic:   topic,
		message: make(chan recv.Message),
	}

	token := client.Subscribe(topic, qos, func(mc mqtt.Client, mm mqtt.Message) {
		receiver.message <- newMessage(mm)
	})

	if token.WaitTimeout(tokenWaitTimeout); token.Error() != nil {
		return nil, token.Error()
	}

	return receiver, nil
}

func (r *receiver) Message() <-chan recv.Message {
	return r.message
}

func (r *receiver) Close() {
	token := r.c.Unsubscribe(r.topic)
	token.WaitTimeout(tokenWaitTimeout)
	close(r.message)
}
