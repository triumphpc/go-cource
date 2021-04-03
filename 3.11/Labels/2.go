package main

import "fmt"

func main() {

	var w = "One"
	var l = "Two"

	for _, w := range w {
		fmt.Printf("%c", w)
	w:
		for _, l := range l {
			fmt.Printf("%c", l)
			continue w
		}
	}
}
