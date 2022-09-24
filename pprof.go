//go:build pprof

package bento

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}
