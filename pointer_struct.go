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

	Book1.title = "Go Language"
	Book1.author = "someone"
	Book1.subject = "Go Lession"
	Book1.book_id = 3435345

	printBook(&Book1)

}

func printBook(book *Books) {
	fmt.Println(book.author)
}
