<p align="center">
    <img src="https://raw.githubusercontent.com/System-Glitch/goyave/master/resources/img/logo/goyave_text.png" alt="Goyave Logo" width="550"/>
</p>

# Goyave Websocket Example

## 🚧 Work in progress

A minimal chat application to showcase [Goyave](https://github.com/System-Glitch/goyave)'s websocket feature.

## Getting Started

### Requirements

- Go 1.13+
- Go modules

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

Run `go run main.go` in your project's directory to start the server, then open your browser to `http://localhost:8080/index.html`.

## Learning Goyave

The Goyave framework has an extensive documentation covering in-depth subjects and teaching you how to run a project using Goyave from setup to deployment.

<a href="https://goyave.dev/guide/installation"><h3 align="center">Read the documentation</h3></a>

<a href="https://pkg.go.dev/github.com/System-Glitch/goyave/v3"><h3 align="center">pkg.go.dev</h3></a>

## Contributing

Thank you for considering contributing to the Goyave framework! You can find the contribution guide in the [documentation](https://goyave.dev/guide/contribution-guide.html).

I have many ideas for the future of Goyave. I would be infinitely grateful to whoever want to support me and let me continue working on Goyave and making it better and better.

You can support me on Github Sponsor or Patreon.

<a href="https://github.com/sponsors/System-Glitch">❤ Sponsor me!</a>

<a href="https://www.patreon.com/bePatron?u=25997573">
    <img src="https://c5.patreon.com/external/logo/become_a_patron_button@2x.png" width="160">
</a>

## License

This example project is MIT Licensed. Copyright © 2021 Jérémy LAMBERT (SystemGlitch) 

The Goyave framework is MIT Licensed. Copyright © 2019 Jérémy LAMBERT (SystemGlitch)
