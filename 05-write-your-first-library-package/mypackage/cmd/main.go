package main

import (
	"fmt"
	mp "github.com/inancgumus/learngo/05-write-your-first-library-package/mypackage"
)

func main() {
	mp.Hello()
	fmt.Println(mp.ReturnResult())
}
