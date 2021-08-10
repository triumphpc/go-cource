package main

import (
	"fmt"
	"github.com/triumphpc/moduleexample/v2"
)

func main() {
	g, err := moduleexample.Hi("Alex", "Additional info")
	if err != nil {
		panic(err)
	}
	fmt.Println(g)
}
