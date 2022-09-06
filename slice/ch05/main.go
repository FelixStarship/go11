package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := make([]int, 2, 6)
	fmt.Println(len(s), cap(s))
	reflect.ValueOf(&s).Elem().SetLen(3)

	fmt.Println(len(s), cap(s))
	reflect.ValueOf(&s).Elem().SetCap(5)

	fmt.Println(len(s), cap(s))

	//切片克隆
	sClone := make([]int, len(s))
	copy(sClone, s)
	fmt.Println(len(sClone))

	sClone = append(make([]int, 0, 0), s...)
	fmt.Println(sClone)

	//删除切片
	s1 := make([]int, 2, 2)
	s1[0] = 100
	s1[1] = 100
	s1 = append(s1[:1], s1[1+1:]...)

	fmt.Println(s1, len(s1), cap(s1))

	//插入切片元素、插入到开头
	s1 = append(s, s1...)
	fmt.Println(s1)

	//插入切片元素、插入到结尾
	s1 = append(s1, s...)
	fmt.Println(s1)
}
