package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/herculesgabriel/payments-processor/adapter/broker/kafka"
	"github.com/herculesgabriel/payments-processor/adapter/factory"
	"github.com/herculesgabriel/payments-processor/adapter/presenter/transaction"
	"github.com/herculesgabriel/payments-processor/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// db
	db, err := sql.Open("sqlite3", "payments_processor.db")

	if err != nil {
		log.Fatal(err)
	}

	// repository
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	// producer
	configMapProducer := &ckafka.ConfigMap{"boostrap.servers": "kafka:9092"}
	kafkaPresenter := transaction.NewKafkaTransactionPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	// consumer
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "payments-processor",
		"group.id":          "payments-processor",
	}
	topics := []string{"new_transaction"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	var messageChan = make(chan *ckafka.Message)
	go consumer.Consume(messageChan)

	// usecase
	usecase := process_transaction.NewProcessTransaction(&repository, producer, "transaction_processed")

	for message := range messageChan {
		var input process_transaction.TransactionDTOInput
		json.Unmarshal(message.Value, &input)
		usecase.Execute(input)
	}
}
