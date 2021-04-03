package main

func main() {
	for {
	s:
		switch 5 {
		case 5:
			switch 3 {
			case 3:
				break s
			}
		}
		break
	}
}
