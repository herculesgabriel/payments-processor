package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/herculesgabriel/payments-processor/adapter/presenter"
)

type Producer struct {
	ConfigMap *ckafka.ConfigMap
	Presenter presenter.Presenter
}

func NewKafkaProducer(configMap *ckafka.ConfigMap, presenter presenter.Presenter) *Producer {
	return &Producer{ConfigMap: configMap, Presenter: presenter}
}

func (p *Producer) Publish(message interface{}, key []byte, topic string) error {
	producer, err := ckafka.NewProducer(p.ConfigMap)

	if err != nil {
		return err
	}

	err = p.Presenter.Bind(message)

	if err != nil {
		return err
	}

	presenterMessage, err := p.Presenter.Show()

	if err != nil {
		return err
	}

	kafkaMessage := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          presenterMessage,
		Key:            key,
	}

	err = producer.Produce(kafkaMessage, nil)

	if err != nil {
		return err
	}

	return nil
}
