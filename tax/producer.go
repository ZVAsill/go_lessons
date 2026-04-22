package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Message struct {
	UserID int    `json:"user_id"`
	Action string `json:"action"`
}

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "demo-topic",
	})

	defer w.Close()

	msg := Message{
		UserID: 1,
		Action: "login",
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Value: data,
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		log.Println("sent", i)
		time.Sleep(time.Second)
	}
}
