package main

type quadcopter struct {
	transport
}

// factory method
func newQuadcopter() iTransport {
	return &quadcopter{
		transport: transport{
			name:  "Quadcopter",
			speed: 14,
		},
	}
}
