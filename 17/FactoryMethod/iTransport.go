package main

// Defines an interface for transport objects
type iTransport interface {
	// Set name of transport
	setName(n string)
	// Get name of transport
	getName() string
	// Set speed of transport
	setSpeed(s uint)
	// Get speed of transport
	getSpeed() uint
}
