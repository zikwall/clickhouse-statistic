package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
	return fmt.Sprintf(
		"Send message to broker event %s: user %d, time %s from host %s & app %s",
		m.Event, m.UserId, m.CreatedAt, m.Host, m.App,
	)
}

func SetupCloseHandler(r *kafka.Reader, w *kafka.Writer) {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig

		fmt.Println("Close by signal")

		err := w.Close()

		if err != nil {
			fmt.Println(err)
		}

		err = r.Close()

		if err != nil {
			fmt.Println(err)
		}

		// todo wait last gorutines

		os.Exit(0)
	}()
}

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "ClickhouseTopic",
		GroupID: "my-group",
	})
	defer r.Close()

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "ClickhouseTopic",
	})
	defer w.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)

	message := make(chan Main)
	ctx := context.Background()
	//SetupCloseHandler(r, w)

	go func() {
		for i := 0; i <= 1000; i++ {
			// or timestamp
			now := time.Now().Format("2006-01-02 15:04:05")
			msg := Main{
				UserId:    user(),
				App:       faker(apps),
				Host:      faker(hosts),
				Event:     faker(events),
				Ip:        "",
				Guid:      "",
				CreatedAt: now,
			}

			producer(ctx, w, message, msg)
		}
	}()

	go lookup(message)
	//go consumer(ctx, r)

	wg.Wait()
}

func producer(ctx context.Context, w *kafka.Writer, message chan<- Main, msg Main) {
	message <- msg

	jsonbytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = w.WriteMessages(ctx,
		kafka.Message{
			Value: jsonbytes,
		},
	)

	if err != nil {
		fmt.Println(err)
	}
}

func lookup(message <-chan Main) {
	for msg := range message {
		fmt.Println(msg)
	}
}

func consumer(ctx context.Context, r *kafka.Reader) {
	for {
		m, err := r.ReadMessage(ctx)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
