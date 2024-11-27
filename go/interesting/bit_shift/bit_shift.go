package main

import "fmt"

func main() {
	var x = 16
	fmt.Println(x >> 2)
	// 16 = 10000 >> 2 = 100 = 4

	var y = 1
	fmt.Println(y << 3)
	// 1 = 01 << 3 = 1000 = 8

	var ex = 7
	fmt.Println(2 << ex)
	// 2 = 01 << 7 == 100000000 = 256
}
