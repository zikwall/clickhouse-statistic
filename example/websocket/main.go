package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/websocket"
	"log"
)

func main() {
	app := fiber.New()

	app.Static("/public", "./public/index.html")

	app.Use(func(c *fiber.Ctx) {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			c.Next()
		}
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		for {
			mt, msg, err := c.ReadMessage()

			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}

				return
			}

			if mt == websocket.TextMessage {
				log.Printf("recv: %s", msg)

				err = c.WriteMessage(mt, msg)

				if err != nil {
					log.Println("write:", err)
					break
				}
			}

			if mt == websocket.CloseMessage {
				fmt.Println("close")
			}
		}
	}))

	app.Listen(3000)
}
