package ecs

import "math"

// Entity ID is considered invalid if set to this value.
const invalid = math.MaxUint32

// Entity is a unique identifer for a game object, which zero or more components are associated with in a world.
type Entity struct {
	id      uint32
	version uint32
}
