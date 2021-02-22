package main

import (
	"os"

	"github.com/go-goyave/websocket-example/http/route"

	"github.com/System-Glitch/goyave/v3"
)

func main() {
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
