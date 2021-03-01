package chat

import (
	"context"
	"sync"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/websocket"
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

	ctx    context.Context
	cancel context.CancelFunc
}

// NewHub create a new chat Hub.
func NewHub() *Hub {
	ctx, cancel := context.WithCancel(context.Background())
	return &Hub{
		broadcast:  make(chan []byte, 256),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]struct{}),
		ctx:        ctx,
		cancel:     cancel,
	}
}

// Run the Hub loop. Should be run in a goroutine. This method
// should not be called more than once.
//
// Registers a shutdown hook, ensuring the hub shutdowns properly
// and closes all active connections before goyave.Start() returns.
func (h *Hub) Run() {
	done := make(chan struct{}, 1)
	defer h.cancel()
	goyave.RegisterShutdownHook(func() {
		h.cancel()
		<-done
	})
	for {
		select {
		case <-h.ctx.Done():
			wg := &sync.WaitGroup{}
			wg.Add(len(h.clients))
			for client := range h.clients {
				delete(h.clients, client)
				go func(c *Client) {
					close(c.send)
					if err := c.conn.CloseNormal(); err != nil {
						goyave.ErrLogger.Println(err)
					}
					<-h.unregister // Wait for readPump to return
					wg.Done()
				}(client)
			}
			wg.Wait()
			done <- struct{}{}
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
		Name:     request.String("name"),
		hub:      h,
		conn:     c,
		send:     make(chan []byte, 256),
		readErr:  make(chan error, 1),
		writeErr: make(chan error, 1),
	}
	h.Broadcast([]byte(client.Name + " joined."))
	err := client.pump()
	h.Broadcast([]byte(client.Name + " left."))

	return err
}

// Broadcast send a message to all connected clients. This method is concurrently safe
// and doesn't do anything is the Hub's context is canceled.
func (h *Hub) Broadcast(message []byte) {
	select {
	case <-h.ctx.Done(): // Don't send if the hub is shutting down
	case h.broadcast <- message:
	}
}
