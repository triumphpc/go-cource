/**
Go (Golang) From simple to great. The Complete Developer's Guide.
Example 2.19: Visibility blocks
go run example_2.19.2.go

@author Alex Versus 2021
*/

// package block
package main

import "fmt"

// function block
func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	//fmt.Println(x, у) // Compiling error

}

func f() int {
	return 1

}

func g(x int) int {
	return x

}
