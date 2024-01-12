package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/go-goyave/websocket-example/http/route"
	"github.com/go-goyave/websocket-example/service/static"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/fsutil"
)

//go:embed resources
var resources embed.FS

func main() {

	resourcesEmbed := fsutil.NewEmbed(resources)

	server, err := goyave.New(goyave.Options{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err.(*errors.Error).String())
		os.Exit(1)
	}

	server.Logger.Info("Registering hooks")
	server.RegisterSignalHook()

	server.RegisterStartupHook(func(s *goyave.Server) {
		server.Logger.Info("Server is listening", "host", s.Host())
	})

	server.RegisterShutdownHook(func(s *goyave.Server) {
		s.Logger.Info("Server is shutting down")
	})

	server.Logger.Info("Registering services")
	staticResources, err := resourcesEmbed.Sub("resources/template")
	if err != nil {
		server.Logger.Error(err)
		os.Exit(1)
	}
	server.RegisterService(static.NewService(staticResources))

	server.Logger.Info("Registering routes")
	server.RegisterRoutes(route.Register)

	if err := server.Start(); err != nil {
		server.Logger.Error(err)
		os.Exit(3)
	}
}
