package ecs

import (
	"io"
	"log"
)

var logger *log.Logger

func Log(out io.Writer) {
	logger = log.New(out, "ecs: ", log.Lshortfile)
}
