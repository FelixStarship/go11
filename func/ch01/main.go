package main

import "fmt"

func Double(n int) int {
	return n + n
}

func Apply(n int, f func(int) int) int {
	return f(n)
}

func main() {
	fmt.Printf("%T\n", Double)
	// 函数类型
	var f func(n int) int
	f = Double
	g := Apply
	fmt.Printf("%T\n", g)
	fmt.Println(f(9))
	fmt.Println(g(6, f))
	fmt.Println(Apply(6, Double))
}
