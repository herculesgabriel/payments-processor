package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/herculesgabriel/payments-processor/adapter/presenter/transaction"
	"github.com/herculesgabriel/payments-processor/domain/entity"
	"github.com/herculesgabriel/payments-processor/usecase/process_transaction"
	"github.com/stretchr/testify/assert"
)

func TestProducer_Publish(t *testing.T) {
	expectedOutput := process_transaction.TransactionDTOOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you do not have enough limit for this transaction",
	}

	configMap := ckafka.ConfigMap{"test.mock.num.brokers": 3}
	producer := NewKafkaProducer(&configMap, transaction.NewKafkaTransactionPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")

	assert.Nil(t, err)
}
