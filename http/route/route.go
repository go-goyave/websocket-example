package route

import (
	"github.com/go-goyave/websocket-example/http/controller/chat"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/cors"
	"goyave.dev/goyave/v3/log"
	"goyave.dev/goyave/v3/websocket"
)

// Register all the application routes. This is the main route registrer.
func Register(router *goyave.Router) {

	router.CORS(cors.Default())
	router.Middleware(log.CombinedLogMiddleware())

	hub := chat.NewHub()
	go hub.Run()

	upgrader := websocket.Upgrader{}
	router.Get("/chat", upgrader.Handler(hub.Serve)).Validate(chat.JoinRequest)
	router.Static("/", "resources/template", false)
}
