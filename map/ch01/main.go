package main

import "fmt"

func main() {
	type T struct{ age int }
	mt := map[string]T{}
	mt["John"] = T{age: 29}
	ma := map[int][5]int{}
	ma[1] = [5]int{1}

	fmt.Println(mt, ma)

	//部分修改一个映射的元素，非法
	//mt["John"].age=1
	//ma[1][0] = 1

	fmt.Println(mt["John"].age)
	fmt.Println(ma[1][0])
}
