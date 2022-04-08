package main

import "fmt"

var factorialValue uint64 = 1

func factorial(n int) int {
	if n == 0 { // base case: factorial of 0
		return 1 // factorial of 0 is 1
	}
	// recursive case: multiply n by (n - 1) factorial
	return n * factorial(n-1)
}

// factorial Iterative algorithm
func factorialIterative(n int) uint64 {
	var factorialValue uint64 = 1
	if n < 0 {
		fmt.Println("Factorial of negative numbers is impossible.")
	} else {
		for i := 1; i <= n; i++ {
			factorialValue *= uint64(i)
		}
	}
	return factorialValue
}

func main() {
	fmt.Println(factorial(10))
	fmt.Println(factorialIterative(5))
}
