package main

import "fmt"

func Factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}
	return result
}

func main() {
	var i int = 15
	fmt.Printf("Factorial of %d is %d\n", i, Factorial(i))
}
