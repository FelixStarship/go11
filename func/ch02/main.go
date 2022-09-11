package main

import "fmt"

func main() {
	// 闭包
	isMultipleOfx := func(x int) func(int) bool {
		return func(n int) bool {
			return n%x == 0
		}
	}
	var isMultipleOfx3 = isMultipleOfx(3)
	fmt.Println(isMultipleOfx3(6))
}
