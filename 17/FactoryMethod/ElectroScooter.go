package main

type electricScooter struct {
	transport
}

// factory method
func newElectricScooter() iTransport {
	return &electricScooter{
		transport: transport{
			name:  "Scooter",
			speed: 4,
		},
	}
}
