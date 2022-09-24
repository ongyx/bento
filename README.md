# Bento 🍱

**A game framework for [Ebitengine]**

**DISCLAIMER**: Bento is still in the early stages of development! The API is subject to breaking changes.

## Rationale

TL;DR My original intent was to make a game, but I ended up creating a game framework `¯\_(ツ)_/¯`

Ebitengine is great for making cross platform games because it abstracts away the low-level stuff.
However, I think that a game framework would be beneficial for simplifying some aspects of game development with it.

To this end, Bento aims to provide a modular framework of utilities.

## Requirements

The minimum Go version is `1.18` for versions of Bento past 0.5.0.

## Getting Started

### Installation

`go get github.com/ongyx/bento` should be sufficient to add it as a go.mod dependency.

### Documentation

The [API] documentation is a good place to start.

General tutorials/how-tos can be found in the [`docs`](docs/) directory.
You may want to take a look at the [`examples`](examples/) too.

### Build Tags

Bento can be configured by specifying build tags when compiling or testing.

Some features require CGo to be enabled, which is the default for native builds.
If you are [cross compiling](#cross-compiling), some setup is needed.

Feature       | Description                               | CGo needed?
---           | ---                                       | ---
`discretegpu` | Prefer using the discrete GPU on Windows  | Y
`pprof`       | Enable pprof server on `localhost:6060`   | N
`ecs.debug`   | Enable debug logging to stdout for ECS    | N

### Cross Compiling

To cross compile pure Go code, it's generally as easy as setting `GOOS` and `GOARCH` to your target platform, such as `windows/amd64`.
However, features using CGo require some setup to work properly.

First, find and install the appropriate cross compiler for your target platform. For Linux -> Windows, this is usually MinGW.
(Consult your package manager for this.)

Next, export the `GOOS`/`GOARCH` from earlier, `CGO_ENABLED=1`, and `CC` to the cross compiler:

```bash
$ # this is for windows/amd64, change the variables accordingly for your platform
$ export GOOS=windows
$ export GOARCH=amd64
$ export CGO_ENABLED=1
$ export CC=x86_64-w64-mingw32-gcc
$ go build
```

## Credits

Hajime Hoshi for creating Ebiten.

## License

Bento is licensed under the MIT License.

[Ebitengine]: https://github.com/hajimehoshi/ebiten
[API]: https://pkg.go.dev/github.com/ongyx/bento
