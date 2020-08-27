package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

type KafkaConnection struct {
	conn *kafka.Conn
}
type KafkaConfig struct {
	address   string
	topic     string
	partition int
	deadline  int
}

func NewKafkaWriter(ctx context.Context, c KafkaConfig) (*KafkaConnection, error) {
	k := new(KafkaConnection)
	conn, err := kafka.DialLeader(ctx, "tcp", c.address, c.topic, c.partition)
	if err != nil {
		return nil, err
	}
	k.conn = conn

	return k, nil
}

func (k *KafkaConnection) Send(deadline int, messages ...kafka.Message) error {
	if err := k.conn.SetWriteDeadline(time.Now().Add(time.Duration(deadline) * time.Second)); err != nil {
		return err
	}

	if _, err := k.conn.WriteMessages(messages...); err != nil {
		return err
	}

	return nil
}

func (k *KafkaConnection) Close() {
	k.conn.Close()
}
