package main

import "fmt"

func main() {
	fmt.Println(sort72([]int{1, 4, 2, 1, 6, 4, 3, 8, 7, 9, 3}))
	fmt.Println(sort73([]int{1, 4, 2, 1, 6, 4, 3, 8, 7, 9, 3}))
	fmt.Println(sort74([]int{1, 4, 2, 1, 6, 4, 3, 8, 7, 9, 3}))
}

func sort72(arr []int) []int {
	k := 0
	for k < len(arr) {
		t := k
		s := k
		for t < len(arr) {
			if arr[t] < arr[s] {
				s = t
			}
			t += 1
		}
		arr[s], arr[k] = arr[k], arr[s]
		k += 1
	}
	return arr
}

func sort73(arr []int) []int {
	k := 1
	for k < len(arr) {
		t := k
		k += 1
		for t > 0 && arr[t] < arr[t-1] {
			arr[t], arr[t-1] = arr[t-1], arr[t]
			t -= 1
		}
	}
	return arr
}
