package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/websocket"
	"log"
)

func main() {
	internal := NewInternal()
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
