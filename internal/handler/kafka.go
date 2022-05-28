package handler

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9882"},
		Topic:    "mytopic",
		GroupID:  "g1",
		MaxBytes: 18,
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Some error occurred", err)
			continue
		}
		fmt.Println("Message is: ", string(m.Value))
	}

}
