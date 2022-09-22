//go:build !ecs.debug

package ecs

import "io"

func init() {
	Log(io.Discard)
}
