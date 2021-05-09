package main

import "fmt"

func main() {

	scooter, _ := getTransport("scooter")
	quad, _ := getTransport("quadcopter")

	fmt.Println(scooter)
	fmt.Println(quad)
}
