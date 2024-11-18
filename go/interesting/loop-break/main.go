package main

import "fmt"

func main() {

	for range 5 {
		fmt.Println("loop 1")
		goto loop2
	}
loop2:
	for range 5 {
		fmt.Println("loop 2")
		break loop2
	}
}
