/**
Go (Golang) From simple to great. The Complete Developer's Guide.
Example 2.3:

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
	// Run function Printf from fmt package
	fmt.Println("This is Printf function exec")
	// Run Round function from math package
	fmt.Println(math.Round(2.3))
	// Run function TrimSpace function from strings package for some string
	fmt.Println(strings.TrimSpace(" 'Some text trimmed' "))
}
