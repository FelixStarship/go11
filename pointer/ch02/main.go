package main

import "fmt"

func double(x *int) {
	*x += *x
	fmt.Println(x)
	x = nil
	fmt.Println(x)
}

func main() {
	var a = 3
	double(&a)
	fmt.Println(a)
	var b *int
	var b1 = new(int)
	fmt.Println(b)
	fmt.Println(b1)
}
