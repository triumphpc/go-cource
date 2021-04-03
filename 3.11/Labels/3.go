package main

import "fmt"

func main() {
	i := 0
m:
	for {
		i++

		switch i {
		case 1:
			fmt.Print("y")
			break
		case 2:
			fmt.Print("e")
			continue m
		default:
			fmt.Println("p")
			goto o
		}
	}
o:
	fmt.Println("done")
}
