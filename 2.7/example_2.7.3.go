/**
Go (Golang) From simple to great. The Complete Developer's Guide.
Example 2.7: using fmt.Sprintf

@author Alex Versus 2021
*/

package main

import "fmt"

// Entrypoint for program
func main() {
	// We use formatting float and return value
	resultString := fmt.Sprintf("with two digits after the decimal point: %0.2f\n", 1.0/5.0)
	fmt.Println("It's Println", resultString)
}
