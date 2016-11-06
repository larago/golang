package main

import "fmt"

func fibonaci(n int) int {
	if n < 2 {
		return n
	} else {
		return fibonaci(n-2) + fibonaci(n-1)
	}
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Println(fibonaci(i), "\t")
	}
}
