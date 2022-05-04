package kafka_producer

import (
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

p, err := kafka.NewProducer(&kafka.ConfigMap{
    "bootstrap.servers": "host1:9092,host2:9092",
    "client.id": socket.gethostname(),
    "acks": "all"})

if err != nil {
    fmt.Printf("Failed to create producer: %s\n", err)
    os.Exit(1)
}