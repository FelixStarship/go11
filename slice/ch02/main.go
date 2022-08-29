package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"Go": 2007}
	m["C"] = 1972
	m["Java"] = 1995
	fmt.Println(m)
	m["Go"] = 2009
	delete(m, "Java")
	fmt.Println(m)

	s := make(map[int]int, 1)
	s[0] = 1
	s[1] = 2
	s[2] = 2
	fmt.Println(len(s))

	t := make([]int, 1, 3)
	fmt.Println(len(t), cap(t))

	c := map[int]int{}
	c[0] = 1
	a := new(map[int]int)
	a = &c
	fmt.Println(*a)

	slice := new([]int)
	slice = &t

	fmt.Println(*slice)

}
