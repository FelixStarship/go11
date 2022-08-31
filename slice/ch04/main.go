package main

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	a := make([]int, 32)
	b := a[1:16]
	// a和b不在共享内存片段，a长度不够、开辟新的内存片段
	a = append(a, 1)
	a[2] = 42

	fmt.Println(a)
	fmt.Println(b, len(b), cap(b))
	//-----------------------append()函数切片扩容-----------------------------
	path := []byte("AAAA/BBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') //4
	dir1 := path[:sepIndex]                //双下标
	dir2 := path[:sepIndex:sepIndex]       //[0:4:4]三下标

	fmt.Println(string(dir1))
	fmt.Println(string(dir2))
	fmt.Println(len(dir1), cap(dir1))
	fmt.Println(len(dir2), cap(dir2))
	//--------------------------------------复制对象----------------------------------
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2))
	//prints: s1 == s2: true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2))

	bufio.NewScanner()
}
