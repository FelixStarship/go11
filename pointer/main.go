package main

import (
	"fmt"
)

func main() {
	p0 := new(int)
	fmt.Println(p0)
	fmt.Println(*p0)
	fmt.Println(&p0)
	x := *p0

	fmt.Println(&*p0)
	x = 100

	p1, p2 := &p0, &p0
	fmt.Println(x)
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
	fmt.Println("p1=", *p1)
	fmt.Println("p2=", *p2)
	//fmt.Println(p0==p1)
}
