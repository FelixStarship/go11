package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("a:%T\n", a)
	fmt.Println(a[:], a[0:len(a):len(a)])

	fmt.Println(len(a[2:5]), cap(a[2:5]))

	//子切片造成暂时性的内存泄露
	fmt.Println(len(f()), cap(f()))

	var z = make([]int, 3, 5)

	var _ = (*[3]int)(z)

	//切片未初始化
	var x []int
	m := map[string]int{}
	fmt.Println(x == nil)
	fmt.Println(m == nil)
	m["1"] = 1
	x[0] = 1
	fmt.Println(x)

}

func f() []int {
	s := make([]int, 10, 100)
	return s[50:50]
}
