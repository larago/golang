package main

import "fmt"

func main() {
	var a int = 20
	var ip *int

	ip = &a

	fmt.Printf("Address of a:%x\n", &a)

	fmt.Printf("Address of ip:%x\n", &ip)

	fmt.Printf("Value of ip: %d\n", *ip)
}
