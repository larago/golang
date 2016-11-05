package main

import "fmt"

func main() {

	var a bool = true
	var b bool = false
	if a && b {
		fmt.Println("won't print out")
	}
	if a || b {
		fmt.Println("would be printed out")
	}

	// alter a and b
	a = false
	b = true
	if a && b {
		fmt.Println("won't printed out")
	} else {
		fmt.Println("would be printed out")
	}
	if !(a && b) {
		fmt.Println("would be printed out")
	}

}
