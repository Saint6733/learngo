package mypackage

import (
	"fmt"
	"runtime"
)

func Hello() {
	fmt.Println(runtime.Version())
}
