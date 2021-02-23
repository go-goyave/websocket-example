package chat

import (
	"context"
	"sync"

	"github.com/System-Glitch/goyave/v3"
	"github.com/System-Glitch/goyave/v3/websocket"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]struct{}

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

// NewHub create a new chat Hub.
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]struct{}),
	}
}

// Run the Hub loop. Should be run in a goroutine. This method
// should not be called more than once.
//
// Registers a shutdown hook, ensuring the hub shutdowns properly
// and closes all active connections before goyave.Start() returns.
func (h *Hub) Run() {
	done := make(chan struct{}, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer goyave.Logger.Println("run return")
	goyave.RegisterShutdownHook(func() {
		goyave.Logger.Println("shutdown hook")
		cancel()
		<-done
		goyave.Logger.Println("done received")
	})
	goyave.Logger.Println("hub run")
	for {
		select {
		case <-ctx.Done():
			done <- struct{}{}
			// TODO graceful shutdown of all active connections
			return
		case client := <-h.register:
			h.clients[client] = struct{}{}
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// Serve is the websocket Handler for the chat clients.
func (h *Hub) Serve(c *websocket.Conn, request *goyave.Request) error {
	client := &Client{
		Name:      request.String("name"),
		hub:       h,
		conn:      c,
		send:      make(chan []byte, 256),
		readErr:   make(chan error, 1),
		writeErr:  make(chan error, 1),
		waitGroup: sync.WaitGroup{},
	}
	h.broadcast <- []byte(client.Name + " joined.")
	err := client.pump()
	h.broadcast <- []byte(client.Name + " left.")
	return err
}
