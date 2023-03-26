package handler

import (
	"fmt"
	"github.com/jefersonalmeida/go-wallet/pkg/events"
	"github.com/jefersonalmeida/go-wallet/pkg/kafka"
	"sync"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		Kafka: kafka,
	}
}

func (h *TransactionCreatedKafkaHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Publish(event, nil, "transactions")
	fmt.Println("TransactionCreatedKafkaHandler: ", event.GetPayload())
}
