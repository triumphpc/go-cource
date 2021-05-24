/**
Go (Golang) From simple to great. The Complete Developer's Guide.
Example 6.10: Function new


@author Alex Versus 2021
*/

package main

import "fmt"

type MyType struct {
	name string
}

// Entrypoint
func main() {
	f := new(MyType)
	i := new(int)

	fmt.Println(f)  // &{}
	fmt.Println(*f) // {}

	fmt.Println(*i) // 0

	f.name = "new value"
	fmt.Println(f)
	fmt.Println(*f)

	*i = 3
	fmt.Println(*i) // 3

	v1 := newVal()
	fmt.Println(*v1) // 0

	v2 := newVal2()
	fmt.Println(*v2) // 0

	// different address
	fmt.Println(newVal())
	fmt.Println(newVal())

	st := &MyType{name: "test"}
	fmt.Println(st)

	st2 := new(MyType)
	*st2 = MyType{name: "test2"}
	fmt.Println(st2)
}

func newVal() *int {
	return new(int)
}

func newVal2() *int {
	var i int
	return &i
}

func del(old, new int) int {
	return new - old
}
