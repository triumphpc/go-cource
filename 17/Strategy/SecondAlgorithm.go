package main

import "fmt"

// defined filters for concrete algorithm
type filterFoSecondAlgorithm struct{}

func (f filterFoSecondAlgorithm) doSearch(filters map[string]int) {
	fmt.Println("Second implements strategy", filters)
}
