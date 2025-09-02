package utils

import (
	"fmt"
	"runtime"
)

func Assert(cond bool, msg string) {
	if !cond {
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, false)
		panic(fmt.Sprintf("Assertion failed: %s\n%s", msg, buf))
	}
}
