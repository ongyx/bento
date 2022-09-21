package ecs

import "github.com/ongyx/bento/ecs/sparse"

// Pool represents a collection of components of type T.
type Pool[T any] struct {
	*sparse.Map[Entity, T]
}

// NewPool returns a new pool of (size) capacity.
func NewPool[T any](size int) Pool[T] {
	return Pool[T]{sparse.NewMap[Entity, T](size)}
}

type pool interface {
	Delete(e Entity)
}
