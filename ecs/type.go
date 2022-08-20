package ecs

import "reflect"

// Type is an opaque identifier that represents a generic component type.
type TypeID struct {
	t reflect.Type
}

// Type returns the type ID for a generic type T.
func Type[T any]() TypeID {
	var zv T
	return TypeID{reflect.TypeOf(zv)}
}
