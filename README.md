# Bento

A framework for 2D games built on [Ebitengine] (née Ebiten).

**DISCLAIMER**: Bento is experimental, and it is under development. The API is subject to breaking changes.

## Rationale

Ebiten is a great game engine, and it is a solid foundation for making cross-platform games.
However, complex games would benefit from a higher-level framework to simplify some common tasks.

This is where Bento attempts to fill the gap: it provides a rendering model with objects such as entities and scenes.

Keep in mind that Bento is meant to complement, not replace Ebiten. 
You can (and should) use Ebiten directly for finer-grained control over graphics if needed.

## Getting Started

The [API] documentation is a good place to start.

Some documents that you might want to read:

- [`concepts`](docs/concepts.md): Core concepts that Bento is built upon.

For some examples, check out the [`examples`](examples/) folder.

## Install

`go get github.com/ongyx/bento`

## Credits

Hajime Hoshi for creating Ebiten.

## License

Bento is licensed under the MIT License.

[Ebitengine]: https://github.com/hajimehoshi/ebiten
[reference]: docs/README.md
[API]: https://pkg.go.dev/github.com/ongyx/bento
