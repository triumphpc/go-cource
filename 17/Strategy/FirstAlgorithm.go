package main

import "fmt"

// defined filters for concrete algorithm
type FilterFoFirstAlgorithm struct{}

func (f FilterFoFirstAlgorithm) doSearch(filters map[string]int) {
	fmt.Println("First implements strategy", filters)
}
