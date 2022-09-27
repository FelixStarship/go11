package main

import "fmt"

type Book struct {
	pages int
}

type Books []Book

func (b Book) SetPages(pages int) {
	b.pages = pages
}

func (b *Books) Modify() {
	(*b)[0].pages = 500
	*b = append(*b, Book{789}, Book{900})
}

func main() {
	var b Book
	b.SetPages(123)
	fmt.Println(b.pages)

	var book = Books{{123}, {456}}
	book.Modify()
	fmt.Println(book)
}
