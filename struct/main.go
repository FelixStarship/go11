package main

import "fmt"

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
	book=Book{}
	fmt.Println(book)
	p:=&Book{pages: 10}
    fmt.Println(p)

}
