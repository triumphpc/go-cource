/**
Go (Golang) From simple to great. The Complete Developer's Guide.
Example 17: Design patterns. Strategy

@author Alex Versus 2021
*/

package main

// define interface for strategy classes
type Strategy interface {
	// define method for doing search algorithm
	doSearch(filters map[string]int)
}
