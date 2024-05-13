package chat

import (
	"fmt"
	"sync"
	"testing"
	"time"

	ws "github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/middleware/parse"
	"goyave.dev/goyave/v5/util/testutil"
	"goyave.dev/goyave/v5/websocket"
)

func TestChat(t *testing.T) {
	server := testutil.NewTestServer(t, "config.test.json")

	hub := NewHub()
	server.RegisterRoutes(func(_ *goyave.Server, router *goyave.Router) {
		router.GlobalMiddleware(&parse.Middleware{})
		router.Subrouter("/chat").Controller(websocket.New(hub))
	})

	wg := sync.WaitGroup{}
	wg.Add(2)

	go hub.Run()

	server.RegisterStartupHook(func(_ *goyave.Server) {
		go func() {
			defer wg.Done()
			connectClient(t, server.Host(), "bob", func(conn *ws.Conn) {
				go func() {
					// Connect the second user when bob is already connected
					// so they can see the "alice joined." message.
					defer wg.Done()
					connectClient(t, server.Host(), "alice", func(conn *ws.Conn) {
						expectMessage(t, conn, "bob: Hi Alice!")
						assert.NoError(t, conn.WriteMessage(ws.TextMessage, []byte("What's up Bob?")))
						expectMessage(t, conn, "alice: What's up Bob?")
					})
				}()
				expectMessage(t, conn, "alice joined.")
				assert.NoError(t, conn.WriteMessage(ws.TextMessage, []byte("Hi Alice!")))
				expectMessage(t, conn, "bob: Hi Alice!")
				expectMessage(t, conn, "alice: What's up Bob?")
				expectMessage(t, conn, "alice left.")
			})
		}()
	})

	go func() {
		assert.NoError(t, server.Start())
	}()
	defer server.Stop()

	wg.Wait()
}

func connectClient(t *testing.T, serverHost, name string, f func(conn *ws.Conn)) {
	addr := fmt.Sprintf("ws://%s/chat?name=%s", serverHost, name)
	conn, _, err := ws.DefaultDialer.Dial(addr, nil)
	assert.NoError(t, err)
	defer func() { _ = conn.Close() }()

	f(conn)

	m := ws.FormatCloseMessage(ws.CloseNormalClosure, "Connection closed by client")
	assert.NoError(t, conn.WriteControl(ws.CloseMessage, m, time.Now().Add(time.Second)))
}

func expectMessage(t *testing.T, conn *ws.Conn, message string) {
	messageType, data, err := conn.ReadMessage()
	assert.NoError(t, err)
	assert.Equal(t, ws.TextMessage, messageType)
	assert.Equal(t, []byte(message), data)
}
