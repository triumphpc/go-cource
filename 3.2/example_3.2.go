/**
Go (Golang) From simple to great. The Complete Developer's Guide.
Example 3.2: Conditional statements
If statement

@author Alex Versus 2021
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Entrypoint
func main() {
	var nowTime time.Time = time.Now() // example of method .Now()
	fmt.Println("Today time:", nowTime)
	var day int = nowTime.Day() // method Day of type Time
	fmt.Println("Day:", day)

	// for read enter from keyboard
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please place your age: ")
	userAge, err := reader.ReadString('\n')

	// check on errors that err does not contain errors
	if err != nil {
		log.Fatal(err)
	}
	userAge = strings.TrimSpace(userAge)
	// convert to int
	currentAge, _ := strconv.Atoi(userAge)

	fmt.Println("Your age: ", currentAge)

	// additional information
	if currentAge != 0 {
		fmt.Println("Please take a prize")
	}

	//
	//userAge = strings.TrimSpace(userAge)
	//// convert to int
	//userNumber, err := strconv.Atoi(userAge)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// if statement
	//if userAge < 18 {
	//	fmt.Println("Your age: ", userAge)
	//} else {
	//
	//}

	//
	//for attempts := 0; attempts < 3; attempts ++ {
	//
	//
	//
	//
	//	numberEnter = strings.TrimSpace(numberEnter)
	//	// convert to int
	//	userNumber, err := strconv.Atoi(numberEnter)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	if userNumber > 37 || userNumber < 0 {
	//		fmt.Println("Your can enter number > 0 and < 38. Please enter again")
	//		continue
	//	} else {
	//		break
	//	}
	//}
	//
	//// Next step, user enter the sum
	//fmt.Print("Please place your value: ")
	//valueEnter, err := reader.ReadString('\n')
	//// check on errors that err does not contain errors
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//valueEnter = strings.TrimSpace(valueEnter)
	//// convert to int
	//userValue, err := strconv.Atoi(valueEnter)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("You bet: ", userNumber, " and your value: ", userValue)
	//
	//
	//// generate number for roulette
	//seconds := time.Now().Unix()
	//rand.Seed(seconds)
	//targetNumber := rand.Intn(37)
	//
	//fmt.Println("Roulette give you number: ", targetNumber)
	//
	//// equal user enter and roulette number
	//if targetNumber == userNumber {
	//	fmt.Println("Congratulation! You win! ")
	//} else {
	//	fmt.Println("You lose (")
	//}

}
