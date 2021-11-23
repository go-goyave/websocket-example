package main

import (
	"os"

	"github.com/go-goyave/websocket-example/http/route"

	"goyave.dev/goyave/v4"
)

func main() {
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
