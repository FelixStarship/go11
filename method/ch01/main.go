package main

import "fmt"

type Age int

//方法、属主参数声明
func (age Age) LargerThan(a Age) bool {
	return age > a
}

func (age *Age) Increase() {
	*age++
}

// 自定义函数类型
type FilterFunc func(in int) bool

func (ff FilterFunc) Filter(in int) bool {
	return ff(in)
}

//自定义映射类型StringSet声明方法
type StringSet map[string]struct{}

func (ss StringSet) Has(key string) bool {
	_, ok := ss[key]
	return ok
}
func (ss StringSet) Add(key string) {
	ss[key] = struct{}{}
}
func (ss StringSet) Remove(key string) {
	delete(ss, key)
}

//自定义结构体类型Book和*Book定义方法
type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}
func (b *Book) SetPages(pages int) {
	b.pages = pages
}

//func Book.Pages(b Book) int {
//	return b.pages
//}

func main() {
	var a int
	a = 1
	fmt.Println(Age(a).LargerThan(10))

	str := StringSet{}
	str.Add("骄傲的少年1")
	str.Add("骄傲的少年2")
	str.Add("骄傲的少年3")
	str.Add("骄傲的少年4")

	//StringSet().Has()

	fmt.Println(str)

	fmt.Println(FilterFunc(func(in int) bool {
		return in > 1
	}).Filter(0))

	//显示调用
	var book Book
	(*Book).SetPages(&book, 985)
	fmt.Println(Book.Pages(book))

	book.SetPages(123)
}
