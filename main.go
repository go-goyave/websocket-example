package main

import (
	"fmt"
	"os"

	"github.com/go-goyave/websocket-example/http/route"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/util/errors"
)

func main() {
	server, err := goyave.New()
	if err != nil {
		fmt.Println(err.(*errors.Error).String())
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

	server.Logger.Info("Registering routes")
	server.RegisterRoutes(route.Register)

	if err := server.Start(); err != nil {
		server.Logger.Error(err)
		os.Exit(3)
	}
}
