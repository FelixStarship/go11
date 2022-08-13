package main

import (
	"fmt"
	"unsafe"
)

type Book struct {
	title, author string
	pages         int
}

func main() {

	book := Book{
		"go语言101",
		"老貘",
		256,
	}

	fmt.Println(book)
	book = Book{}
	fmt.Println(book)
	p := &Book{pages: 10}
	fmt.Println(p)

	var book1 = Book{}
	p1 := &book1.pages
	*p1 = 123
	fmt.Println(book1)

	Book{}.pages = 123
	p2 := &Book{}.pages

	p3 := &Book{pages: 100}
	(*p3).pages = 100
	unsafe.Pointer
}
