package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go Language"
	Book1.author = "someone"
	Book1.subject = "Go Lession"
	Book1.book_id = 3245345

	Book2.title = "Pyhton Language"
	Book2.author = "someone"
	Book2.subject = "Python Lession"
	Book2.book_id = 456456

	fmt.Println(Book1.title)
}
