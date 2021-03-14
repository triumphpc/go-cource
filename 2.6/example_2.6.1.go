/**
Go (Golang) From simple to great. The Complete Developer's Guide.
Example 2.5: Standard library packages

@author Alex Versus 2021
*/

package main

// Import packages format
import (
	"fmt"
	"math" // import package math from standard library
	"strings"
)

// Entrypoint for program
func main() {
	// Run function Println from fmt package
	fmt.Println("This is Println function exec")
	// Run Round function from math package
	math.Round(2.3)
	// Run function TrimSpace function from strings package for some string
	strings.TrimSpace(" 'Some text trimmed' ")
}
