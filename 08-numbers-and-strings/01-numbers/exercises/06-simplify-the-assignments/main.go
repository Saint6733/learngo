// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Simplify the Assignments
//
//  Simplify the code (refactor)
//
// RESTRICTION
//  Use only the incdec and assignment operations
//
// EXPECTED OUTPUT
//  3
// ---------------------------------------------------------

func main() {
	width, height := 10, 2

	// 增加宽度
	width += 1
	// 增加高度
	width += height
	// 减少宽度
	width -= 1
	// 减少高度
	width -= height
	// 宽度乘以20
	width *= 20
	// 宽度除以25
	width /= 25
	// 宽度模5
	width %= 5

	fmt.Println(width)
}
