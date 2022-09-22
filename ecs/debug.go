//go:build ecs.debug

package ecs

import (
	"os"
)

func init() {
	Log(os.Stdout)
}
