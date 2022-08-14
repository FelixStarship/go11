package main

import (
	"fmt"
)

func main() {
	s, s2 := []int{1, 2, 3}, []bool{}
	fmt.Println(len(s), cap(s))
	fmt.Println(len(s2), cap(s2))

	m, m2 := map[int]bool{1: true, 0: false}, map[int]int{}
	fmt.Println(len(m))
	fmt.Println(len(m2))
	if n, present := m[0]; present {
		m[0] = true
		fmt.Println(n, present)
	}
	fmt.Println(m)

	m = nil
	fmt.Println(m[0])
	var m1 map[int]int
	fmt.Println(m1 == nil)
	fmt.Println(m1[0])
	m1[0] = 1

}
