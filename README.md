# Bento

A game engine based on [Ebiten].

**DISCLAIMER**: Bento is experimental, and it is under development. Its API is subject to breaking changes.

## Rationale

Ebiten is a great game engine, but understandably does not have a higher-level API akin to Unity to keep things simple.

This is where Bento attempts to fill the gap: it provides a rendering model with objects such as entities and scenes.
However, you can (and should) use Ebiten directly for finer-grained control over graphics if needed.

## Getting Started

The general [reference] and [API] documentation is a good place to start.

Since Bento is a library, to install it `go get github.com/ongyx/bento` should suffice.

## License

Bento is licensed under the MIT License.

[Ebiten]: https://github.com/hajimehoshi/ebiten
[reference]: docs/README.md
[API]: https://pkg.go.dev/github.com/ongyx/bento
