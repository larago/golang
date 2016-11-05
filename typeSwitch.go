package main

import "fmt"

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Println(i)
	case int:
		fmt.Printf("int")
	case float64:
		fmt.Println("float64")
	case func(int) float64:
		fmt.Printf("func(int) float64")
	case bool, string:
		fmt.Println("x is bool or string")
	default:
		fmt.Println("Unknown")
	}
}
