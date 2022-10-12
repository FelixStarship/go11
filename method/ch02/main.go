package main

import "fmt"

type Book struct {
	pages int
}
type StringSet map[string]struct{}

func (s StringSet) Has(key string) bool {
	_, present := s[key]
	return present
}

type Age int

func (a *Age) IsNil() bool {
	return a == nil
}
func (a *Age) Increase() {
	*a++
}
func (b Book) Pages() int {
	return b.pages
}
func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func (b Book) SetPage(pages int) {
	b.pages = pages
}
func main() {
	var book Book
	fmt.Printf("%T\n", book.Pages)
	fmt.Printf("%T\n", (&book).SetPages)
	fmt.Printf("%T\n", (&book).Pages)

	//调用方法
	(&book).SetPages(123)
	book.SetPages(123) //等价于上一行
	fmt.Println(book.Pages())
	fmt.Println((&book).Pages())

	StringSet(nil).Has("key")
	//(*Age)(nil).Increase()  //产生恐慌

	book.SetPage(123) //属主参数的传递是值复制的过程
	fmt.Println(book.pages)
}
