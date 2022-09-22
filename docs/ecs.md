# Entity Component System

## Introduction

ECS, or Entity Component System, is a programming paradigm commonly found in gamedev which emphasizes separation of behaviour from data [[1]].

https://github.com/ongyx/bento/blob/main/examples/helloworld/main.go

## Technical Details

With the introduction of generics in Go 1.18, it is now possible to create a ECS implementation without having to use reflection for the most part.

The only reflection involved is `reflect.TypeOf` to generate a unique identifer for a generic component type:

```go
func typeof[T any]() reflect.Type {
  var zeroValue T
  return reflect.TypeOf(zeroValue)
}
```

A bitset is used with the unique identifer to efficiently query entities in the world with the correct components.
Afterwards, `Query[T]()` and `Register[T]()` uses regular type assertions to get a `Component[T]`, which is much cheaper than using plain reflection.

Part of this hybrid bitset-generics approach was inspired by this blog post [[2]], which was originally written in C++.

[1]: https://github.com/SanderMertens/ecs-faq
[2]: https://austinmorlan.com/posts/entity_component_system
