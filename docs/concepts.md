# Concepts

## Component

The foundation of Bento's model is the component.
A component encapsulates some state and logic that is updated every tick by calling it's `Update` method, similar to ebiten's `Game.Update`.

Components should update their subcomponents by calling their `Update` method before using their state.

## Entity

An entity is a component that draws directly to an image.

## Animation

An animation is a entity that draws on top of a sprite/scene for a finite number of ticks.

## Transition

A transition is an animation that draws over a entity when it's entering or exiting the stage.
Transitions also control the visibility of a sprite.

## Scene

A scene is a special kind of meta-entity: it has a bunch of entities that Bento updates and renders, and entity are scripted in the `Script` method.
This can be considered analogous to a game level.

## Stage

A stage draws and changes scenes on screen, optionally with transitions.
It implements the `ebiten.Game` interface, and therefore an instance can be passed directly to `ebiten.RunGame`.
