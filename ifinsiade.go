package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	if a == 100 {
		if b == 200 {
			fmt.Println("a is equal to 100, b is equal to 200")
		}
	}
	fmt.Println(a)
	fmt.Println(b)
}
