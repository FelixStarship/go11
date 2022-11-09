package main

import (
	"fmt"
)

func main() {
	var data interface{} = "great"

	if data, ok := data.(int); ok {
		fmt.Println("[is an int] value =>", data)
	} else {
		fmt.Println("[not an int] value =>", data) //prints: [not an int] value => 0 (not "great")
	}

	if res, ok := data.(int); ok {
		fmt.Println("[is an int] value =>", res)
	} else {
		fmt.Println("[not an int] value =>", data) //prints: [not an int] value => great (as expected)
	}

}

func f1(s []int) {
	_ = s[0] // 第5行： 需要边界检查
	_ = s[1] // 第6行： 需要边界检查
	_ = s[2] // 第7行： 需要边界检查
}

func f2(s []int) {
	_ = s[2] // 第11行： 需要边界检查
	_ = s[1] // 第12行： 边界检查消除了！
	_ = s[0] // 第13行： 边界检查消除了！
}

func f3(s []int, index int) {
	_ = s[index] // 第17行： 需要边界检查
	_ = s[index] // 第18行： 边界检查消除了！
}

func f4(a [5]int) {
	_ = a[4] // 第22行： 边界检查消除了！
}
