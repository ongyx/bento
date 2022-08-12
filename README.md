# Bento

A game framework for [Ebitengine].

**DISCLAIMER**: Bento is still in the early stages of development! The API is subject to breaking changes.

## Rationale

Bento was created mainly to make gamedev with Ebitengine easier, by providing a framework of utilities to simplify common tasks.
Modularity is a key facet of Bento, allowing you to use Bento in your game's codebase as you see fit.

## Requirements

The minimum Go version is `1.18` for versions of Bento past v0.5.0.

Bento version v0.5.0 was rewritten from scratch to take advantage of newer Go features such as generics.

## Installation

`go get github.com/ongyx/bento` should be sufficient to add it as a go.mod dependency.

## Getting Started

The [API] documentation is a good place to start.

General tutorials/how-tos can be found in the [`docs`](docs/) directory.
You may want to take a look at the [`examples`](examples/) too.

## Credits

Hajime Hoshi for creating Ebiten.

## License

Bento is licensed under the MIT License.

[Ebitengine]: https://github.com/hajimehoshi/ebiten
[API]: https://pkg.go.dev/github.com/ongyx/bento
