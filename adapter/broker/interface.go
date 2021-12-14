package broker

type ProducerInterface interface {
	Publish(message interface{}, key []byte, topic string) error
}
