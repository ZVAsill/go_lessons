package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
)

type Message struct {
	UserID int    `json:"user_id"`
	Action string `json:"action"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// ловим Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "demo-topic",
		GroupID: "group-1",
	})

	defer func() {
		log.Println("closing kafka reader...")
		r.Close()
	}()

	// goroutine для сигнала остановки
	go func() {
		<-sigChan
		log.Println("shutdown signal received")
		cancel()
	}()

	log.Println("consumer started...")

	for {
		select {
		case <-ctx.Done():
			log.Println("context cancelled, exiting loop")
			return

		default:
			msg, err := r.ReadMessage(ctx)
			if err != nil {
				log.Println("read error:", err)
				continue
			}

			var m Message

			err = json.Unmarshal(msg.Value, &m)
			if err != nil {
				log.Println("invalid json:", err)
				continue
			}

			log.Printf("user=%d action=%s\n", m.UserID, m.Action)
		}
	}
}
