package main

import "fmt"

func exam(d []int) {
	d[0] = 1
}

func main() {
	a := 1
	s := new([]int)
	s1 := append(*s, a)
	fmt.Println(s1)

	//拷贝完整切片
	s2 := append(*s, *s...)
	fmt.Println(s2)

	p := make([]int, 1)
	p[0] = 2
	fmt.Println(p)
	exam(p)
	fmt.Println(p)
}
