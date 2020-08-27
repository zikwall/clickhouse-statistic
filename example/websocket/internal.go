package main

import (
	"github.com/gofiber/websocket"
	"log"
)

type (
	client   struct{}
	Internal struct {
		Handlers
		clients    map[*websocket.Conn]client
		register   chan *websocket.Conn
		unregister chan *websocket.Conn
		event      chan string
	}
	Handlers struct {
		eventHandler func(string) error
		errorHandler func(error)
	}
)

func NewInternal() *Internal {
	i := new(Internal)
	i.clients = map[*websocket.Conn]client{}
	i.register = make(chan *websocket.Conn)
	i.unregister = make(chan *websocket.Conn)
	i.event = make(chan string)

	return i
}

func (i *Internal) SetEventHandler(handler func(string) error) {
	i.eventHandler = handler
}

func (i *Internal) SetErrorHandler(handler func(error)) {
	i.errorHandler = handler
}

func (i Internal) handleEvent(event string) {
	if err := i.eventHandler(event); err != nil {
		i.errorHandler(err)
	}
}

func (i *Internal) handleConnect(connection *websocket.Conn) {
	i.clients[connection] = client{}
	log.Println("connection registered")
}

func (i *Internal) handleDisconnect(disconnection *websocket.Conn) {
	delete(i.clients, disconnection)
	log.Println("connection unregistered")
}

func (i *Internal) Listen() {
	for {
		select {
		case connection := <-i.register:
			i.handleConnect(connection)
		case event := <-i.event:
			i.handleEvent(event)
		case connection := <-i.unregister:
			i.handleDisconnect(connection)
		}
	}
}

func (i Internal) Event(event string) {
	i.event <- event
}

func (i Internal) Connect(connection *websocket.Conn) {
	i.register <- connection
}

func (i Internal) Disconnect(connection *websocket.Conn) {
	i.unregister <- connection
}
