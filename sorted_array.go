package main

import "fmt"

func isSorted(x []int) bool {
	n := len(x)
	if n == 1 {
		return true
	}
	if x[n-1] < x[n-2] {
		return false
	}
	return isSorted(x[:n-1])
}

func main() {
	x := []int{10, 20, 23, 23, 45, 78, 58}
	fmt.Println(isSorted(x))
	x = []int{10, 20, 3, 23, 45, 78, 88}
	fmt.Println(isSorted(x))
}
