package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/websocket"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	// todo init inside service container
	kfk, err := NewKafkaWriter(context.Background(), KafkaConfig{
		address:   "localhost:9092",
		topic:     "ClickhouseTopic",
		partition: 0,
	})

	if err != nil {
		fmt.Println(err)
	}

	internal := NewInternal()
	// todo error pool
	internal.SetErrorHandler(func(err error) {
		fmt.Println(err)
	})
	// todo batch write to kafka
	internal.SetEventHandler(func(event string) error {
		err := kfk.Send(10, kafka.Message{
			Value: []byte(event),
		})

		return err
	})

	websocketHandler := func(c *websocket.Conn) {
		// Удаляем клиента из соединений
		defer func() {
			internal.Disconnect(c)
			_ = c.Close()
		}()

		internal.Connect(c)

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}

				return
			}

			if messageType == websocket.TextMessage {
				internal.Event(string(message))
				_ = c.WriteMessage(websocket.TextMessage, message)
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}

	app := fiber.New()
	// Делаем доступным наш простой фронтент клиент
	app.Static("/public", "./public/index.html")
	// простой middleware
	app.Use(func(c *fiber.Ctx) {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			c.Next()
		}
	})
	app.Get("/ws", websocket.New(websocketHandler))

	go internal.Listen()

	if err := app.Listen(3000); err != nil {
		fmt.Println(err)
	}
}
