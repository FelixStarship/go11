package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("a:%T\n", a)
	fmt.Println(a[:], a[0:len(a):len(a)])

	fmt.Println(len(a[2:5]), cap(a[2:5]))

	//子切片造成暂时性的内存泄露
	fmt.Println(len(f()), cap(f()))
}

func f() []int {
	s := make([]int, 10, 100)
	return s[50:50]
}
