package main

import "fmt"

// Entrypoint
func main() {
	// Create new categories
	c1 := &category{"category 1"}
	c2 := &category{"category 2"}
	c3 := &category{"category 3"}

	d1 := &directory{
		children: []prototype{c1},
		name:     "Directory 1",
	}

	d2 := &directory{
		children: []prototype{d1, c2, c3},
		name:     "Directory 2",
	}

	fmt.Println("\nOpen directory 2")
	d2.show("  ")

	cloneFolder := d2.clone()
	fmt.Println("\nClone and open directory 2")
	cloneFolder.show("  ")
}
