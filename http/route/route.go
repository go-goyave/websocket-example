package route

import (
	"github.com/System-Glitch/goyave/v3"
	"github.com/System-Glitch/goyave/v3/cors"
	"github.com/System-Glitch/goyave/v3/log"
	"github.com/System-Glitch/goyave/v3/websocket"
	"github.com/go-goyave/websocket-example/http/controller/chat"
)

// Register all the application routes. This is the main route registrer.
func Register(router *goyave.Router) {

	router.CORS(cors.Default())
	router.Middleware(log.CombinedLogMiddleware())

	hub := chat.NewHub()
	go hub.Run()

	upgrader := websocket.Upgrader{}
	router.Get("/chat", upgrader.Handler(hub.Serve))
	router.Static("/", "resources/template", false)
}
