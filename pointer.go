package main

import "fmt"

func main() {

	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	ptr = &a
	fmt.Println(a)
	fmt.Println(*ptr)
}
