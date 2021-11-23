<p align="center">
    <img src="https://raw.githubusercontent.com/go-goyave/goyave/master/resources/img/logo/goyave_text.png" alt="Goyave Logo" width="550"/>
</p>

# Goyave Websocket Example

![https://github.com/go-goyave/websocket-example/actions](https://github.com/go-goyave/websocket-example/workflows/Test/badge.svg)

A minimal chat application to showcase [Goyave](https://github.com/go-goyave/goyave)'s websocket feature. This project is based on [Gorilla's chat example](https://github.com/gorilla/websocket/tree/master/examples/chat).

## Getting Started

### Directory structure

```
.
├── http
│   ├── controller           // Business logic of the application
│   │   └── ...
│   └── route
│       └── route.go         // Routes definition
│
├── resources
│   └── template             // Static resources
│       └── ...
│
├── test                     // Functional tests
|   └── ...
|
├── .gitignore
├── .golangci.yml            // Settings for the Golangci-lint linter
├── config.example.json      // Example config for local development
├── config.test.json         // Config file used for tests
├── go.mod
└── main.go                  // Application entrypoint
```

### Running the project

First, make your own configuration for your local environment. You can copy `config.example.json` to `config.json`.

Run `go run main.go` in your project's directory to start the server, then open your browser to `http://localhost:8080`.

## Learning Goyave

The Goyave framework has an extensive documentation covering in-depth subjects and teaching you how to run a project using Goyave from setup to deployment.

<a href="https://goyave.dev/guide/installation"><h3 align="center">Read the documentation</h3></a>

<a href="https://pkg.go.dev/goyave.dev/goyave/v4"><h3 align="center">pkg.go.dev</h3></a>

## License

This example project is MIT Licensed. Copyright © 2021 Jérémy LAMBERT (SystemGlitch) 

The Goyave framework is MIT Licensed. Copyright © 2019 Jérémy LAMBERT (SystemGlitch)
