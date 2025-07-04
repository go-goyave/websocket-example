package chat

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/websocket"

	ws "github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {

	// Buffered channel of outbound messages.
	send chan []byte

	readErr  chan error
	writeErr chan error

	// The websocket connection.
	conn *websocket.Conn

	hub *Hub

	Name string
}

func (c *Client) pump() error {
	c.hub.register <- c
	go c.writePump()
	go c.readPump()

	var err error
	select {
	case e := <-c.readErr:
		err = e
	case e := <-c.writeErr:
		err = e
		if err == nil {
			// Hub closing, wait for readPump to return
			<-c.readErr
		}
	}

	return err
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
	}()
	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { return c.conn.SetReadDeadline(time.Now().Add(pongWait)) })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err) {
				c.readErr <- err
				return
			}
			c.readErr <- errors.Errorf("read: %w", err)
			return
		}
		message = bytes.TrimSpace(bytes.ReplaceAll(message, newline, space))
		c.hub.Broadcast([]byte(fmt.Sprintf("%s: %s", c.Name, message)))
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.writeErr <- nil
				return
			}

			w, err := c.conn.NextWriter(ws.TextMessage)
			if err != nil {
				c.writeErr <- errors.Errorf("next writer: %w", err)
				return
			}
			if !c.write(w, message) {
				return
			}

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				if !c.write(w, newline) {
					return
				}
				if !c.write(w, <-c.send) {
					return
				}
			}

			if err := w.Close(); err != nil {
				c.writeErr <- errors.Errorf("writer close: %w", err)
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(ws.PingMessage, nil); err != nil {
				c.writeErr <- errors.Errorf("ping: %w", err)
				return
			}
		}
	}
}

func (c *Client) write(w io.Writer, p []byte) bool {
	_, err := w.Write(p)
	if err != nil {
		c.writeErr <- errors.Errorf("write: %w", err)
		return false
	}
	return true
}
