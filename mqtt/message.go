package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

type message struct {
	m mqtt.Message
}

func newMessage(m mqtt.Message) *message {
	return &message{
		m: m,
	}
}

func (m *message) Ack() {
	m.m.Ack()
}

func (m *message) Topic() string {
	return m.m.Topic()
}

func (m *message) Body() []byte {
	return m.m.Payload()
}
