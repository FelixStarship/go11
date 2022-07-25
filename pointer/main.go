package main

import (
	"fmt"
)

func main() {
	p0 := new(int)
	fmt.Println(p0)
	fmt.Println(*p0)
	fmt.Println(&p0)

}
