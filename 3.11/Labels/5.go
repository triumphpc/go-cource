package main

func main() {
s:
	for {
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
