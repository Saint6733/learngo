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
// EXERCISE: Print the literals
//
//  1. Print a few integer literals
//
//  2. Print a few float literals
//
//  3. Print true and false bool constants
//
//  4. Print your name using a string literal
//
//  5. Print a non-english sentence using a string literal
//
// ---------------------------------------------------------

func main() {
	// Use fmt.Println()
	fmt.Println(1, 2, 3, 4, 5, -5, -4, -3, -2, -1)
	fmt.Println(1.0, 2.0, 3.0, 4.0, 5.0)
	fmt.Println(true, false)
	fmt.Println("hello")
	fmt.Println("你好")
}
