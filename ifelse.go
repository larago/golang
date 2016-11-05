package main

import "fmt"

func main() {
	var a int = 100
	if a < 20 {
		fmt.Println("a is smaller than 20")
	} else {
		fmt.Println("a is not smaller than 20")
	}
	fmt.Println("a is equal to", a)
}
