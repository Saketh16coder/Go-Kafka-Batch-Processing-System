package producer

import (
	"context"
	"encoding/json"

	"Demo/model"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitProducer() {
	writer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "transactions",
		Balancer: &kafka.LeastBytes{},
	}
}

func Publish(txn model.Transaction) error {
	data, _ := json.Marshal(txn)

	return writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: data,
		},
	)
}
