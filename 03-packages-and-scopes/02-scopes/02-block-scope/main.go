// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
)

func nope() string { // block scope starts

	// hello and ok are only visible here
	const ok = true
	var hello = "Hello"

	_ = hello
	return hello
} // block scope ends

func main() { // block scope starts

	// hello and ok are not visible here
	nope()
	fmt.Println(nope())

	// ERROR:
	// fmt.Println(hello, ok)

} // block scope ends
