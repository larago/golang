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

	Book1.title = "Go language"
	Book1.author = "someone"
	Book1.subject = "Go lession"
	Book1.book_id = 435345

	printBook(Book1)

}

func printBook(book Books) {
	fmt.Println(book.title)
	fmt.Println(book.author)
	fmt.Println(book.subject)
	fmt.Println(book.book_id)
}
