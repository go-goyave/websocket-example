package route

import (
	"github.com/go-goyave/websocket-example/http/controller/chat"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/cors"
	"goyave.dev/goyave/v5/log"
	"goyave.dev/goyave/v5/middleware/parse"
	"goyave.dev/goyave/v5/websocket"
)

// Register all the application routes. This is the main route registrer.
func Register(server *goyave.Server, router *goyave.Router) {

	router.CORS(cors.Default())
	router.GlobalMiddleware(&parse.Middleware{})
	router.GlobalMiddleware(log.CombinedLogMiddleware())

	hub := chat.NewHub(server)
	go hub.Run()

	router.Subrouter("/chat").Controller(websocket.New(hub))
	router.Static("/", "resources/template", false)
}
