package main

import (
	"main/github.com/triumphpc/go-cource/9.4/logger"
)

func main() {
	logger.Log("Info")

	logger.Debug(true)
	logger.Log("This is a debug")

	// Compilation error

	//fmt.Println(logger.debug)
}
