<p align="center">
    <img src="./.github/img/goyave_banner.png#gh-light-mode-only" alt="Goyave Logo" width="550"/>
    <img src="./.github/img/goyave_banner_dark.png#gh-dark-mode-only" alt="Goyave Logo" width="550"/>
</p>

# Goyave Websocket Example

![https://github.com/go-goyave/websocket-example/actions](https://github.com/go-goyave/websocket-example/workflows/Test/badge.svg)

A minimal chat application to showcase [Goyave](https://github.com/go-goyave/goyave)'s websocket feature. This project is based on [Gorilla's chat example](https://github.com/gorilla/websocket/tree/master/examples/chat).

**Disclaimer:** This example project cannot be used in a real-life scenario, as you would need to be able to serve clients across multiple instances of the application. This is a typical scenario in cloud environments. The hub in this example could use a PUB/SUB mechanism (for example with [redis](https://redis.io/docs/interact/pubsub/)) to solve this issue.

## Getting Started

### Directory structure

```
.
├── http
│   ├── controller
│   │   └── chat             // Chat hub implementation
│   └── route
│       └── route.go         // Routes definition
│
├── resources
│   └── template             // Static resources
│       └── ...
│
├── .gitignore
├── .golangci.yml            // Settings for the Golangci-lint linter
├── config.example.json      // Example config for local development
├── config.test.json         // Config file used for tests
├── go.mod
├── go.sum
└── main.go                  // Application entrypoint
```

### Running the project

First, make your own configuration for your local environment. You can copy `config.example.json` to `config.json`.

Run `go run main.go` in your project's directory to start the server, then open your browser to `http://localhost:8080`.

## Resources

- [Documentation](https://goyave.dev)
- [go.pkg.dev](https://pkg.go.dev/goyave.dev/goyave/v5)
