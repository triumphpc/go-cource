package main

import (
	"fmt"
	"github.com/triumphpc/moduleexample"
	moduleexampleSV "github.com/triumphpc/moduleexample/v2"
)

func main() {
	fmt.Println(moduleexample.Hi("Mister"))

	g, err := moduleexampleSV.Hi("Alex", "Additional info")
	if err != nil {
		panic(err)
	}
	fmt.Println(g)
}
