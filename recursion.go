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

func main() {
	twoFactorial := factorial(10)
	fmt.Println(twoFactorial)
}
