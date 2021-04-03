package main

import "fmt"

func main() {

	for {
	run:
		fmt.Println("yep")

		switch {
		case true:
			continue run
		}
	}
}
