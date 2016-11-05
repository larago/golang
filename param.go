package main

import "fmt"

var a int = 20

func main() {
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Println(a)
	c = sum(a, b)
	fmt.Println(c)
}

func sum(a, b int) int {
	fmt.Println(a)
	fmt.Println(b)

	return a + b
}
