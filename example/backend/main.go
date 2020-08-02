package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"math/rand"
	"sync"
	"time"
)

type Main struct {
	UserId    uint32 `json:"user_id"`
	App       string `json:"app"`
	Host      string `json:"host"`
	Event     string `json:"event"`
	Ip        string `json:"ip"`
	Guid      string `json:"guid"`
	CreatedAt string `json:"created_at"`
}

func (m Main) String() string {
	return fmt.Sprintf("Send message to broker: user %d, time %s", m.UserId, m.CreatedAt)
}

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "Topic1",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	defer r.Close()

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "Topic1",
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	wg := &sync.WaitGroup{}
	message := make(chan Main)

	for i := 0; i <= 10; i++ {
		go producer(message, wg)
	}

	go sender(message, w)
	go consumer(r)

	wg.Wait()
}

func producer(message chan<- Main, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	// or timestamp
	now := time.Now().Format("2006-01-02 15:04:05")

	message <- Main{
		UserId:    uint32(rand.Intn(30-10) + 10),
		App:       "",
		Host:      "",
		Event:     "",
		Ip:        "",
		Guid:      "",
		CreatedAt: now,
	}
}

func sender(message <-chan Main, w *kafka.Writer) {
	for {
		msg := <-message
		jsonbytes, err := json.Marshal(msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte("Key-A"),
				Value: jsonbytes,
			},
		)

		fmt.Println(msg)
	}
}

func consumer(r *kafka.Reader) {
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
